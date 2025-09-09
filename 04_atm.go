package main

import "fmt"

func main() {
	var balance float64 = 5000.0
	var choice int
	var amount float64

	fmt.Println("===== Welcome to GoLang ATM =====")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	if choice == 1 {
		fmt.Println("Your Balance is:", balance)
	} else if choice == 2 {
		fmt.Print("Enter amount to deposit: ")
		fmt.Scanln(&amount)
		balance += amount
		fmt.Println("Deposited successfully! New Balance:", balance)
	} else if choice == 3 {
		fmt.Print("Enter amount to withdraw: ")
		fmt.Scanln(&amount)
		if amount <= balance {
			balance -= amount
			fmt.Println("Withdraw successful! Remaining Balance:", balance)
		} else {
			fmt.Println("Insufficient balance.")
		}
	} else if choice == 4 {
		fmt.Println("Thank you for using our ATM.")
	} else {
		fmt.Println("Invalid choice.")
	}
}
