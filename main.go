package main

import (
	"fmt"
	condition "myFirstProject/Condition"
	function "myFirstProject/Function"
	guesgame "myFirstProject/GuesGame"
	loops "myFirstProject/Loops"
	"myFirstProject/arrays"
	maplesson "myFirstProject/mapLesson"
	"time"
)

// "math"

func main() {
	fmt.Println("Hello Dev, be welcome on my mess")
	arrays.GetArrayAndSlice()

	// Arrays
	sliceWithIndex := []int{0, 1, 2, 3}
	arrays.IsInbounds(sliceWithIndex)

	// My first Game in Go
	guesgame.PlayGameGuess()

	function.MakeFunction()
	fmt.Println(function.Divid(2, 2))
	fmt.Println(function.Sum(2, 2))
	fmt.Println(function.Swap(1, 2))

	// if condition
	condition.MySuperCondition()
	fmt.Println(condition.IsWeekend(time.Time{}))

	// loops
	loops.SuperLoop()
	loops.RangeBased()
	loops.GoRoutinesLoop()
	loops.Matrix()

	// map
	maplesson.GetMap()
	maplesson.ForMap()
}
