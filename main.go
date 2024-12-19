package main

import (
	"fmt"
	"time"

	condition "github.com/frankdias92/myFirstProject/Condition"
	function "github.com/frankdias92/myFirstProject/Function"
	loops "github.com/frankdias92/myFirstProject/Loops"
	quizproject "github.com/frankdias92/myFirstProject/QuizProject"
	structs "github.com/frankdias92/myFirstProject/Structs"
	"github.com/frankdias92/myFirstProject/arrays"
	errorhandler "github.com/frankdias92/myFirstProject/errorHandler"
	maplesson "github.com/frankdias92/myFirstProject/mapLesson"
)

func Test() {
	fmt.Println("Hello Dev, be welcome on my mess")
	arrays.GetArrayAndSlice()

	/* Arrays */
	sliceWithIndex := []int{0, 1, 2, 3}
	arrays.IsInbounds(sliceWithIndex)

	/* My first Game in Go */
	/* guesgame.PlayGameGuess() */

	function.MakeFunction()
	fmt.Println(function.Divid(2, 2))
	fmt.Println(function.Sum(2, 2))
	fmt.Println(function.Swap(1, 2))

	/* if condition */
	condition.MySuperCondition()
	fmt.Println(condition.IsWeekend(time.Time{}))
	condition.GetResult(6)

	/* loops */
	loops.SuperLoop()
	loops.RangeBased()
	loops.GoRoutinesLoop()
	loops.Matrix()

	/* map */
	maplesson.GetMap()
	maplesson.ForMap()

	/*  struct  */
	structs.Struct()
	structs.Interface()
	structs.SwitchStatement()
	structs.TypeStruct()

	/* Error handling */
	errorhandler.HandleError()
	errorhandler.Error1()
	errorhandler.Error2()
	errorhandler.Error3()
	errorhandler.Error4()
	errorhandler.Error5()
	errorhandler.Reader()
}

func main() {
	// Run the quiz game
	quizproject.AskQuestion()
}
