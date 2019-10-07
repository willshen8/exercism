package raindrops

import (
	"bytes"
	"strconv"
)

func Convert(n int) string {
	var b bytes.Buffer
	if n%3 == 0 {
		b.WriteString("Pling")
	}
	if n%5 == 0 {
		b.WriteString("Plang")
	}
	if n%7 == 0 {
		b.WriteString("Plong")
	}
	if b.String() == "" {
		return strconv.Itoa(n)
	}
	return b.String()
}
