package main

<<<<<<< HEAD
import (
	"fmt"
	"myFirstProject/arrays"
)

func main() {
	fmt.Println("hello")
	arrays.GetArrayAndSlice()

	fmt.Println("Print slice")
	// sliceWitoutIndexThree := []int{0,1,2} // this will cause a panic because there is no index on element 3
	sliceWithIndex := []int{0, 1, 2, 3}
	arrays.IsInbounds(sliceWithIndex)
=======
import guesgame "myFirstProject/GuesGame"

func main() {
	guesgame.PlayGameGuess()
>>>>>>> first-project-guessing-game
}
