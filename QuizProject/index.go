package quizproject

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

// GameState represents the state of the quiz game, including player name, score, and loaded questions.
type GameState struct {
	Name      string
	Points    int
	Questions []Question
}

// Init initializes the game by asking for the player's name.
func (g *GameState) Init() {
	fmt.Println("Welcome to the quiz!")
	fmt.Println("What's your name?")

	// Reads input from the user for their name
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		panic("Error reading name") // Handle input errors
	}

	g.Name = name[:len(name)-1] // Remove trailing newline character from input
	fmt.Printf("Let's play the game, %s!\n", g.Name)
}

// ProccessCSV loads quiz questions from a CSV file
func (g *GameState) ProccessCSV(filePath string) {
	// Open the CSV file
	f, err := os.Open(filePath)
	if err != nil {
		panic("Unable to read the file") // Handle file read errors
	}
	defer f.Close()

	// Parse the CSV file
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		panic("Error reading CSV data") // Handle CSV parsing errors
	}

	// Loop through the CSV records, skipping the header row
	for index, record := range records {
		if index > 0 { // Skip the first row (header)
			correctAnswer, _ := toInt(record[5]) // Convert answer to integer
			question := Question{
				Text:    record[0],
				Options: record[1:5],
				Answer:  correctAnswer,
			}

			// Add parsed question to the game's question list
			g.Questions = append(g.Questions, question)
		}
	}
}

// AskQuestion is the main entry point of the game where the player selects a quiz theme.
func AskQuestion() {
	game := &GameState{Points: 0} // Initialize a new game state
	game.Init()

	// Display quiz themes
	fmt.Println("\033[44m Choose your quiz theme: \033[0m")
	fmt.Println("\033[34m 1. General Knowledge \033[0m")
	fmt.Println("\033[34m 2. Math \033[0m")
	fmt.Println("\033[34m 3. Science \033[0m")

	// Loop until the player selects a valid option
	var choice int
	for {
		reader := bufio.NewReader(os.Stdin)
		readChoice, _ := reader.ReadString('\n')
		var err error
		choice, err = toInt(readChoice[:len(readChoice)-1])

		if err != nil || choice < 1 || choice > 3 {
			fmt.Println("Invalid choice. Please choose between 1 and 3.") // Validation feedback
			continue
		}
		break
	}

	// Load the appropriate CSV file based on the theme choice
	switch choice {
	case 1:
		game.ProccessCSV("QuizProject/quiz-go.csv")
	case 2:
		game.ProccessCSV("QuizProject/math-quiz.csv")
	case 3:
		game.ProccessCSV("QuizProject/science-quiz.csv")
	default:
		fmt.Println("Invalid choice.") // Redundant, but safe
		return
	}

	// Run the quiz
	game.Run()

	// Display final score and performance feedback
	fmt.Printf("End of the game! You scored \033[33m %d points. \033[0m \n", game.Points)
	if game.Points >= 20 {
		fmt.Println("\033[33m Congrats Dev! You passed on my test! \033[33m") // Success message
	} else {
		fmt.Println("\033[41m Oops, looks like you need a bit more practice! \033[0m") // Failure message
	}
}

// Run executes the quiz by iterating through all questions and evaluating answers.
func (g *GameState) Run() {
	for index, question := range g.Questions {
		// Display question and options
		fmt.Printf("\033[33m %d. %s \033[0m\n", index+1, question.Text)
		for j, option := range question.Options {
			fmt.Printf("[%d] %s\n", j+2, option)
		}

		fmt.Println("Type the number of the correct answer")

		// Set a timeout for answering each question
		answerCh := make(chan int) // Channel to capture user input
		go func() {
			for {
				reader := bufio.NewReader(os.Stdin)
				readAnswer, _ := reader.ReadString('\n')
				answerInt, err := toInt(readAnswer[:len(readAnswer)-1])
				if err == nil {
					answerCh <- answerInt // Send answer through the channel
					return
				} else {
					fmt.Println(err.Error()) // Input validation feedback
				}
			}
		}()

		// Handle answer input or timeout
		select {
		case answer := <-answerCh:
			if answer == question.Answer {
				fmt.Println("Congrats! You got it right.") // Correct answer feedback
				g.Points += 10                             // Increment score
			} else {
				fmt.Println("\033[41m Oops! Wrong answer. \033[0m") // Incorrect answer feedback
			}
		case <-time.After(10 * time.Second): // Timeout of 10 seconds
			fmt.Println("Time's up! Moving to the next question.") // Timeout message
		}

		fmt.Println("----------------------------------") // Separator for clarity
	}
}
