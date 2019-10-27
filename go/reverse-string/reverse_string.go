package reverse

import (
	"strings"
)

func Reverse(input string) string {
	inputAsRunes := []rune(input)
	var sb strings.Builder
	for i := len(inputAsRunes) - 1; i > -1; i-- {
		sb.WriteRune(inputAsRunes[i])
	}
	return sb.String()
}
