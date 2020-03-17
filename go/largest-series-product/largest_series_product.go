package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(input string, span int) (int64, error) {
	if span < 0 || span > len(input) {
		return -1, errors.New("Error with input")
	} else if span == 0 {
		return 1, nil
	}
	var largestProduct int64
	for i := 0; i+span <= len(input); i++ {
		var product = 1
		for j := 0; j < span; j++ {
			digit, err := strconv.Atoi(string(input[i+j]))
			if err != nil {
				return -1, errors.New("Error with input")
			}
			product *= digit
		}
		if int64(product) > largestProduct {
			largestProduct = int64(product)
		}
	}
	return largestProduct, nil
}
