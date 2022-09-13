package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Transaction struct {
	date, name, category, account string
	amount                        float64
}

// Receives standard user input, terminates at newline, and returns whitespace trimmed input string
func userInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(input)
}

// Main menu - Prompts user for keyboard input and triggers chosen function
func mainMenu() {
	fmt.Print("Enter command (Enter \"help\" for help): ")
	switch input := userInput(); input {
	case "add":
		addTransMenu()
	case "areg":
		fmt.Println("Show transactions in particular account")
	case "accounts":
		fmt.Println("Show accounts")
	case "print":
		fmt.Println("Show transactions")
	case "help":
		fmt.Println("Help")
	default:
		fmt.Println("Unknown command")
		mainMenu()
	}
}

func addTransMenu() {
	fmt.Println("Enter transaction date: ")
	fmt.Println("Enter name for transaction: ")
	fmt.Println("Enter transaction catagory: ")
	fmt.Println("Enter payment account: ")
}

func main() {
	mainMenu()
}
