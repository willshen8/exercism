package perfect

import "errors"

type Classification int

const (
	ClassificationPerfect Classification = iota
	ClassificationAbundant
	ClassificationDeficient
)

var ErrOnlyPositive = errors.New("Error, only natural numbers are allowed")

// Classify determines if a natural number is a perfect , abundant or deficient number
func Classify(num int64) (Classification, error) {
	factors := Factors(num)
	sum := Sum(factors)

	if num <= 0 {
		return -1, ErrOnlyPositive
	}
	if num == 1 {
		return ClassificationDeficient, nil
	}

	if sum < num {
		return ClassificationDeficient, nil
	} else if sum > num {
		return ClassificationAbundant, nil
	}
	return ClassificationPerfect, nil
}

// Factors determines outputs a string of factors of a natural number
func Factors(num int64) []int64 {
	var result []int64
	result = append(result, 1)

	for i := 2; int64(i) <= num/2; i++ {
		factor := num % int64(i)
		if factor == 0 {
			result = append(result, factor, int64(i))
		}
	}
	return result
}

// Sum returns the sum of the slice
func Sum(factors []int64) int64 {
	var result int64
	for _, v := range factors {
		result += v
	}
	return result
}
