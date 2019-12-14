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
	if m < 0 {
		h = h - 1 - (-m)/60
		m = 60 - (-m)%60
	}
	if m >= 60 {
		h = h + m/60
		m = m % 60
	}
	if h >= 24 {
		h = h % 24
	}
	if h < 0 {
		h = 24 - (-h)%24
	}
	return Clock{hour: h, minute: m}
}

// Add will plus a certain minutes to a clock
func (c Clock) Add(minutes int) Clock {
	var totalMinutes = minutes + c.minute
	if totalMinutes >= 60 {
		c.minute = totalMinutes % 60
		c.hour += (totalMinutes - c.minute) / 60
		c.hour = c.hour % 24
	} else {
		c.minute += minutes
	}
	return c
}

// Subtract take away x minutes from a clock
func (c Clock) Subtract(minutes int) Clock {
	var hoursToSubtract, minutesToSubtract int
	if minutes >= 60 {
		hoursToSubtract = minutes / 60
		minutesToSubtract = minutes - 60*hoursToSubtract
		c.hour -= hoursToSubtract
	}

	if minutesToSubtract > c.minute || (minutes > c.minute && minutes%60 != 0 && hoursToSubtract == 0) {
		c.hour--
		c.minute = 60 - minutes + c.minute
	} else {
		c.minute -= minutes
	}
	return c.FormatClock()
}

// FormatClock format a clock into the correct format
func (c Clock) FormatClock() Clock {
	if c.hour < 0 {
		c.hour = 24 - (-c.hour)%24
	}
	if c.hour >= 24 {
		c.hour = c.hour % 24
	}
	if c.minute < 0 {
		c.minute = 60 - (-c.minute)%60
	}
	if c.minute >= 60 {
		c.minute = c.minute % 60
	}
	return c
}
