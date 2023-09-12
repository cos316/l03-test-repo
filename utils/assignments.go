package utils

import (
	"time"
)

func AssignmentDueToday(t time.Time) bool {
	return t.Weekday() == time.Wednesday
}

