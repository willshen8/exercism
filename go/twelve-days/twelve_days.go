package twelve

import (
	"fmt"
	"strings"
)

type verse struct {
	verseNum string
	verse    string
}

var (
	verses = []verse{
		{verseNum: "first", verse: "a Partridge in a Pear Tree"},
		{verseNum: "second", verse: "two Turtle Doves"},
		{verseNum: "third", verse: "three French Hens"},
		{verseNum: "fourth", verse: "four Calling Birds"},
		{verseNum: "fifth", verse: "five Gold Rings"},
		{verseNum: "sixth", verse: "six Geese-a-Laying"},
		{verseNum: "seventh", verse: "seven Swans-a-Swimming"},
		{verseNum: "eighth", verse: "eight Maids-a-Milking"},
		{verseNum: "ninth", verse: "nine Ladies Dancing"},
		{verseNum: "tenth", verse: "ten Lords-a-Leaping"},
		{verseNum: "eleventh", verse: "eleven Pipers Piping"},
		{verseNum: "twelfth", verse: "twelve Drummers Drumming"},
	}
)

// Verse takes an integer n and output the corresponding verse of the day
func Verse(n int) string {
	baseVerse := fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", verses[n-1].verseNum)
	var returnVerse []string
	for i := n - 1; i >= 0; i-- {
		if i == 0 && n != 1 {
			returnVerse = append(returnVerse, "and "+verses[i].verse)
		} else if n == 1 {
			returnVerse = append(returnVerse, verses[i].verse)
		} else {
			returnVerse = append(returnVerse, verses[i].verse)
		}
	}
	return baseVerse + strings.Join(returnVerse, ", ") + "."
}

// Song output the entires verses as a single string
func Song() string {
	var verses []string
	for i := 1; i <= 12; i++ {
		verses = append(verses, Verse(i))

	}
	return strings.Join(verses, "\n")
}
