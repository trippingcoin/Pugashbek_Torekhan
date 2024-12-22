package main

import (
	"fmt"

	"github.com/trippingcoin/Pugashbek_Torekhan/Assaignment_1/Assaignment_1/Bank/bank"
)

func main() {
	var account bank.BankAccount

	fmt.Println("Welcome to the Bank Account System")
	fmt.Print("Enter Account Number: ")
	fmt.Scanln(&account.AccountNumber)
	fmt.Print("Enter Holder Name: ")
	fmt.Scanln(&account.HolderName)
	account.Balance = 0.0 // Initialize with zero balance

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Get Balance")
		fmt.Println("4. Exit")
		fmt.Print("> ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var amount float64
			fmt.Print("Enter deposit amount: ")
			fmt.Scanln(&amount)
			account.Deposit(amount)

		case 2:
			var amount float64
			fmt.Print("Enter withdrawal amount: ")
			fmt.Scanln(&amount)
			account.Withdraw(amount)

		case 3:
			account.GetBalance()

		case 4:
			fmt.Println("Exiting program...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
