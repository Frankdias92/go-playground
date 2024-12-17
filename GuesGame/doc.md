Let's break down the provided Go code step by step to understand its structure, functionality, and logic. This code implements a simple number guessing game where the user has to guess a randomly generated number between 0 and 100.

## Code Breakdown

### 1. **Package Declaration and Imports**
```go
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)
```
- **package main**: This indicates that this file is part of the main package, which is the entry point for a Go program.
- **Imports**:
  - **bufio**: Provides buffered I/O operations, particularly useful for reading input from the console.
  - **fmt**: Used for formatted I/O operations (e.g., printing to the console).
  - **math/rand**: Used for generating random numbers.
  - **os**: Provides functions for interacting with the operating system, such as reading from standard input.
  - **strconv**: Contains functions for converting strings to other types (e.g., integers).
  - **strings**: Provides functions for manipulating strings.

### 2. **Main Function**
```go
func main() {
	fmt.Println("Guess game!")
	fmt.Println("A random number will be drawn. Try to get it right. The number is an integer between 0 and 100")
```
- The `main` function is where the program starts executing.
- It prints a welcome message explaining the game to the user.

### 3. **Random Number Generation**
```go
x := int64(rand.Intn(101)) // Ensure x is of type int64
```
- A random integer between `0` and `100` is generated using `rand.Intn(101)`. 
- The result is converted to `int64` and stored in variable `x`. This ensures that `x` can be compared with user inputs later, which will also be converted to `int64`.

### 4. **Input Scanner Initialization**
```go
scanner := bufio.NewScanner(os.Stdin)
attempts := make([]int64, 0, 10)
```
- A scanner is created to read input from standard input (the console).
- An empty slice `attempts` of type `int64` is initialized with a capacity of `10`, which will store the user's guesses.

### 5. **Game Loop**
```go
for i := 0; i < 10; i++ {
	fmt.Print("What is your guess? ")
```
- A loop runs for a maximum of `10` iterations, allowing the user up to `10` attempts to guess the number.
- Inside the loop, it prompts the user for their guess.

### 6. **Reading User Input**
```go
scanner.Scan()
attempt := scanner.Text()
attempt = strings.TrimSpace(attempt)
```
- The scanner reads a line of input from the user.
- The input is trimmed of any leading or trailing whitespace using `strings.TrimSpace()`.

### 7. **Input Conversion and Error Handling**
```go
intAttempt, err := strconv.ParseInt(attempt, 10, 64)
if err != nil {
	fmt.Println("The guess needs to be an integer number")
	i-- // Allow retry on invalid input
	continue
}
```
- The trimmed string input is converted to an `int64` using `strconv.ParseInt()`.
- If the conversion fails (e.g., if the user enters non-numeric input), an error message is printed, and the loop counter `i` is decremented (`i--`) so that the user can try again without counting this attempt.
- The loop continues back to the start.

### 8. **Storing Attempts**
```go
attempts = append(attempts, intAttempt)
```
- The valid guess (`intAttempt`) is appended to the `attempts` slice for tracking all guesses made by the user.

### 9. **Guess Evaluation**
```go
switch {
case intAttempt < x:
	fmt.Println("Oops! Try again; the guessed number is bigger than", intAttempt)
case intAttempt > x:
	fmt.Println("Oops! Try again; the guessed number is less than", intAttempt)
case intAttempt == x:
	fmt.Printf("Well done! You got it! The number was %d.\nYour attempts were: %v\n", x, attempts)
	return
}
```
- A switch statement evaluates the user's guess against the target number (`x`):
  - If `intAttempt < x`: It informs the user that their guess was too low.
  - If `intAttempt > x`: It informs them that their guess was too high.
  - If `intAttempt == x`: It congratulates them on guessing correctly and displays both the correct number and all their attempts. The program then exits with a return statement.

### 10. **End of Game Message**
```go
fmt.Printf("SORRY! You missed all the attempts. The correct number was %d.\nYour attempts were: %v\n", x, attempts)
```
- If all attempts are used without a correct guess (after exiting the loop), it prints a message indicating that they have used all attempts and reveals the correct number along with all their guesses.
