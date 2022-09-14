package main

import (
	"fmt"
	"os"
)

var ACCOUNT_PATH string = "./files/accounts.txt"

type Account struct {
	name, category	string
	balance			float64
}

func addAccount(a *Account) string {
	a.category = "Credit"
	a.name = "Amex"
	a.balance = 0.00
	return fmt.Sprintf("%s %s %40.2f", a.name, a.category, a.balance)
}

func printAccounts(filePath string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
        fmt.Print(err)
    }

	fmt.Print(string(file))
}
