package atbash

import (
	"unicode"
)

var cipher = map[rune]rune{
	'a': 'z', 'b': 'y', 'c': 'x',
	'd': 'w', 'e': 'v', 'f': 'u',
	'g': 't', 'h': 's', 'i': 'r',
	'j': 'q', 'k': 'p', 'l': 'o',
	'm': 'n', 'n': 'm', 'o': 'l',
	'p': 'k', 'q': 'j', 'r': 'i',
	's': 'h', 't': 'g', 'u': 'f',
	'v': 'e', 'w': 'd', 'x': 'c',
	'y': 'b', 'z': 'a'}

func Atbash(input string) string {
	var cleanedInput string
	for _, c := range input {
		if unicode.IsLetter(c) {
			cleanedInput += (string(cipher[unicode.ToLower(c)]))
		} else if unicode.IsDigit(c) {
			cleanedInput += (string(c))
		}
	}

	var result string
	for i, c := range cleanedInput {
		if i != 0 && i%5 == 0 {
			result += " "
		}
		result += string(c)
	}
	return result
}
