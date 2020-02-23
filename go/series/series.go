package series

// All returns a list of all substrings of s with length n.
func All(n int, s string) []string {
	var output []string
	if len(s) < n {
		return output
	}

	for i := 0; i+n < len(s)+1; i++ {
		output = append(output, s[i:i+n])
	}
	return output
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	if len(s) < n {
		return ""
	}
	return s[0:n]
}

// First returns a boolean and first string with substring length of n
func First(n int, s string) (first string, ok bool) {
	if len(s) < n {
		return "", false
	}
	return s[0:n], true
}
