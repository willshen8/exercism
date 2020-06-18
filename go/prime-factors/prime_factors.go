package prime

func Factors(input int64) []int64 {
	var primeFactors = make([]int64, 0)
	var currentFactor int64 = 2

	for input != 1 {
		if input%currentFactor == 0 {
			input = input / currentFactor
			primeFactors = append(primeFactors, currentFactor)
		} else {
			currentFactor++
		}
	}
	return primeFactors
}
