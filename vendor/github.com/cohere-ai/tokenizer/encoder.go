package tokenizer

import (
	"bufio"
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/pkg/errors"
)

//go:embed vocab/*
var f embed.FS

var (
	splitRegex                        = regexp2.MustCompile(`(?i:'s|'t|'re|'ve|'m|'ll|'d| ?\p{L}+| ?\p{N}+| ?[^\s\p{L}\p{N}]+|\s+(?!\S)|\s+)`, 0)
	bytesEncoder, bytesEncoderInverse = bytesToUnicode()
)

type Encoder struct {
	Encoder   map[string]int64
	Decoder   map[int64]string
	BPERanks  map[[2]string]int64
	Cache     map[string]string
	VocabSize int64
}

func NewFromReaders(encoderReader, vocabReader io.Reader) (*Encoder, error) {
	bpeMerges := [][2]string{}

	vocabScanner := bufio.NewScanner(vocabReader)
	for vocabScanner.Scan() {
		// each line will look something like:
		// fanta stic 4234234
		// we ignore the last count column for encoding purposes
		split := strings.Split(vocabScanner.Text(), " ")

		bpeMerges = append(bpeMerges, [2]string{split[0], split[1]})
	}

	encoderContents, err := ioutil.ReadAll(encoderReader)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read encoder file")
	}

	encoderMap := map[string]int64{}
	if err := json.Unmarshal(encoderContents, &encoderMap); err != nil {
		return nil, errors.Wrap(err, "corrupted encoder file")
	}

	return New(encoderMap, bpeMerges)
}

func NewFromPrebuilt(name string) (*Encoder, error) {
	encoderPath := fmt.Sprintf("vocab/%s/encoder.json", name)
	vocabPath := fmt.Sprintf("vocab/%s/vocab.bpe", name)

	_, encoderOpenErr := f.Open(encoderPath)
	_, vocabOpenErr := f.Open(vocabPath)
	if vocabOpenErr != nil || encoderOpenErr != nil {
		return nil, errors.New("failed to load prebuilt tokenizer")
	}
	encoderContents, err := f.ReadFile(encoderPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read encoder file")
	}
	encoderMap := map[string]int64{}
	if err := json.Unmarshal(encoderContents, &encoderMap); err != nil {
		return nil, errors.Wrap(err, "encoder file had invalid json")
	}

	vocabContents, err := f.ReadFile(vocabPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read vocab file")
	}
	vocabScanner := bufio.NewScanner(bytes.NewReader(vocabContents))

	bpeMerges := [][2]string{}
	for vocabScanner.Scan() {
		split := strings.Split(vocabScanner.Text(), " ")

		bpeMerges = append(bpeMerges, [2]string{split[0], split[1]})
	}

	return New(encoderMap, bpeMerges)
}

func New(encoder map[string]int64, bpeMerges [][2]string) (*Encoder, error) {
	vocabSize := int64(0)
	decoder := map[int64]string{}
	for k, v := range encoder {
		decoder[v] = k
		vocabSize++
	}

	bpeRanks := map[[2]string]int64{}
	for i := int64(0); i < int64(len(bpeMerges)); i++ {
		bpeRanks[bpeMerges[i]] = i
	}

	return &Encoder{
		Encoder:   encoder,
		Decoder:   decoder,
		BPERanks:  bpeRanks,
		Cache:     map[string]string{},
		VocabSize: vocabSize,
	}, nil
}

func getPairs(wordPieces []string) [][2]string {
	if len(wordPieces) == 0 {
		return nil
	}

	pairs := [][2]string{}
	prevChar := wordPieces[0]
	for _, wordPiece := range wordPieces[1:] {
		pairs = append(pairs, [2]string{prevChar, wordPiece})
		prevChar = wordPiece
	}

	return pairs
}

func (e *Encoder) getMinPair(pairs [][2]string) [2]string {
	minimumPair := pairs[0]
	outOfVocab := int64(len(e.BPERanks)) + 1
	for _, pair := range pairs[1:] {
		pairVal, ok := e.BPERanks[pair]
		if !ok {
			pairVal = outOfVocab
		}

		minimumVal, ok := e.BPERanks[minimumPair]
		if !ok {
			minimumVal = outOfVocab
		}

		if pairVal < minimumVal {
			minimumPair = pair
		}
	}

	return minimumPair
}

