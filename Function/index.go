package function

import (
	"fmt"
)

func MakeFunction() {

	// Example call to the `swap` function commented out.
	// `swap` returns two swapped values, but this line is commented:
	// fmt.Println(swap(2, 2))

	// Example of calling `swap` with multiple assignments.
	// Values returned by `swap` would be stored in `a` and `b`.
	// a, b := swap(10, 20)
	// fmt.Println(a, b)

	// Example call to the `divid` function, which returns quotient and remainder.
	// res, rem := divid(10, 3)
	// fmt.Println(res, rem)

	// Example of a higher-order function (function that returns another function).
	// `sumHighOrder(2)` returns a function that adds 2 to an argument.
	// x := sumHighOrder(2)(3) // Adds 2 + 3 and stores the result in `x`.
	// f := sumHighOrder(2)    // `f` is now a function.
	// x := f(3)               // Adds 2 + 3 using the function stored in `f`.
	// fmt.Println(x)

	// Calls the variadic function `sums`, which accepts multiple arguments.
	fmt.Println("sums of 10 + 20 + 20 = ", sums(10, 20, 20)) // Adds 10 + 20 + 20 and prints 50.
}

// Function that adds two integers.
func Sum(a int, b int) int { // or `a, b int` to declare types together.
	return a + b
}

// Function that swaps the values of `a` and `b`.
// Returns the two values in reversed order.
func Swap(a, b int) (int, int) {
	return b, a
}

// Function that divides two integers and returns quotient and remainder.
// Named return values (`res`, `rem`) are optional but make the code more readable.
func Divid(a, b int) (res int, rem int) {
	res = a / b
	rem = a % b
	return res, rem
}

// Higher-order function: returns another function.
// The returned function adds `a` to the given parameter.
func SumHighOrder(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

// Variadic function: accepts a variable number of `int` arguments.
func sums(nums ...int) int {
	var out int // Variable to store the sum.
	// Iterates over the provided arguments.
	for _, n := range nums {
		out += n // Adds each argument to the sum.
	}
	return out // Returns the total.
}
