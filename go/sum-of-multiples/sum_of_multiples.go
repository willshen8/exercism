package summultiples

// Found is used inside the map to denote the number exists
// The empty struct takes advantages of not allocating any values
type Found struct{}

func SumMultiples(limit int, divisors []int) int {
	multiples := make(map[int]Found)
	for i := 1; i < limit; i++ {
		for _, v := range divisors {
			if v != 0 && i%v == 0 {
				if _, ok := multiples[i]; !ok {
					multiples[i] = Found{}
				}
			}
		}
	}
	return SumIntSlice(multiples)
}

//SumIntMap sums up all the ints in the int slice
func SumIntSlice(s map[int]Found) int {
	var result int
	for k, _ := range s {
		result += k
	}
	return result
}
