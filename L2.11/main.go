package main

import (
	"fmt"
	"maps"
	"sort"
	"strings"
	"unicode/utf8"
)

func main() {

	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}
	fmt.Println(groupAnagrams(words))

}

func groupAnagrams(words []string) map[string][]string {

	if len(words) < 2 {
		return nil
	}

	sort.Strings(words)

	res := make(map[string][]string)
	processed := make(map[string]struct{})

	for _, word := range words {
		if _, ok := processed[word]; !ok {
			if notSingle := (len(getAnagrams(processed, word, words)) > 1); notSingle {
				res[strings.ToLower(word)] = getAnagrams(processed, word, words)
			}
		}
	}

	return res

}

func getAnagrams(processed map[string]struct{}, word string, words []string) []string {

	var res []string
	hm := make(map[rune]int)

	for _, char := range word {
		hm[char]++
	}

	for _, currentWord := range words {

		copyMap := make(map[rune]int)
		maps.Copy(copyMap, hm)

		valid := true
		for _, ch := range currentWord {

			if utf8.RuneCountInString(currentWord) != utf8.RuneCountInString(word) {
				valid = false
				break
			}

			copyMap[ch]--
			if copyMap[ch] < 0 {
				valid = false
				break
			}

		}

		if valid {
			res = append(res, strings.ToLower(currentWord))
			processed[currentWord] = struct{}{}
		}

	}

	return res

}
