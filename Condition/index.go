package condition

import (
	"fmt"
	"math"
	"time"
)

// "math"

func MySuperCondition() {
	if x := math.Sqrt(4); x < 1 {
		fmt.Println(x)
	} else if x > 1 {
		fmt.Println(" else value ")
	}
}

func Do(x int) {
	switch x {
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("Another thing")
	}
}

func IsWeekend(x time.Time) bool {
	switch {
	case x.Weekday() > 0 && x.Weekday() < 6:
		return false
	default:
		return true
	}
}

// func doType(x any) {
// 	switch t := x.(type) {
// 	case string:
// 		takeString(t)
// 	case int:
// 	case nil:
// 	}
// }

// func takeString(s string) {
// 	fmt.Println(s)
// }
