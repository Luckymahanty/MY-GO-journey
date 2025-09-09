package main

import "fmt"

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
func main() {
	var a int
	fmt.Println("Enter the number you want :")
	fmt.Scanln(&a)

	fmt.Println("The of factorial was :", factorial(a))

}
