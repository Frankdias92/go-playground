package main

import (
	"fmt"
	function "myFirstProject/Function"
)

func main() {
	function.MakeFunction()
	fmt.Println(function.Divid(2, 2))
	fmt.Println(function.Sum(2, 2))
	fmt.Println(function.Swap(1, 2))
}
