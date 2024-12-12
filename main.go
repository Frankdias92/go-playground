package main

import (
	"fmt"
)

func main() {
	x := doDefer()
	fmt.Println(x)
}

func doDefer() int {
	defer fmt.Println(3)
	defer fmt.Println(2)
	fmt.Println(1)

	return 10
}
