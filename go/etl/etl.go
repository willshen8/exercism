package etl

import (
	"strings"
)

func Transform(input map[int][]string) map[string]int {
	var output = make(map[string]int, 26)
	for i, v := range input {
		for _, letter := range v {
			output[strings.ToLower(letter)] = i
		}
	}
	return output
}
