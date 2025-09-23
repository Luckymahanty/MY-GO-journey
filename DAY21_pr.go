package main

import (
	"fmt"
)

type ATM struct {
	Balance float64
}

func (a *ATM) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("Invalid amount : %.2f", amount)
	}
	if amount > a.Balance {
		return fmt.Errorf("Insufficiante balance :  You need %.2f but you have %.2f", amount, a.Balance)
	}
	a.Balance -= amount
	fmt.Printf("Successfuly withdraw %.2f . Remaining balance : %.2f\n", amount, a.Balance)
	return nil

}

func main() {

	atm := ATM{Balance: 1000.0}

	if err := atm.Withdraw(200); err != nil {
		fmt.Println("Error:", err)
	}
	if err := atm.Withdraw(900); err != nil {
		fmt.Println("Error:", err)
	}
	if err := atm.Withdraw(-50); err != nil {
		fmt.Println("Error:", err)
	}
}
