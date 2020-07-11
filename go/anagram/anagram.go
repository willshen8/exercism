package anagram

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var results []string
	subjectSlice := strings.Split(strings.ToLower(subject), "")
	sort.Strings(subjectSlice)

	subjectMap := make(map[string]int)
	for _, v := range subjectSlice {
		if _, ok := subjectMap[v]; !ok {
			subjectMap[v] = 1
		} else {
			subjectMap[v]++
		}
	}

	for _, candidate := range candidates {
		if strings.EqualFold(subject, candidate) {
			break
		}
		candidateSlice := strings.Split(strings.ToLower(candidate), "")
		sort.Strings(candidateSlice)
		candidateMap := make(map[string]int)
		for _, v := range candidateSlice {
			if _, ok := candidateMap[v]; !ok {
				candidateMap[v] = 1
			} else {
				candidateMap[v]++
			}
		}
		if reflect.DeepEqual(subjectMap, candidateMap) {
			results = append(results, candidate)
		}
	}

	fmt.Println("Subject=", subject, subjectMap)
	return results
}
