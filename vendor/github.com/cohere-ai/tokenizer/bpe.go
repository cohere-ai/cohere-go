package tokenizer

import (
	"errors"
	"fmt"
	"math"
)

var (
	specialTokens = int64(1) // this is a wit meme, but for now it shifts vocab indices by 1 because 0 is reserved for padding
)

type Merge struct {
	Merge [2]string
	Count int64
}

type byWordCount []WordCount

func (s byWordCount) Len() int      { return len(s) }
func (s byWordCount) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s byWordCount) Less(i, j int) bool {
	first, second := s[i], s[j]
	switch {
	case first.Count > second.Count:
		return true
	case first.Count < second.Count:
		return false
	default:
		min := int(math.Min(float64(len(first.Pieces)), float64(len(second.Pieces))))
		for k := 0; k < min; k++ {
			if first.Pieces[k] < second.Pieces[k] {
				return true
			} else if first.Pieces[k] > second.Pieces[k] {
				return false
			}
		}
	}

	return false
}

func getPairStatistics(vocab []WordCount) (map[[2]string]int64, map[[2]string]map[int64]int64) {
	stats := map[[2]string]int64{}
	indices := map[[2]string]map[int64]int64{}

	var prevChar string
	for i, symbol := range vocab {
		if len(symbol.Pieces) < 2 {
			continue
		}
		prevChar = symbol.Pieces[0]
		for _, c := range symbol.Pieces[1:] {
			key := [2]string{prevChar, c}
			stats[key] += symbol.Count
			if _, ok := indices[key]; !ok {
				indices[key] = make(map[int64]int64)
			}
			indices[key][int64(i)]++
			prevChar = c
		}
	}

	return stats, indices
}

func getMaxStat(stats map[[2]string]int64) [2]string {
	var maxKeys [][2]string

	maximum := int64(-1)
	for key, count := range stats {
		if count == maximum {
			maxKeys = append(maxKeys, key)
		}
		if count > maximum {
			maximum = count
			maxKeys = [][2]string{key}
		}
	}
	maxKey := maxKeys[0]
	for _, key := range maxKeys[1:] {
		for k := 0; k < 2; k++ {
			if maxKey[k] > key[k] {
				maxKey = key
			} else if maxKey[k] < key[k] {
				break
			}
		}
	}
	return maxKey
}

func pruneStats(stats, bigStats map[[2]string]int64, threshold float64) {
	var pruneCount int64
	for item, freq := range stats {
		if float64(freq) >= threshold {
			continue
		}
		delete(stats, item)
		pruneCount++
		if freq < 0 {
			bigStats[item] += freq
		} else {
			bigStats[item] = freq
		}
	}
}

func deepCopyStats(stats map[[2]string]int64) map[[2]string]int64 {
	newStats := map[[2]string]int64{}
	for k, v := range stats {
		newStats[k] = v
	}
	return newStats
}

type change struct {
	Index     int64
	Word      []string
	OldWord   []string
	Frequency int64
}

func replacePair(bigram [2]string, sortedVocab []WordCount, indices map[[2]string]map[int64]int64) []change {
	changes := []change{}
	for i, freq := range indices[bigram] {
		if freq < 1 {
			continue
		}

		symbol := sortedVocab[i]
		word, freq := symbol.Pieces, symbol.Count
		newWordPieces := replace(word, bigram)
		sortedVocab[i] = WordCount{Pieces: newWordPieces, Count: freq}
		changes = append(changes, change{int64(i), newWordPieces, word, freq})
	}
	return changes
}

func strSliceIndexOf(s []string, elem string, from int) int {
	for i := from; i < len(s); i++ {
		if s[i] == elem {
			return i
		}
	}
	return -1
}

