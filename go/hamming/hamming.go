package hamming

import (
	"errors"
)

// Distance calculates the hamming distance between two strings/DNA strands
func Distance(a, b string) (int, error) {
	if a == "" && b == "" {
		return 0, nil
	}
	if len(a) != len(b) {
		return 0, errors.New("Two strings have different length")
	}
	var distance int
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
