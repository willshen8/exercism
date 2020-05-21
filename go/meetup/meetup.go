package meetup

import (
	"time"
)

type WeekSchedule int

const (
	First  WeekSchedule = 1
	Second              = 8
	Third               = 15
	Fourth              = 22
	Last                = -7
	Teenth              = 13
)

func Day(week WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	if week == Last {
		t := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)
		date := time.Date(year, month, t.Day(), 0, 0, 0, 0, time.UTC)
		result := date.Day()
		if date.Weekday() == weekday {
			return result
		}
		for date.Weekday() != weekday {
			date = date.AddDate(0, 0, 1)
			result++
		}
		return result + Last
	}
	var result = int(week)
	date := time.Date(year, month, int(week), 0, 0, 0, 0, time.UTC)
	for date.Weekday() != weekday {
		date = date.AddDate(0, 0, 1)
		result++
	}
	return result
}
