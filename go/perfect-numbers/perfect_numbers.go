package perfect

import "errors"

type Classification struct {
	name string
}

var (
	ClassificationPerfect   = Classification{name: "ClassificationPerfect"}
	ClassificationAbundant  = Classification{name: "ClassificationAbundant"}
	ClassificationDeficient = Classification{name: "ClassificationDeficient"}
	ErrOnlyPositive         = errors.New("Error, only natural numbers are allowed")
)

// Classify determines if a natural number is a perfect , abundant or deficient number
func Classify(num int64) (Classification, error) {
	factors := Factors(num)
	sum := Sum(factors)

	if num <= 0 {
		return Classification{}, ErrOnlyPositive
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
	// instead of going through the whole array, we only need to go through
	// half of the items
	var midPoint int64

	if num%2 == 0 {
		midPoint = num / 2
	} else {
		midPoint = (num + 1) / 2
	}
	for i := 2; int64(i) <= midPoint; i++ {
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
