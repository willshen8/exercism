package isbn

import (
	"fmt"
	"unicode"
)

// IsValidISBN return a boolean if the input string is a ISBN
func IsValidISBN(isbn string) bool {
	if len(isbn) < 9 {
		return false
	}
	var sum int
	if isbn[len(isbn)-1] == 'X' { // if check character is X, add 10 to sum
		sum += 10
	}
	intSlice := ConvertToIntSlice(isbn)
	if len(intSlice) == 9 && isbn[len(isbn)-1] != 'X' {
		return false
	}
	if len(intSlice) > 10 {
		return false
	}
	for i, s := range intSlice {
		sum += (10 - i) * int(s-'0') // convert rune to int
	}
	fmt.Println("ISBN = ", isbn, "sum = ", sum)
	return sum%11 == 0
}

// ConvertToIntSlice convert a string into int by striping out non alphabet characters
func ConvertToIntSlice(input string) []int {
	var output []int
	for i, s := range input {
		if !unicode.IsDigit(s) {
			continue
		}
		output = append(output, int(input[i]))
	}
	return output
}
