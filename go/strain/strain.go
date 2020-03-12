package strain

type Ints []int
type Lists [][]int
type Strings []string

func (input Ints) Keep(predicate func(int) bool) (output Ints) {
	for _, v := range input {
		if predicate(v) {
			output = append(output, v)
		}

	}
	return
}

//    (Ints) Discard(func(int) bool) Ints

func (input Ints) Discard(predicate func(int) bool) (output Ints) {
	for _, v := range input {
		if !predicate(v) {
			output = append(output, v)
		}

	}
	return
}

func (input Lists) Keep(predicate func([]int) bool) (output Lists) {
	for _, v := range input {
		if predicate(v) {
			output = append(output, v)
		}
	}
	return
}

func (input Strings) Keep(predicate func(string) bool) (output Strings) {
	for _, v := range input {
		if predicate(v) {
			output = append(output, v)
		}

	}
	return
}
