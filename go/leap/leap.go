// Package leap calculates whether a year is a leap year
package leap

// IsLeapYear returns true if a year is int
func IsLeapYear(year int) bool {
	if year%100 == 0 {
		if year%400 == 0 {
			return true
		}
		return false
	}
	if year%4 == 0 {
		return true
	}
	return false
}
