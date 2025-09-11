package main

import "fmt"

type Account struct {
	balance int
}

func (a *Account) deposit(amount int) {
	a.balance += amount
}
func (a *Account) withdraw(amount int) {
	if a.balance >= amount {
		a.balance -= amount
	} else {
		fmt.Println("!you dont have isufficent money!")
	}
}
func main() {
	acc := Account{balance: 1000}
	acc.deposit(500)
	acc.withdraw(300)
	fmt.Println("Your total balance is now :", acc.balance)

}
