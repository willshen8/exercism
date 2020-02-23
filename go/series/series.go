package series

// All returns a list of all substrings of s with length n.
func All(n int, s string) []string {
	var output []string
	if len(s)< n{
		return output
	}

	for i:=0; i+n<len(s)+1; i++{
		output = append(output, s[i:i+n])
	}
	return output
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	var output []string
	if len(s)< n{
		return output[0]
	}
	output = append(output, s[0:n])
	return output[0]
}

// First returns a boolean and first string with substring length of n
func First(n int, s string) (first string, ok bool) {
	var output []string
	if len(s)< n{
		return  output[0], false 
	}
	output =  append(output, s[0:n])
	return output[0], true
}