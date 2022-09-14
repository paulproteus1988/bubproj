package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var TRANSACTION_PATH string = "./files/transactionlog.txt"

type Transaction struct {
	date, name, category, account string
	amount                        float64
}

// Receive pointer to Transaction struct, prompt user for input for Transaction fields
func addTransaction(t *Transaction) string {
	fmt.Print("Date: ")
	// Validate date, returns empty string if invalid
	date, err := time.Parse("01-02-2006", userInput())
	if err != nil {
		fmt.Println(err)
		return ""
	}
	t.date = date.Format("01-02-2006")
	fmt.Print("Name: ")
	t.name = strings.TrimSpace(userInput())
	fmt.Print("Account: ")
	t.category = strings.TrimSpace(userInput())
	fmt.Print("Category: ")
	t.account = strings.TrimSpace(userInput())
	fmt.Print("Amount: ")
	// Validate float, returns empty string if invalid
	amount, err2 := strconv.ParseFloat(userInput(), 64)
	if err2 != nil {
		fmt.Println("Error: Must be a valid value")
		return ""
	}
	t.amount = amount
	return fmt.Sprintf("%v %v \n\t%-39v %.2f\n\t%-40v %.2f", t.date, t.name, t.account, -t.amount, t.category, t.amount)
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

// Writes string to file
func writeFile(filePath string, input string) {
	f, err := os.Create(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	data := []byte(input)

	_, err2 := f.Write(data)

	if err2 != nil {
		log.Fatal(err2)
	}
}

// Main menu - Prompts user for keyboard input and triggers chosen function
func mainMenu() {
	fmt.Print("Enter command (Enter \"help\" for help): ")
	switch input := userInput(); input {
	case "add":
		transaction := Transaction{}
		transText := addTransaction(&transaction)
		writeFile(TRANSACTION_PATH, transText)
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

func main() {
	mainMenu()
}
