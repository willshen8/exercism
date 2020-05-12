package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(input int) bool {
	strInt := strconv.Itoa(input)
	var result float64
	for _, v := range strInt {
		digit, _ := strconv.Atoi(string(v))
		result += math.Pow(float64(digit), float64(len(strInt)))
	}
	return int(result) == input
}
