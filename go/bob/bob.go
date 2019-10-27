// Package bob returns a response based on a question/comment
package bob

import (
	"regexp"
	"strings"
	"unicode"
)

// Hey gives Bob's response based a remark by Alice
func Hey(remark string) string {
	trimmedRemark := strings.TrimSpace(remark)

	if IsAskingQuestion(trimmedRemark) {
		if IsYelling(trimmedRemark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	}
	if IsYelling(trimmedRemark) {
		return "Whoa, chill out!"
	}
	if IsNotAddressing(trimmedRemark) {
		return "Fine. Be that way!"
	}
	return "Whatever."
}

// IsYelling checks if the remark is a shout (all capital letters)
func IsYelling(remark string) bool {
	re := regexp.MustCompile("[a-zA-Z]")
	if !re.MatchString(remark) {
		return false
	}

	for _, c := range remark {
		if unicode.IsLetter(c) && unicode.IsLower(c) {
			return false
		}
	}
	return true
}

// IsAskingQuestion checks if the remark is a question
func IsAskingQuestion(remark string) bool {
	re := regexp.MustCompile("[?]$")
	return re.MatchString(remark)
}

// IsNotAddressing checks if a remark is saying nothing
func IsNotAddressing(remark string) bool {
	return strings.TrimSpace(remark) == ""
}
