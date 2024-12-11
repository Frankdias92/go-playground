package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Explicit conversion examples
	var x int = 10084
	f := float64(x) // Converts int to float64
	fmt.Println("Converted to float64:", f)

	// Proper string conversion using strconv
	y := strconv.FormatInt(int64(x), 10) // Converts int64 to a string
	fmt.Println("Converted to string:", y)

	// Demonstrating constant flexibility
	const z = 10 // Untyped constant
	takeInt32(z) // Untyped constant adapts to int32
	takeInt64(z) // Untyped constant adapts to int64

	// Typed constant
	const typedConst int32 = 15
	takeInt32(typedConst)

	// String example
	takeString("This is a constant string")
}

// Accepts int32
func takeInt32(x int32) {
	fmt.Println("Received int32:", x)
}

// Accepts int64
func takeInt64(x int64) {
	fmt.Println("Received int64:", x)
}

// Accepts string
func takeString(s string) {
	fmt.Println("Received string:", s)
}
