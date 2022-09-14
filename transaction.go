package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
	"time"
)

var TRANSACTION_PATH string = "./files/sledger.txt"

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

func printTransactions(filePath string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
        fmt.Print(err)
    }

	fmt.Print(string(file))
}
