package clock

import (
	"fmt"
)

// MINUTESINADAY gives the number of minutes in one day
const MINUTESINADAY = 1440

// Clock represent a real clock with minute and hour
type Clock struct {
	minute int
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minute/60, c.minute%60)
}

// New function create a clock with the minute and hour passed in
func New(h, m int) Clock {
	return Clock{minute: NormaliseClock(m + h*60)}
}

// Add will add a certain minutes to a clock
func (c Clock) Add(minutes int) Clock {
	return New(0, c.minute+minutes)
}

// Subtract takes away x minutes from a clock
func (c Clock) Subtract(minutes int) Clock {
	return New(0, c.minute-minutes)
}

// NormaliseClock normalise a clock in the right format
func NormaliseClock(m int) int {
	if m < 0 {
		m = m % MINUTESINADAY
		m = MINUTESINADAY + m
	}
	return m % MINUTESINADAY
}
