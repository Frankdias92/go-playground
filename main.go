package main

import (
	"fmt"
	"myFirstProject/math"
)

func main() {
	fmt.Println("Hello, word!")

	sum := math.Add(3, 3)
	fmt.Println("Sum:", sum)

	sub := math.Subtract(5, 3)
	fmt.Println("sub:", sub)
}

// ro run this run this command
// go run main.go

// you can make the build as well with this command:
// go build main.go

// let's compile to linux
// GOARCH=amd64 go build main.go
