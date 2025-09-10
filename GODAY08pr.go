package main

import "fmt"

func main() {
	transactoin := []int{500000, 3876, -34233, -232, -111, 676}
	fmt.Println("Your Transaction History:")
	balance := 0
	for i, t := range transactoin {
		balance += t
		fmt.Printf("Trancation %d: %d, Balance is : %d\n", i+1, t, balance)

	}
	fmt.Println("Now your balance is : ", balance)
}
