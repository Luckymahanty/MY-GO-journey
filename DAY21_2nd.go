package main

import (
	"fmt"
)

func withdraw(balance, amount float64) (float64, error) {
	if amount > balance {
		return balance, fmt.Errorf("Insufficient Balance:  tried to withdraw %.2f but blance is amount is %.2f", amount, balance)
	}
	return balance - amount, nil
}

func main() {
	balance := 500.0
	newBalance, err := withdraw(balance, 600.0)
	if err != nil {
		fmt.Println("Transaction failed:", err)
	} else {
		fmt.Println("Withdrawal Successful ! new balance is :", newBalance)
	}
}
