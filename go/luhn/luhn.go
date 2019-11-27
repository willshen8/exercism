// Package luhn uses luhn algorithm is a checksum using modulus 10
package luhn

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var space = regexp.MustCompile(`\s`)

// Valid takes a string and determines if the input is valid based on
// Luhn's algorithm
func Valid(input string) bool {
	// take out empty spaces from input first
	s := strings.TrimSpace(space.ReplaceAllString(input, ""))
	// check for non-number characters and return false if found
	if len(s) <= 1 {
		return false
	}
	if _, err := strconv.Atoi(s); err != nil {
		return false
	}
	result := SumOfString(DoubleSecondDigit(s))
	return result%10 == 0
}

// DoubleSecondDigit doubles every second digit from right and if result is more than 9
// 9 will be subtracted from it
func DoubleSecondDigit(input string) string {
	var result string
	for i := len(input) - 1; i > 0; i -= 2 {
		firstDigit, err := strconv.Atoi(string(input[i]))
		if err != nil {
			fmt.Println(err)
		}
		result += strconv.Itoa(firstDigit)

		secondDigit, err := strconv.Atoi(string(input[i-1]))
		if err != nil {
			fmt.Println(err)
		}
		if (secondDigit * 2) > 9 {
			result += strconv.Itoa(secondDigit*2 - 9)
		} else {
			result += strconv.Itoa(secondDigit * 2)
		}
	}
	return result
}

// SumOfString takes a string representation of a number and sum up its digits
func SumOfString(input string) int {
	var sum int
	for i := 0; i < len(input); i++ {
		digit, err := strconv.Atoi(string(input[i]))
		if err != nil {
			fmt.Println(err)
		}
		sum += digit
	}
	return sum
}
