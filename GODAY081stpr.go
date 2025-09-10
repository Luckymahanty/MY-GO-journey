package main

import "fmt"

func main() {
	fmt.Println("\n RANGE USE IN ARRAY:\n")
	numbers := [5]int{12, 34, 23, 54, 65}
	for i, val := range numbers {
		fmt.Println("Index:", i, "Value:", val)
	}
	fmt.Println("\n RANGE USE IN SLICE:\n")
	fruits := []string{"apple", "banana", "chery", "carrot"}
	for j, fru := range fruits {
		fmt.Println("Cart number:", j, "is", fru)

	}
	fmt.Println("\n RANGE USE IN MAP:\n")
	Phonebook := map[string]string{
		"Lucky": "567654567",
		"vicky": "343445455",
	}
	for name, number := range Phonebook {

		fmt.Println(name, ":", number)
	}
	fmt.Println("\n NOW FOR STRING:\n")

	Word := "GOLANG"
	for i, ch := range Word {
		fmt.Printf("index: %d ,  Character: %c\n", i, ch)
	}

}
