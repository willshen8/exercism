package pangram

import (
	"unicode"
)

// IsPangram determines if the input string is a pangram or not
func IsPangram(input string) bool {
	if input == "" {
		return false
	}
	lettersMap := make(map[rune]int)
	for _, v := range input {
		if unicode.IsLetter(v) {
			lettersMap[unicode.ToLower(v)]++
		}
	}

	check := "abcdefghijklmnopqrstuvwxyz"
	for _, v := range check {
		if found := lettersMap[v]; found == 0 {
			return false
		}
	}
	return true
}
