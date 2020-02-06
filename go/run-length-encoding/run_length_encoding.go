// Package encode implement run-length encoding and decoding.
package encode

import (
	"strconv"
	"strings"
)

// RunLengthEncode takes an input and encoded it as RLE
func RunLengthEncode(input string) string {
	if input == "" {
		return ""
	}
	if len(input) == 1 {
		return input
	}

	var curr, prev int
	var counter = 1
	var strBuilder strings.Builder

	for curr < len(input) {
		if curr == 0 {
			curr++
			continue
		}
		//check for the end of the string
		if curr == len(input)-1 {
			if input[curr] != input[prev] {
				if counter == 1 {
					strBuilder.WriteString(string(input[prev]))
					strBuilder.WriteString(string(input[curr]))
					curr++
					prev++
					counter = 1
					continue
				} else if counter > 1 {
					strBuilder.WriteString(strconv.Itoa(counter))
					strBuilder.WriteString(string(input[prev]))
					strBuilder.WriteString(string(input[curr]))
					curr++
					prev++
					counter = 1
					continue
				}
			} else if input[curr] == input[prev] {
				counter++
				strBuilder.WriteString(strconv.Itoa(counter))
				strBuilder.WriteString(string(input[curr]))
				curr++
				continue
			}
		}

		if input[curr] != input[prev] {
			if counter == 1 {
				strBuilder.WriteString(string(input[prev]))
				curr++
				prev++
				counter = 1
				continue
			} else if counter > 1 {
				strBuilder.WriteString(strconv.Itoa(counter))
				strBuilder.WriteString(string(input[prev]))
				curr++
				prev++
				counter = 1
				continue
			}
		}

		if input[curr] == input[prev] {
			counter++
			curr++
			prev++
		}

	}
	return strBuilder.String()
}

// RunLengthDecode takes an input and decoded it as RLE
func RunLengthDecode(input string) string {
	return ""
}
