package main

import "fmt"

func safeDrive(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from Panic:", r)
		}

	}()
	fmt.Println("Result:", a/b)

}
func main() {
	safeDrive(10, 2)
	safeDrive(10, 0)
	fmt.Println("Program Continue")

}
