## Here are the main concepts about errors in Go

```
package main

import (
	"errors"
	"fmt"
	"math"
)

// Custom Error Type Example
type SqrtError struct {
	Value float64
	Msg   string
}

func (e SqrtError) Error() string {
	return fmt.Sprintf("error: %s, value: %f", e.Msg, e.Value)
}

// Function to calculate square root with custom error handling
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, SqrtError{x, "negative value"}
	}
	return math.Sqrt(x), nil
}

// Example of a predefined error
var ErrNotFound = errors.New("not found")

// Function that adds context to an error using fmt.Errorf
func fetchResource(id int) error {
	if id == 0 {
		return fmt.Errorf("fetching resource failed: %w", ErrNotFound)
	}
	return nil
}

// Example demonstrating the use of errors.Is and errors.As
func checkErrorHandling() {
	// Checking if an error matches a specific predefined error
	if err := fetchResource(0); err != nil {
		if errors.Is(err, ErrNotFound) {
			fmt.Println("Resource not found")
		}
	}

	// Example of using a custom error
	res, err := Sqrt(-10)
	if err != nil {
		var sqrtErr SqrtError
		if errors.As(err, &sqrtErr) {
			fmt.Printf("Custom Error - %s\n", sqrtErr)
			return
		}
	}
	fmt.Println("Result:", res)
}

// Example demonstrating multiple errors using errors.Join (Go 1.20+)
var ErrA = errors.New("error A")
var ErrB = errors.New("error B")

func joinErrors() error {
	var combinedError error
	combinedError = errors.Join(ErrA, ErrB)
	return combinedError
}

func handleMultipleErrors() {
	err := joinErrors()
	if err != nil {
		fmt.Println("Combined error:", err)
		if errors.Is(err, ErrA) {
			fmt.Println("Detected ErrA")
		}
		if errors.Is(err, ErrB) {
			fmt.Println("Detected ErrB")
		}
	}
}

// Example demonstrating panic and recover
type SafeExecutor struct{}

func (SafeExecutor) Execute(f func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	f()
}

func demonstratePanicRecovery() {
	executor := SafeExecutor{}
	executor.Execute(func() {
		panic("unexpected error")
	})
}

// Main function to demonstrate all concepts
func main() {
	fmt.Println("--- Error Handling Demonstrations ---")
	fmt.Println("1. Checking Custom Error Handling")
	checkErrorHandling()

	fmt.Println("\n2. Handling Multiple Errors")
	handleMultipleErrors()

	fmt.Println("\n3. Demonstrating Panic Recovery")
	demonstratePanicRecovery()
}
```

> The code above is a comprehensive guide on error handling in Go. It includes commented examples of:

**Creating Custom Errors**:  
Demonstrated through the `SqrtError` type, which encapsulates additional information about the error.

**Using Predefined Errors**:  
An example with `ErrNotFound` to represent specific error conditions.

**Error Handling with `errors.Is` and `errors.As`**:  
Identifies errors with exact matches or attempts to convert them to custom types.

**Combining Multiple Errors (`errors.Join`)**:  
Introduced in Go 1.20, it allows grouping multiple errors into a single instance.

**Panic and Recovery (`panic` and `recover`)**:  
An example that demonstrates how to catch and handle unexpected runtime errors.