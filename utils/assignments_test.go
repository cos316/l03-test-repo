package utils

import (
	"testing"
	"math/rand"
	"time"
)

// https://stackoverflow.com/a/43497333, CC BY-SA 3.0 @mkopriva
func randdate() time.Time {
    min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
    max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
    delta := max - min

    sec := rand.Int63n(delta) + min
    return time.Unix(sec, 0)
}

func TestAssignmentDueToday(t *testing.T) {
	for i := 0; i < 1000; i++ {
		// Generate a random date. If its a Wednesday,
		// an assignment should be due...
		date := randdate()

		t.Logf(
			"Testing date %s (%s)! Assignment due today: %t",
			date,
			date.Weekday(),
			AssignmentDueToday(date),
		);

		isWednesday := date.Weekday() == time.Wednesday

		if isWednesday != AssignmentDueToday(date) {
			t.Errorf("Assignment should be due on %s!", date);
		}
	}
}

