# Quiz Game in Go

This repository contains a simple quiz game implemented in Go. The game allows players to choose a quiz theme, answer questions, and receive feedback based on their performance. The questions are loaded from CSV files, and the game features a timeout for answering each question.

## Concepts Covered

### 1. Structs
- **GameState**: Represents the state of the quiz game, including the player's name, score, and loaded questions.
- **Question**: A structure that holds the text of the question, possible options, and the index of the correct answer.

### 2. Methods
- Methods are defined on the `GameState` struct to initialize the game, process CSV files, run the quiz, and handle user input.

### 3. Context Management
- The game uses goroutines to handle user input concurrently while enforcing a timeout for answers.

### 4. CSV Processing
- Questions are loaded from CSV files using Go's `encoding/csv` package. This allows for easy modification and addition of questions without changing the code.

### 5. Input Handling
- User input is managed using `bufio` to read from standard input, with error handling to ensure valid responses.

## Code Overview

### Main Components

1. **GameState Struct**
   - Holds player information and questions.
   - Initializes the game by asking for the player's name.

2. **ProccessCSV Method**
   - Loads questions from a specified CSV file.
   - Parses each question and its options.

3. **AskQuestion Function**
   - Displays available quiz themes.
   - Calls `ProccessCSV` based on user selection.
   - Executes the quiz using the `Run` method.

4. **Run Method**
   - Iterates through all loaded questions.
   - Displays each question and options.
   - Captures user input with a timeout for each question.

### Example Usage

To run the quiz game:

1. Ensure you have Go installed on your machine.
2. Clone this repository:
   ```
   git clone https://github.com/Frankdias92/go-playground 
   git switch quiz-project
   ```
3. Create or modify CSV files in the `QuizProject` directory:
   - `quiz-go.csv`
   - `math-quiz.csv`
   - `science-quiz.csv`

4. Run the game:
   ```
   go run main.go
   ```

### Sample CSV Format

Each CSV file should have the following format:

```
Question,Option 1,Option 2,Option 3,Option 4,Correct Answer Index
What is the capital of France?,Paris,Berlin,Madrid,Rome,1
```

## Conclusion

This quiz game demonstrates basic concepts of Go programming such as structs, methods, concurrency with goroutines, and file handling with CSVs. It provides an interactive way to test knowledge across various topics while illustrating effective error handling and user input management.

Feel free to modify and expand upon this project by adding more questions or features!
```

### Explanation of Sections

- **Title**: Clearly states what the repository is about.
- **Concepts Covered**: Lists key programming concepts demonstrated in the code.
- **Code Overview**: Provides a high-level summary of how different parts of the code work together.
- **Example Usage**: Gives instructions on how to set up and run the game.
- **Sample CSV Format**: Shows how to format CSV files for loading questions.
- **Conclusion**: Summarizes what users can learn from this project.