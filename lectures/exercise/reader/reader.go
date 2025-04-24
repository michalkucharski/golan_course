//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	// Statistics
	nonBlankLines := 0
	commandsEntered := 0

	fmt.Println("Enter commands (type 'Q' or 'q' to exit):")

	for scanner.Scan() {
		input := scanner.Text()

		// Check for exit condition
		if input == "Q" || input == "q" {
			break
		}

		// Count non-blank lines
		if strings.TrimSpace(input) != "" {
			nonBlankLines++
		}

		// Check for commands and respond
		switch strings.ToLower(input) {
		case "hello":
			fmt.Println("Hello! How are you ?")
			commandsEntered++
		case "hi":
			fmt.Println("Hi there! How are you ?")
			commandsEntered++
		case "bye":
			fmt.Println("Goodbye!")
			commandsEntered++
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}

	// Print statistics
	fmt.Println("\nProgram Statistics:")
	fmt.Printf("Non-blank lines entered: %d\n", nonBlankLines)
	fmt.Printf("Commands entered: %d\n", commandsEntered)
}
