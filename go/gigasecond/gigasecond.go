// Package gigasecond converts a time by adding one gigasecoond
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond returns the new time by adding one giga second
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1000000000 * time.Second)
}
