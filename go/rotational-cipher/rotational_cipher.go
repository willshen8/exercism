package rotationalcipher

import "unicode"

func RotationalCipher(input string, inputShiftKey int) string {
	var result []byte
	for _, v := range []byte(input) {
		var rotatedCipher byte
		if unicode.IsLower(rune(v)) || unicode.IsUpper(rune(v)) {
			rotatedCipher = v + byte(inputShiftKey%26)
			if rotatedCipher > 90 && unicode.IsUpper(rune(v)) ||
				rotatedCipher > 122 && unicode.IsLower(rune(v)) {
				rotatedCipher -= 26
			}
			result = append(result, rotatedCipher)
			continue
		}
		result = append(result, v)
	}
	return string(result)
}