func updatePairStatistics(pair [2]string, changed []change, stats map[[2]string]int64, indices map[[2]string]map[int64]int64) {
	stats[pair] = 0
	indices[pair] = make(map[int64]int64)
	first, second := pair[0], pair[1]
	newPair := first + second
	var prev [2]string
	for _, change := range changed {
		// find all instances of pair, and update frequency/indices around it
		i := 0
		for {
			i = strSliceIndexOf(change.OldWord, first, i)
			if i < 0 {
				break
			}

			// if first symbol is followed by second symbol, we've found an occurrence of pair (old_word[i:i+2])
			if i < len(change.OldWord)-1 && change.OldWord[i+1] == second {
				// assuming a symbol sequence "A B C", if "B C" is merged, reduce the frequency of "A B"
				if i > 0 {
					prev = [2]string{change.OldWord[i-1], change.OldWord[i]}
					stats[prev] -= change.Frequency
					if _, ok := indices[prev]; !ok {
						indices[prev] = map[int64]int64{}
					}
					indices[prev][change.Index]--
				}

				if i < len(change.OldWord)-2 {
					// assuming a symbol sequence "A B C B", if "B C" is merged, reduce the frequency of "C B".
					// however, skip this if the sequence is A B C B C, because the frequency of "C B" will be reduced by the previous code block
					if change.OldWord[i+2] != first || i >= len(change.OldWord)-3 || change.OldWord[i+3] != second {
						nex := [2]string{change.OldWord[i+1], change.OldWord[i+2]}
						stats[nex] -= change.Frequency
						if _, ok := indices[nex]; !ok {
							indices[nex] = map[int64]int64{}
						}
						indices[nex][change.Index]--
					}
				}

				i += 2
			} else {
				i++
			}
		}

		i = 0
		for {
			i = strSliceIndexOf(change.Word, newPair, i)
			if i < 0 {
				break
			}

			if i > 0 {
				prev = [2]string{change.Word[i-1], change.Word[i]}
				stats[prev] += change.Frequency
				if _, ok := indices[prev]; !ok {
					indices[prev] = map[int64]int64{}
				}
				indices[prev][change.Index]++
			}

			if i < len(change.Word)-1 && change.Word[i+1] != newPair {
				nex := [2]string{change.Word[i], change.Word[i+1]}
				stats[nex] += change.Frequency
				if _, ok := indices[nex]; !ok {
					indices[nex] = map[int64]int64{}
				}
				indices[nex][change.Index]++
			}
			i++
		}
	}
}

func baseEncoder() map[string]int64 {
	encoder := make(map[string]int64)
	for k, v := range bytesEncoderInverse {
		encoder[string(k)] = int64(v) + specialTokens
	}
	return encoder
}

func BPE(freq map[string]int64, numSymbols, minFrequency int64) (map[string]int64, []*Merge, error) {
	frequencies := mapToSortedWordCount(freq)

	if minFrequency <= 0 {
		return nil, nil, errors.New("min frequency can't be 0")
	}

	stats, indices := getPairStatistics(frequencies)
	bigStats := deepCopyStats(stats)

	merges := []*Merge{}
	encoder := baseEncoder()
	encoderIdx := int64(len(encoder)) + specialTokens
	if len(freq) == 0 {
		return encoder, merges, nil
	}

	threshold := float64(stats[getMaxStat(stats)]) / 10.0

	for i := int64(0); i < numSymbols; i++ {
		var mostFrequent [2]string
		if len(stats) > 0 {
			mostFrequent = getMaxStat(stats)
		}

		// we probably missed the best pair because of pruning; go back to full statistics
		if len(stats) == 0 || (i > 0 && float64(stats[mostFrequent]) < threshold) {
			pruneStats(stats, bigStats, threshold)
			stats = deepCopyStats(bigStats)
			mostFrequent = getMaxStat(stats)
			threshold = float64(stats[mostFrequent]) * float64(i) / (float64(i) + 10000.0)
			pruneStats(stats, bigStats, threshold)
		}

		if stats[mostFrequent] < minFrequency {
			break
		}

		merges = append(merges, &Merge{Merge: mostFrequent, Count: stats[mostFrequent]})

		encoder[fmt.Sprintf("%s%s", mostFrequent[0], mostFrequent[1])] = encoderIdx
		encoderIdx++

		changes := replacePair(mostFrequent, frequencies, indices)

		updatePairStatistics(mostFrequent, changes, stats, indices)
		stats[mostFrequent] = 0
	}

	return encoder, merges, nil
}
