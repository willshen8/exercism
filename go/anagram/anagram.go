package anagram

import (
	"reflect"
	"strings"
)

// Detect returns a slice of strings from candidates that is an anagram of subject
func Detect(subject string, candidates []string) []string {
	var results []string
	subjectMap := convertStringToCharMap(subject)

	for _, candidate := range candidates {
		if strings.EqualFold(subject, candidate) {
			break
		}
		candidateMap := convertStringToCharMap(candidate)
		if reflect.DeepEqual(subjectMap, candidateMap) {
			results = append(results, candidate)
		}
	}
	return results
}

// convertStringToCharMap convert a string into map with its char and its frequencies
func convertStringToCharMap(str string) map[string]int {
	strSlice := strings.Split(strings.ToLower(str), "")
	charMap := make(map[string]int)
	for _, v := range strSlice {
		if _, ok := charMap[v]; !ok {
			charMap[v] = 1
		} else {
			charMap[v]++
		}
	}
	return charMap
}
