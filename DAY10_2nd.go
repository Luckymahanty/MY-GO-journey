package main

import "fmt"

type Payment interface {
	pay(amount float64)
}
type Creditcard struct {
	cardNumber string
}

func (c Creditcard) pay(amount float64) {
	fmt.Printf("Paid %.2f using Creditcard (%s)\n", amount, c.cardNumber)

}

type Paypal struct {
	email string
}

func (p Paypal) pay(amount float64) {
	fmt.Printf("Paid %.2f using paypal (%s)\n", amount, p.email)

}
func main() {
	var p Payment
	p = Creditcard{cardNumber: "5678765434567"}
	p.pay(4545.23)

	p = Paypal{email: "hjkjhjoi@gmail.com"}
	p.pay(56765.567)

}
