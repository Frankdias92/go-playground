package main

import (
	"fmt"
	condition "myFirstProject/Condition"
	"time"
)

// "math"

func main() {
	condition.MySuperCondition()
	fmt.Println(condition.IsWeekend(time.Time{}))
}
