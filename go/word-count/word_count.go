package wordcount

import (
	"log"
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	output := Frequency{}
	parsedWords := ParseWords(phrase)
	formattedSlices := FormatSlice(parsedWords)
	for _, word := range formattedSlices {
		if _, ok := output[word]; !ok {
			output[word] = 1
		} else {
			output[word] += 1
		}
	}
	return output
}

// formatString returns a new string by removing all non-alanumerical and '
func FormatSlice(input []string) []string {
	reg, err := regexp.Compile("[^a-zA-Z0-9']+")
	if err != nil {
		log.Fatal(err)
	}
	output := make([]string, len(input))
	for i, word := range input {
		output[i] = reg.ReplaceAllString(strings.Trim(word, "'"), "")
	}
	return output
}

// ParseWords takes a string and output a slice of words
// It will parse based on any type of whitespaces
func ParseWords(input string) []string {
	return strings.Fields(strings.Replace(strings.ToLower(input), ",", " ", -1))
}
