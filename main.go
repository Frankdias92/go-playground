package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guess game!")
	fmt.Println("A random number will be drawn. Try to get it right. The number is an integer between 0 and 100")

	x := int64(rand.Intn(101)) // Ensure x is of type int64
	scanner := bufio.NewScanner(os.Stdin)
	attempts := make([]int64, 0, 10)

	for i := 0; i < 10; i++ {
		fmt.Print("What is your guess? ")

		scanner.Scan()
		attempt := scanner.Text()
		attempt = strings.TrimSpace(attempt)

		intAttempt, err := strconv.ParseInt(attempt, 10, 64)
		if err != nil {
			fmt.Println("The guess needs to be an integer number")
			i-- // Allow retry on invalid input
			continue
		}

		attempts = append(attempts, intAttempt)

		switch {
		case intAttempt < x:
			fmt.Println("Oops! Try again; the guessed number is bigger than", intAttempt)
		case intAttempt > x:
			fmt.Println("Oops! Try again; the guessed number is less than", intAttempt)
		case intAttempt == x:
			fmt.Printf("Well done! You got it! The number was %d.\nYour attempts were: %v\n", x, attempts)
			return
		}
	}

	fmt.Printf("SORRY! You missed all the attempts. The correct number was %d.\nYour attempts were: %v\n", x, attempts)
}
