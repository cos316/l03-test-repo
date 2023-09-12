package main

import (
	"fmt"
	"time"
	"cos316test/utils"
)

func main() {
	fmt.Println("Hello COS316!")

	assignment_due := utils.AssignmentDueToday(time.Now())
	fmt.Println("Is there an assignment due today?", assignment_due)
}
