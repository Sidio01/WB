package main

import (
	"fmt"
	"regexp"
	"sort"
)

type wordCounter struct{}

func (wc *wordCounter) countWords(sentence string) map[string]int {
	re := regexp.MustCompile("\\W+")
	words := re.Split(sentence, -1)
	// words := strings.Split(sentence, re)

	result := make(map[string]int)
	for _, word := range words {
		if _, ok := result[word]; ok {
			result[word]++
		} else {
			result[word] = 1
		}
	}
	return result
}

type textProcessor struct{}

func (tp *textProcessor) processText(sys system, s string) {
	sys.getProcessedText(s)
}

type system interface {
	getProcessedText(s string)
}

type wordCounterAdapter struct {
	wc *wordCounter
}

func (wca *wordCounterAdapter) getProcessedText(s string) {
	m := wca.wc.countWords(s)
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] < m[keys[j]]
	})
	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}

func main() {
	s := "work in progress, progress in progress"

	tp := &textProcessor{}
	wc := &wordCounter{}
	wca := &wordCounterAdapter{wc: wc}
	tp.processText(wca, s)
}
