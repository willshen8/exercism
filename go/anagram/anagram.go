package anagram

import (
	"reflect"
	"sort"
	"strings"
)

// Detect returns a slice of strings from candidates that is an anagram of subject
func Detect(subject string, candidates []string) []string {
	var results []string
	subjectSlice := strings.Split(strings.ToLower(subject), "")
	sort.Strings(subjectSlice)

	for _, candidate := range candidates {
		if strings.EqualFold(subject, candidate) {
			break
		}
		candidateSlice := strings.Split(strings.ToLower(candidate), "")
		sort.Strings(candidateSlice)
		if reflect.DeepEqual(subjectSlice, candidateSlice) {
			results = append(results, candidate)
		}
	}
	return results
}
