// package proverb helps to build proverbs given a slice of strings
package proverb

import (
	"fmt"
)

// Proverb returns a string slice of proverbs given a slice of verbs
func Proverb(rhyme []string) []string {
	var proverb []string
	for i := 0; i+1 < len(rhyme); i++ {
		verb := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		proverb = append(proverb, verb)
	}
	if len(rhyme) > 0 {
		proverb = append(proverb, "And all for the want of a "+rhyme[0]+".")
	}
	return proverb
}
