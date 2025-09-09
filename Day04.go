package main

import "fmt"

func main() {
	groceries := []string{}

	groceries = append(groceries, "Milk")
	groceries = append(groceries, "butter")
	groceries = append(groceries, "Salt")
	groceries = append(groceries, "Sugar")

	fmt.Println("Groceries List")
	for i := 0; i < len(groceries); i++ {
		fmt.Println(i+1, "-", groceries[i])
	}

}
