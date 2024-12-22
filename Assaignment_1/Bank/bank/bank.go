package bank

import "fmt"

type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

func (b *BankAccount) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("Deposit amount must be greater than zero.")
		return
	}
	b.Balance += amount
	fmt.Printf("Successfully deposited %.2f. New Balance: %.2f\n", amount, b.Balance)
}

func (b *BankAccount) Withdraw(amount float64) {
	if amount <= 0 {
		fmt.Println("Withdrawal amount must be greater than zero.")
		return
	}
	if amount > b.Balance {
		fmt.Println("Insufficient balance.")
		return
	}
	b.Balance -= amount
	fmt.Printf("Successfully withdrew %.2f. New Balance: %.2f\n", amount, b.Balance)
}

func (b *BankAccount) GetBalance() {
	fmt.Printf("Current Balance: %.2f\n", b.Balance)
}

func Transaction(account *BankAccount, transactions []float64) {
	for _, amount := range transactions {
		if amount > 0 {
			account.Deposit(amount)
		} else {
			account.Withdraw(-amount)
		}
	}
}
