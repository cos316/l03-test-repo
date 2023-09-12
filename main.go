package main

import (
	"fmt"
	"time"
	"cos316test/utils"
)

func main() {
	fmt.Println("Hello COS316!")

	assignment_due := utils.AssignmentDueToday(time.Now())
	fmt.Printf(
		"Today is a %s. Could there be an assignment due today? %t\n",
		time.Now().Weekday(),
		assignment_due,
	)
}
