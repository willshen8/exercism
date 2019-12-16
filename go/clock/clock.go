package clock

import (
	"fmt"
)

// Clock represent a real clock with minute and hour
type Clock struct {
	hour   int
	minute int
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// New function create a clock with the minute and hour passed in
func New(h, m int) Clock {
	return Clock{hour: h, minute: m}.FormatClock()
}

// Add will plus a certain minutes to a clock
func (c Clock) Add(minutes int) Clock {
	c.minute += minutes
	return c.FormatClock()
}

// Subtract take away x minutes from a clock
func (c Clock) Subtract(minutes int) Clock {
	// var hoursToSubtract, minutesToSubtract int
	// if minutes >= 60 {
	// 	hoursToSubtract = minutes / 60
	// 	minutesToSubtract = minutes - 60*hoursToSubtract
	// 	c.hour -= hoursToSubtract
	// }

	// if minutesToSubtract > c.minute || (minutes > c.minute && minutes%60 != 0 && hoursToSubtract == 0) {
	// 	c.hour--
	// 	c.minute = 60 - minutes + c.minute
	// } else {
	// 	c.minute -= minutes
	// }
	c.minute-= minutes
	return c.FormatClock()
}

// FormatClock format a clock into the correct format
func (c Clock) FormatClock() Clock {
	if c.minute < 0 {
		c.hour = c.hour - 1 - (-c.minute)/60
		c.minute = 60 - (-c.minute)%60
	}
	if c.hour < 0 {
		c.hour = 24 - (-c.hour)%24
	}
	if c.minute >= 60 {
		c.hour = c.hour + c.minute/60
		c.minute = c.minute % 60
	}
	if c.hour >= 24 {
		c.hour = c.hour % 24
	}
	return c
}
