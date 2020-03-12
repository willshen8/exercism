// Package encode implement run-length encoding and decoding.
package encode

import (
	"strconv"
	"strings"
)

// RunLengthEncode takes an input and encoded it as RLE
func RunLengthEncode(input string) string {
	if len(input) <= 1 {
		return input
	}

	var curr int
	var counter = 1
	var strBuilder strings.Builder

	for curr < len(input) {
		if curr == 0 {
			curr++
		}
		//check for the end of the string
		if curr == len(input)-1 {
			if input[curr] != input[curr-1] {
				if counter == 1 {
					strBuilder.WriteString(string(input[curr-1]))
					strBuilder.WriteString(string(input[curr]))
					curr++
					counter = 1
					continue
				} else if counter > 1 {
					strBuilder.WriteString(strconv.Itoa(counter))
					strBuilder.WriteString(string(input[curr-1]))
					strBuilder.WriteString(string(input[curr]))
					curr++
					counter = 1
					continue
				}
			} else if input[curr] == input[curr-1] {
				counter++
				strBuilder.WriteString(strconv.Itoa(counter))
				strBuilder.WriteString(string(input[curr]))
				curr++
				continue
			}
		}
		// when current char is not equal to previous one
		if input[curr] != input[curr-1] {
			if counter == 1 {
				strBuilder.WriteString(string(input[curr-1]))
				curr++
				counter = 1
				continue
			} else if counter > 1 {
				strBuilder.WriteString(strconv.Itoa(counter))
				strBuilder.WriteString(string(input[curr-1]))
				curr++
				counter = 1
				continue
			}
		}
		// when current char equals to previous one
		if input[curr] == input[curr-1] {
			counter++
			curr++
		}
	}
	return strBuilder.String()
}

// RunLengthDecode takes an input and decoded it as RLE
func RunLengthDecode(input string) string {
	if len(input) <= 1 {
		return input
	}

	var curr int
	var output, repeatFactor string

	for curr < len(input) {
		_, err := strconv.Atoi(string(input[curr]))
		if err == nil {
			repeatFactor += string(input[curr])
			curr++
			continue
		} else { // else we have read a character
			result, _ := strconv.Atoi(repeatFactor)
			if result == 0 {
				output += string(input[curr])
			} else {
				for i := 0; i < result; i++ {
					output += string(input[curr])
				}
			}
			curr++
			repeatFactor = ""
		}

	}

	return output
}
