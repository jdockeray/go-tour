package tour

import (
	"strings"
)

// WordCount It should return a map of the counts of each “word” in the string s.
// The wc. Test function runs a test suite against the provided function and prints
// success or failure.
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordCount := make(map[string]int)

	for _, key := range words {
		_, ok := wordCount[key]

		if ok {
			wordCount[key]++
		} else {
			wordCount[key] = 1
		}

	}
	return wordCount
}

// func main() {
// 	wc.Test(WordCount)
// }
