package sublist

type Relation string

// Sublist output a result whether a is a superlist, sublist, equal or unequal list of b
func Sublist(a, b []int) Relation {
	lenDiff := len(a) - len(b)
	if lenDiff < 0 { // a is the smaller slice
		if len(a) == 0 {
			return "sublist"
		}
		indices := Find(b, a[0])
		for _, v := range indices {
			if Equal(a, b[v:v+len(a)]) {
				return "sublist"
			}
		}
		return "unequal"
	}
	if lenDiff > 0 {
		if len(b) == 0 {
			return "superlist"
		}
		indices := Find(a, b[0])
		for _, v := range indices {
			if Equal(b, a[v:v+len(b)]) {
				return "superlist"
			}
		}
		return "unequal"
	}
	// case when len(a) == len(b)
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return "unequal"
		}
	}
	return "equal"
}

// Find returns the index of the the slice where i is found
func Find(ints []int, s int) []int {
	var foundInts []int
	for i, v := range ints {
		if v == s {
			foundInts = append(foundInts, i)
		}
	}
	return foundInts
}

// Equal determiens if slices a and b are equal
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
