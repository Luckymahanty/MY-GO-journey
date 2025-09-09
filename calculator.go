package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func multi(a, b int) int {
	return a * b

}
func divd(a, b int) int {
	return a / b

}
func main() {
	var a, b int
	fmt.Println("Enter two number you want :")
	fmt.Scanln(&a, &b)

	fmt.Println("ADDITION WAS:", add(a, b))
	fmt.Println("Substracton was:", sub(a, b))
	fmt.Println("Multiply was :", multi(a, b))
	fmt.Println("Divide was:", divd(a, b))

}
