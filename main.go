package main

import "fmt"

// Global variable for use across multiple functions.
var age int

func main() {
	// Explicit declaration with a defined type.
	var x int8 = 10

	// Constant declaration for immutable values.
	const year string = "1992"

	// Initializing a global variable.
	age = 32

	// Grouped variable declarations with explicit types.
	var (
		name     string = "Franklin"
		lastName string = "Macedo"
	)

	// Printing variable values.
	fmt.Println(name, lastName, age, year, x)
}
