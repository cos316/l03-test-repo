package utils

import (
	"time"
)

func AssignmentDueToday(t time.Time) bool {
	start, _ := time.Parse(time.RFC822, "05 Sep 23 00:00 ET")
	end, _ := time.Parse(time.RFC822, "16 Dec 23 00:00 ET")

	if t.Before(start) || t.After(end) {
		return false
	}

	return t.Weekday() == time.Wednesday
}

