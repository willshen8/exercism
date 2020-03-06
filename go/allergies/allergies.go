package allergies

import (
	"math"
	"math/bits"
)

type allergicResult struct {
	substance string
	result    bool
}

var allergies = map[uint]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

func Allergies(score uint) []string {
	var result []string
	for i := 0; i < bits.Len(score); i++ {
		if (score>>i)&1 == 1 {
			if found := allergies[uint(math.Pow(float64(2), float64(i)))]; found != "" {
				result = append(result, found)
			}
		}
	}
	return result
}

func AllergicTo(score uint, allergy string) bool {
	allergyResult := Allergies(score)
	for _, v := range allergyResult {
		if v == allergy {
			return true
		}
	}
	return false
}
