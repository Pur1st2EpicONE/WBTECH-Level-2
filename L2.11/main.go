package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {

	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}
	fmt.Println(groupAnagrams(words))

}

// groupAnagrams returns a map of anagram groups from a slice of words.
// Key is the first word in lowercase; value is a sorted slice of all words in the group.
func groupAnagrams(words []string) map[string][]string {

	groups := make(map[[33]byte][]string)
	for _, word := range words {
		word = strings.ToLower(word)
		key := makeHistogram(word)
		groups[key] = append(groups[key], word)
	}

	result := make(map[string][]string)
	for _, group := range groups {
		slices.Sort(group)
		group = slices.Compact(group)
		if len(group) > 1 {
			result[group[0]] = group
		}
	}

	return result

}

// makeHistogram creates a letter histogram for a word.
// It returns an array of length 33, where each element corresponds
// to a letter of the Russian alphabet ('а'–'я') and stores
// the number of occurrences of that letter in the word.
func makeHistogram(word string) [33]byte {
	var key [33]byte
	for _, letter := range word {
		key[letter-'а']++
	}
	return key
}
