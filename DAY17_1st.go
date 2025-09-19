package main

import "fmt"

func main() {
	fmt.Println("START")

	defer fmt.Println("Clean up started")
	defer fmt.Println("Clean up Complited")

	fmt.Println("END")
}
