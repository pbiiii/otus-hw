package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(input string) []string {
	if input == "" {
		return []string{}
	}

	wordsArray := strings.Fields(input)
	wordsMap := map[string]int{}

	for i := 0; i < len(wordsArray); i++ {
		word := wordsArray[i]
		if _, ok := wordsMap[word]; ok {
			wordsMap[word]++
		} else {
			wordsMap[word] = 1
		}
	}

	keys := make([]string, 0, len(wordsMap))

	for k := range wordsMap {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		if wordsMap[keys[i]] > wordsMap[keys[j]] {
			return true
		}

		if wordsMap[keys[i]] < wordsMap[keys[j]] {
			return false
		}

		return keys[i] < keys[j]
	})

	return keys[0:10]
}
