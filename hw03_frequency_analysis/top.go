package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type WordWithFreq struct {
	String string
	Freq   int
}

type WordsByFreq []WordWithFreq

func (w WordsByFreq) Len() int {
	return len(w)
}

func (w WordsByFreq) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w WordsByFreq) Less(i, j int) bool {
	iWord := w[i]
	jWord := w[j]

	if iWord.Freq > jWord.Freq {
		return true
	}

	if iWord.Freq < jWord.Freq {
		return false
	}

	return iWord.String < jWord.String
}

func Top10(input string) []string {
	if input == "" {
		return nil
	}

	wordsArray := strings.Fields(input)
	wordsMap := map[string]int{}

	for i := 0; i < len(wordsArray); i++ {
		wordsMap[wordsArray[i]]++
	}

	byFreqMap := make(WordsByFreq, len(wordsMap))
	for word, freq := range wordsMap {
		byFreqMap = append(byFreqMap, WordWithFreq{
			String: word,
			Freq:   freq,
		})
	}

	sort.Stable(byFreqMap)

	var firstWords []WordWithFreq

	if byFreqMap.Len() > 10 {
		firstWords = byFreqMap[0:10]
	} else {
		firstWords = byFreqMap
	}

	keys := make([]string, len(firstWords))
	for _, w := range firstWords {
		keys = append(keys, w.String)
	}

	return keys
}
