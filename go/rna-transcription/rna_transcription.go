package strand

import "bytes"

var dnaMapping = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	buf := bytes.Buffer{}
	for _, s := range dna {
		if val, ok := dnaMapping[s]; ok {
			buf.WriteRune(val)
		}
	}
	return buf.String()
}
