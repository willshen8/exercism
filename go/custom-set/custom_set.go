package stringset

import (
	"sort"
)

type Set []string

func New() Set {
	return Set{}
}

func NewFromSlice(input []string) Set {
	var uniqueInputs = make(map[string]int, 0)
	for _, v := range input {
		if _, ok := uniqueInputs[v]; !ok {
			uniqueInputs[v] = 1
		}
	}
	var result []string
	for k, v := range uniqueInputs {
		if v == 1 {
			result = append(result, k)
		}
	}
	return Set(result)
}

func (s Set) String() string {
	result := "{"
	for i, v := range s {
		result += `"` + v + `"`
		if i != len(s)-1 {
			result += ", "
		}

	}
	result += "}"
	return result
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(compareStr string) bool {
	for _, v := range s {
		if v == compareStr {
			return true
		}
	}
	return false
}

// Subset returns true if s1 is a subset of s2
func Subset(s1, s2 Set) bool {
	if Equal(s1, s2) || len(s1) == 0 {
		return true
	}
	if len(s2) == 0 || len(s2) < len(s1) {
		return false
	}
	var startingIndex int
	for i, v := range s1 {
		if v == s2[i] {
			startingIndex = i
			break
		}
	}
	return Equal(s1, s2[startingIndex:len(s1)])
}

// Disjoint returns true if there are not elements in common
func Disjoint(s1, s2 Set) bool {
	if len(s1) == 0 || len(s2) == 0 {
		return true
	}
	for _, v := range s1 {
		if v == s2[0] {
			return false
		}
	}
	return true
}

// Equals return true if s1 and s2 are exactly the same
func Equal(s1, s2 Set) bool {
	sort.Strings(s1)
	sort.Strings(s2)
	if len(s1) != len(s2) {
		return false
	}
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}

// Add append a new string to s
func (s *Set) Add(input string) {
	if s.Has(input) {
		return
	}
	*s = append(*s, input)
}

// Intersection returns the common elements of s1 and s2
func Intersection(s1, s2 Set) Set {
	result := New()
	for _, v := range s1 {
		if s2.Has(v) {
			result = append(result, v)
		}
	}
	return result
}

// Difference (or Complement) of a set is a set of all elements that are only in the first set
func Difference(s1, s2 Set) Set {
	result := New()
	for _, v := range s1 {
		if !s2.Has(v) {
			result = append(result, v)
		}
	}
	return result
}

// Union contains all the elements in both a and b
func Union(s1, s2 Set) Set {
	result := NewFromSlice(s1)
	for _, v := range s2 {
		if !s1.Has(v) {
			result.Add(v)
		}
	}
	return result
}
