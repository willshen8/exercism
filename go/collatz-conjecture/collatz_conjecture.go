package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(number int) (int, error) {
	if number <= 0 {
		return 0, errors.New("Only numbers above 0 are allowed")
	}
	var steps int
	for number != 1 {
		if number%2 == 0 {
			number = number / 2
		} else {
			number = number*3 + 1
		}
		steps++
	}
	return steps, nil
}
