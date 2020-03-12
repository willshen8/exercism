package strain

type Ints []int
type Lists [][]int
type Strings []string

func (input Ints) Keep(predicate func(int) bool) Ints {
	var output Ints
	for _, v := range input {
		if predicate(v) {
			output = append(output, v)
		}

	}
	return output
}

//    (Ints) Discard(func(int) bool) Ints

func (input Ints) Discard(predicate func(int) bool) Ints {
	var output Ints
	for _, v := range input {
		if !predicate(v) {
			output = append(output, v)
		}

	}
	return output
}

func (input Lists) Keep(predicate func([]int) bool) Lists {
	var output Lists
	for _, v := range input {
		if predicate(v) {
			output = append(output, v)
		}
	}
	return output
}

func (input Strings) Keep(predicate func(string) bool) Strings {
	var output Strings
	for _, v := range input {
		if predicate(v) {
			output = append(output, v)
		}

	}
	return output
}
