package isogram

import (
	"unicode"
)

// IsIsogram outputs true if input is a isogram
// isogram is a word without a repeating letter
func IsIsogram(input string) bool {
	letters := make(map[rune]int)
	for _, c := range input {
		searchChar := unicode.ToLower(c)
		if unicode.IsLetter(searchChar) {
			if _, found := letters[searchChar]; found {
				return false
			}
			letters[searchChar] = 1
		}
	}
	return true
}
