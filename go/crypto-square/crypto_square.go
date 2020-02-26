package cryptosquare

import (
	"log"
	"math"
	"regexp"
	"strings"
)

func Encode(s string) string {
	if s == "" {
		return ""
	}
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	checkError(err)
	newStr := strings.ToLower(reg.ReplaceAllString(s, ""))

	rows := int(math.Floor(math.Sqrt(float64(len(newStr)))))
	charPerRow := int(math.Ceil(float64(len(newStr)) / float64(rows)))
	// normalise rows and charPerRow as per condition: c >= r and c - r <= 1
	if charPerRow-rows > 1 {
		rows = (rows + charPerRow) / 2
		charPerRow = rows
	}

	offset := rows*charPerRow - len(newStr) // used to add paddings/spaces in the end

	var chunks []string
	for i := 0; i < offset; i++ {
		newStr += " "
	}

	for i := 0; i < len(newStr); i += charPerRow {
		chunks = append(chunks, newStr[i:i+charPerRow])
	}

	var result strings.Builder
	for i := 0; i < charPerRow; i++ {
		for j := 0; j < rows; j++ {
			result.WriteString(string(chunks[j][i]))
		}
		if i != charPerRow-1 {
			result.WriteString(" ")
		}
	}

	return result.String()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
