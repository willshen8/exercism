package rotationalcipher

func RotationalCipher(input string, inputShiftKey int) string {
	var result []byte
	for _, v := range []byte(input) {
		var rotatedCipher byte
		if v >= 65 && v <= 90 { // uppercase
			rotatedCipher = v + byte(inputShiftKey%26)
			if rotatedCipher > 90 {
				rotatedCipher -= 26
			}
			result = append(result, rotatedCipher)
			continue
		} else if v >= 97 && v <= 122 { // lowercase
			rotatedCipher = v + byte(inputShiftKey%26)
			if rotatedCipher > 122 {
				rotatedCipher -= 26
			}
			result = append(result, rotatedCipher)
			continue
		}
		result = append(result, v)
	}
	return string(result)
}
