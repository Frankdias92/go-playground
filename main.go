package main

import (
	"fmt"
	"time"

	"github.com/frankdias92/myFirstProject/database"
	condition "github.com/frankdias92/myFirstProject/first-steps/Condition"
	function "github.com/frankdias92/myFirstProject/first-steps/Function"
	loops "github.com/frankdias92/myFirstProject/first-steps/Loops"
	quizproject "github.com/frankdias92/myFirstProject/first-steps/QuizProject"
	structs "github.com/frankdias92/myFirstProject/first-steps/Structs"
	"github.com/frankdias92/myFirstProject/first-steps/arrays"
	competitionandresources "github.com/frankdias92/myFirstProject/first-steps/competitionAndResources"
	errorhandler "github.com/frankdias92/myFirstProject/first-steps/errorHandler"
	maplesson "github.com/frankdias92/myFirstProject/first-steps/mapLesson"
	myfirtsapigo "github.com/frankdias92/myFirstProject/my-firts-api-go"
	advanceconcepts "github.com/frankdias92/myFirstProject/projects/advance-concepts"
	postgres_db "github.com/frankdias92/myFirstProject/sql"
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

	/* Competition and Resources */
	competitionandresources.TypeParamenters()
	competitionandresources.GoRoutines()
	competitionandresources.GetContextRequest()

	/* Projects / challenges */
	advanceconcepts.GetMeasure()

	/* Run the quiz game */
	quizproject.AskQuestion()
	myfirtsapigo.Server()

	/* multi-database-integration */
	database.DatabaseFoo()
	postgres_db.PostgresMain()
}

func main() {
	myfirtsapigo.Server()
}
