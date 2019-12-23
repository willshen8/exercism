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
	return Clock{minute: NormaliseClock(m + h*60)}
}

// Add will plus a certain minutes to a clock
func (c Clock) Add(minutes int) Clock {
	c.minute = NormaliseClock(c.minute + minutes)
	return c
}

// Subtract take away x minutes from a clock
func (c Clock) Subtract(minutes int) Clock {
	c.minute = NormaliseClock(c.minute - minutes)
	return c
}

// NormaliseClock normalise a clock in the right format
func NormaliseClock(m int) int {
	hour := m / 60
	minute := m % 60

	if minute < 0 {
		hour = hour - 1 - (-minute)/60
		minute = 60 - (-minute)%60
	}
	if hour < 0 {
		hour = 24 - (-hour)%24
	}
	if minute >= 60 {
		hour = hour + m/60
		minute = minute % 60
	}
	if hour >= 24 {
		hour = hour % 24
	}
	return hour*60 + minute
}
