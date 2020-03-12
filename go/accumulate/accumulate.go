package accumulate

// Accumulate takesan input and transform the them using the converter specified
func Accumulate(input []string, converter func(string) string) []string {
	var output []string
	for _, v := range input {
		output = append(output, converter(v))
	}
	return output
}