func (e *Encoder) bPE(token string) []string {
	wordPieces := strings.Split(token, "")
	pairs := getPairs(wordPieces)
	if len(pairs) == 0 {
		return []string{token}
	}

	for {
		bigram := e.getMinPair(pairs)
		if _, ok := e.BPERanks[bigram]; !ok {
			break
		}

		newWord := replace(wordPieces, bigram)
		wordPieces = newWord
		if len(wordPieces) == 1 {
			break
		} else {
			pairs = getPairs(wordPieces)
		}
	}

	return wordPieces
}

func (e *Encoder) encodeWords(words []string) []int64 {
	bpeTokens := []int64{}

	for _, word := range words {
		token := unicodeEncode(word)
		bpeEncoded := e.bPE(token)
		for _, bpeEnc := range bpeEncoded {
			if _, ok := e.Encoder[bpeEnc]; ok {
				bpeTokens = append(bpeTokens, e.Encoder[bpeEnc])
			}
		}
	}

	return bpeTokens
}

func (e *Encoder) Encode(text string) []int64 {
	words := wordSplit(text)
	return e.encodeWords(words)
}

func (e *Encoder) Decode(tokens []int64) string {
	var decodeBuffer bytes.Buffer
	for _, token := range tokens {
		for _, dt := range e.Decoder[token] {
			decodeBuffer.WriteByte(bytesEncoderInverse[dt])
		}
	}

	return decodeBuffer.String()
}

func unicodeEncode(word string) string {
	var tokenBuffer bytes.Buffer

	for _, b := range []byte(word) {
		encodedRune := bytesEncoder[b]
		tokenBuffer.WriteRune(encodedRune)
	}

	word = tokenBuffer.String()
	return word
}

func wordSplit(s string) []string {
	results := make([]string, 0)
	wordsMatch, _ := splitRegex.FindStringMatch(s)
	if wordsMatch == nil {
		return nil
	}

	for {
		word := wordsMatch.String()

		if word != "" {
			results = append(results, word)
		}

		wordsMatch, _ = splitRegex.FindNextMatch(wordsMatch)
		if wordsMatch == nil {
			break
		}
	}

	return results
}

func runeContains(bs []int, b int) bool {
	for _, v := range bs {
		if b == v {
			return true
		}
	}
	return false
}

func bytesToUnicode() (map[byte]rune, map[rune]byte) {
	bs := []int{}
	for i := 33; i < 127; i++ {
		bs = append(bs, i)
	}
	for i := 161; i < 173; i++ {
		bs = append(bs, i)
	}
	for i := 174; i < 256; i++ {
		bs = append(bs, i)
	}

	cs := make([]int, 0)
	for i := 0; i < len(bs); i++ {
		cs = append(cs, bs[i])
	}

	n := 0
	for b := 0; b < 256; b++ {
		if !runeContains(bs, b) {
			bs = append(bs, b)
			cs = append(cs, 256+n)
			n++
		}
	}

	result := map[byte]rune{}
	for i := range bs {
		result[byte(bs[i])] = rune(cs[i])
	}

	resultInverse := map[rune]byte{}
	for k, v := range result {
		resultInverse[v] = k
	}
	return result, resultInverse
}

func indexOf(wordPieces []string, word string, i int64) int64 {
	for j := i; j < int64(len(wordPieces)); j++ {
		if word == wordPieces[j] {
			return j
		}
	}

	return -1
}
func replace(wordPieces []string, bigram [2]string) []string {
	first, second := bigram[0], bigram[1]
	pairStr := fmt.Sprintf("%s%s", first, second)
	newWord := []string{}
	i := int64(0)
	for i < int64(len(wordPieces)) {
		j := indexOf(wordPieces, first, i)
		if j >= 0 {
			newWord = append(newWord, wordPieces[i:j]...)
			i = j
		} else {
			newWord = append(newWord, wordPieces[i:]...)
			break
		}

		if wordPieces[i] == first && i < int64(len(wordPieces)-1) && wordPieces[i+1] == second {
			newWord = append(newWord, pairStr)
			i += 2
		} else {
			newWord = append(newWord, wordPieces[i])
			i++
		}
	}
	return newWord
}
