package sublist

type Relation string

// Sublist output a result whether a is a superlist, sublist, equal or unequal list of b
func Sublist(a, b []int) Relation {
	lenDiff := len(a) - len(b)
	if lenDiff > 0 && Contains(b, a) {
		return "superlist"
	} else if lenDiff < 0 && Contains(a, b) {
		return "sublist"
	} else if lenDiff == 0 && Contains(a, b) || len(a) == 0 {
		return "equal"
	}
	return "unequal"
}

// Contains determines if a is a substring of b
func Contains(a, b []int) bool {
	if len(a) == len(b) {
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
	}
	if len(b) == 0 {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return Contains(a, b[1:])
		}
	}
	return true
}
