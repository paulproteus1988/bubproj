package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// var ACCOUNT_PATH string = "./files/accounts.txt"
// var TRANSACTION_PATH string = "./files/sledger.txt"

// type Transaction struct {
// 	date, name, category, account string
// 	amount                        float64
// }

// type Account struct {
// 	name, category	string
// 	balance			float64
// }

// Receive pointer to Transaction struct, prompt user for input for Transaction fields
// func addTransaction(t *Transaction) string {
// 	fmt.Print("Date: ")
// 	// Validate date, returns empty string if invalid
// 	date, err := time.Parse("01-02-2006", userInput())
// 	if err != nil {
// 		fmt.Println(err)
// 		return ""
// 	}
// 	t.date = date.Format("01-02-2006")
// 	fmt.Print("Name: ")
// 	t.name = strings.TrimSpace(userInput())
// 	fmt.Print("Account: ")
// 	t.category = strings.TrimSpace(userInput())
// 	fmt.Print("Category: ")
// 	t.account = strings.TrimSpace(userInput())
// 	fmt.Print("Amount: ")
// 	// Validate float, returns empty string if invalid
// 	amount, err2 := strconv.ParseFloat(userInput(), 64)
// 	if err2 != nil {
// 		fmt.Println("Error: Must be a valid value")
// 		return ""
// 	}
// 	t.amount = amount
// 	return fmt.Sprintf("%v %v \n\t%-39v %.2f\n\t%-40v %.2f", t.date, t.name, t.account, -t.amount, t.category, t.amount)
// }

// Receives standard user input, terminates at newline, and returns whitespace trimmed input string
func userInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(input)
}

// Prints transaction text file
// func printTransactions(filePath string) {
// 	file, err := os.ReadFile(filePath)
// 	if err != nil {
//         fmt.Print(err)
//     }

// 	fmt.Print(string(file))
// }

// Writes string to file
func writeFile(filePath string, input string) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	if _, err := f.WriteString(input + "\n\n"); err != nil {
		log.Println(err)
	}
}

// func addAccount(a *Account) string {
// 	a.category = "Credit"
// 	a.name = "Amex"
// 	a.balance = 0.00
// 	return fmt.Sprintf("%s %s %40.2f", a.name, a.category, a.balance)
// }

// func printAccounts(filePath string) {
// 	file, err := os.ReadFile(filePath)
// 	if err != nil {
//         fmt.Print(err)
//     }

// 	fmt.Print(string(file))
// }

// Main menu - Prompts user for keyboard input and triggers chosen function
func mainMenu() {
	fmt.Print("Enter command (Enter \"help\" for help): ")
	switch input := userInput(); input {
	case "add":
		transaction := Transaction{}
		writeFile(TRANSACTION_PATH, addTransaction(&transaction))
	case "areg":
		fmt.Println("Show transactions in particular account")
	case "addaccount":
		account := Account{}
		writeFile(ACCOUNT_PATH, addAccount(&account))	
	case "print":
		printTransactions(TRANSACTION_PATH)
	case "printaccounts":
		printAccounts(ACCOUNT_PATH)
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
