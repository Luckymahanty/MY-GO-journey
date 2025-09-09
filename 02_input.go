package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name) // note: stops at space. "Laxmi" works; "Laxmi Kanta" will only read "Laxmi"

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	fmt.Println("Hello,", name, "You are", age, "years old.")
}
