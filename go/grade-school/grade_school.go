package school

import (
	"sort"
)

// Grade contains a list of students in a particular grade
type Grade struct {
	level    int
	students []string
}

// School consists a list of grades
type School struct {
	levels []Grade
}

// New is a constructor that initialize a new school
func New() *School {
	return &School{}
}

// Add adds a new student to a particular level
func (s *School) Add(name string, level int) {
	for i, v := range s.levels {
		if v.level == level {
			s.levels[i].students = append(s.levels[i].students, name)
			return
		}
	}
	s.levels = append(s.levels, Grade{level: level, students: []string{name}})
}

// Grade returns a list of all students in a particular grade sorted by name
func (s *School) Grade(grade int) []string {
	var result []string
	for _, v := range s.levels {
		if v.level == grade {
			result = v.students
		}
	}
	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
	return result
}

// Enrollment returns a sorted list of all students by grade first then name
func (s *School) Enrollment() []Grade {
	sort.Slice(s.levels, func(i, j int) bool { return s.levels[i].level < s.levels[j].level })
	for _, v := range s.levels {
		sort.Slice(v.students, func(i, j int) bool { return v.students[i] < v.students[j] })
	}
	return s.levels
}
