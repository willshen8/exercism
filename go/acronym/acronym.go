// Package acronym gives the acronym of a string
package acronym

import (
	"log"
	"regexp"
	"strings"
)

// Abbreviate returns with the acronym of a string
func Abbreviate(s string) string {
	if s == "" {
		return ""
	}

	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}

	formattedString := strings.Replace(s, "-", " ", -1)
	stringSlice := strings.Split(formattedString, " ")

	var result string
	for _, word := range stringSlice {
		trimmedWord := strings.TrimSpace(word)
		processedWord := reg.ReplaceAllString(trimmedWord, "")
		if processedWord != "" {
			firstChar := processedWord[:1]
			result += strings.ToUpper(firstChar)
		}
	}
	return result
}
