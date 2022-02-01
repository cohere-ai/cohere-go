package tokenizer

import (
	"bufio"
	"io"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"sync"

	"github.com/dlclark/regexp2"
)

var (
	tabRegex     = regexp.MustCompile(`\t`) // Multiple tabs in a row are relevant, but not multiple spaces or newlines
	newlineRegex = regexp.MustCompile(`[\n\r]+`)
	spaceRegex   = regexp.MustCompile(`\p{Z}+`)
	repeatRegex  = regexp2.MustCompile(`.*(.)\1{5,}.*`, 0)
)

type WordCount struct {
	Pieces []string `json:"pieces"`
	Count  int64    `json:"count"`
}

func CountString(s string) map[string]int64 {
	s = tabRegex.ReplaceAllString(s, "\t")
	s = newlineRegex.ReplaceAllString(s, "\n")
	s = spaceRegex.ReplaceAllString(s, " ")

	words := WordSplit(s)
	frequencies := map[string]int64{}
	for _, word := range words {
		token := unicodeEncode(word)

		// only add non-repeating tokens to the frequencies
		if _, ok := frequencies[token]; !ok {
			if match, err := repeatRegex.MatchString(token); match || err != nil {
				continue
			}
		}

		frequencies[token]++
	}
	return frequencies
}

func mapToSortedWordCount(freq map[string]int64) []WordCount {
	counts := make([]WordCount, len(freq))
	idx := 0
	for word, count := range freq {
		counts[idx] = WordCount{Pieces: strings.Split(word, ""), Count: count}
		idx++
	}
	sort.Sort(byWordCount(counts))
	return counts
}

func MergeCounts(a map[string]int64, b map[string]int64) {
	for k, v := range b {
		a[k] += v
	}
}

func CountReader(reader io.Reader) (map[string]int64, error) {
	bufReader := bufio.NewReader(reader)
	wg := sync.WaitGroup{}
	mergeWG := sync.WaitGroup{}
	countJobs := make(chan map[string]int64)
	totalCount := map[string]int64{}
	mergeWorker := func(counts <-chan map[string]int64) {
		for count := range counts {
			MergeCounts(totalCount, count)
		}
		mergeWG.Done()
	}

	mergeWG.Add(1)
	go mergeWorker(countJobs)

	worker := func(jobs <-chan string) {
		for text := range jobs {
			countJobs <- CountString(text)
		}
		wg.Done()
	}

	// assuming each line is 1kb~, this gives us around a 100mb mem cap on this queue
	jobs := make(chan string, 100000)
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go worker(jobs)
	}

	EOF := false
	for {
		if EOF {
			break
		}

		line, err := bufReader.ReadString('\n')
		if err != nil {
			if err != io.EOF && err != io.ErrUnexpectedEOF {
				return nil, err
			}
			EOF = true
		}

		jobs <- line
	}
	close(jobs)
	wg.Wait()

	close(countJobs)
	mergeWG.Wait()

	return totalCount, nil
}
