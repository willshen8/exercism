package clock

import (
	"fmt"
)

// Clock represent a real clock with minute and hour
type Clock struct {
	minute int
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minute/60, c.minute%60)
}

// New function create a clock with the minute and hour passed in
func New(h, m int) Clock {
	m = m + h*60
	if m < 0 {
		m %= 24 * 60
		m += 24 * 60
	}
	m %= 24 * 60
	return Clock{minute: m}
}

// Add will add a certain minutes to a clock
func (c Clock) Add(minutes int) Clock {
	return New(0, c.minute+minutes)
}

// Subtract takes away x minutes from a clock
func (c Clock) Subtract(minutes int) Clock {
	return New(0, c.minute-minutes)
}
