package Defer

import (
	"fmt"
	"os"
)

func RunDefer() int {
	defer fmt.Println(3)
	defer fmt.Println(2)
	fmt.Println(1)

	return 10
}

func LetsOpen() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // Ensures file is closed when main returns

	// Use the file...
}
