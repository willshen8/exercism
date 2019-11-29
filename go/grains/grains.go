// Package grains calculates the total number of grains given the square
package grains

import (
	"errors"
)

// Square returns the amount of grains on a particular square
func Square(n int) (uint64, error) {
	var result uint64 = 1
	if n < 1 || n > 64 {
		return 0, errors.New("the Square on the cheese board must be greater than 0 and less than 65")
	}
	return uint64(result << uint64(n-1)), nil
}

// Total calculates the number of grains based on the square and output
// The result should be a 64 bit digit of all 1s
func Total() uint64 {
	var result uint64 = 1
	for i := 1; i < 65; i++ {
		nextResult := result | (result << uint64(1))
		result = nextResult
	}
	return result
}
