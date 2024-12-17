package main

import (
	"fmt"
	condition "myFirstProject/Condition"
	function "myFirstProject/Function"
	guesgame "myFirstProject/GuesGame"
	"myFirstProject/arrays"
	"time"
)

// "math"

func main() {
	fmt.Println("hello")
	arrays.GetArrayAndSlice()

	fmt.Println("Print slice")
	// sliceWitoutIndexThree := []int{0,1,2} // this will cause a panic because there is no index on element 3
	sliceWithIndex := []int{0, 1, 2, 3}
	arrays.IsInbounds(sliceWithIndex)

	guesgame.PlayGameGuess()

	function.MakeFunction()
	fmt.Println(function.Divid(2, 2))
	fmt.Println(function.Sum(2, 2))
	fmt.Println(function.Swap(1, 2))

	// if condition
	condition.MySuperCondition()
	fmt.Println(condition.IsWeekend(time.Time{}))
}
