package romannumerals

import (
	"errors"
	"strconv"
)

// I=1, V=5, X=10, L=50, C=100, D=500, M=1000

func ToRomanNumeral(arabic int) (string, error) {
	if arabic < 1 || arabic > 3000 {
		return "", errors.New("Input can't be less than 1 or greater than 3000")
	}
	s := strconv.Itoa(arabic)
	var output string
	for i := 0; i < len(s); i++ {
		length := len(s) - i
		digit, err := strconv.Atoi(string(s[i]))
		if err != nil {
			panic(digit)
		}
		if length == 4 {
			for j := 0; j < digit; j++ {
				output += "M"
			}
		}
		if length == 3 {
			output += CalculateDigit(digit, "C", "D", "M")
		}
		if length == 2 {
			output += CalculateDigit(digit, "X", "L", "C")
		}
		if length == 1 {
			output += CalculateDigit(digit, "I", "V", "X")
		}
	}
	// Algorithm works by converting each digit into roman numerals
	return output, nil
}

// CalculateDigit calculates the corresponding roman numeral given the string for 1, 5 and 10
func CalculateDigit(digit int, one string, five string, ten string) string {
	switch digit {
	case 0:
		return ""
	case 1:
		return one
	case 2:
		return one + one
	case 3:
		return one + one + one
	case 4:
		return one + five
	case 5:
		return five
	case 6:
		return five + one
	case 7:
		return five + one + one
	case 8:
		return five + one + one + one
	case 9:
		return one + ten
	}
	return "Error"
}
