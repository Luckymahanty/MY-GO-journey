package main

import "fmt"

func main() {
	phonebook := map[string]string{
		"Lucky":      "54637823476",
		"Laxmikanta": "6758433756",
		"Subrat":     "78987898789",
	}
	fmt.Println(phonebook)
	for name, number := range phonebook {
		fmt.Println(name, ":", number)
	}
	searchName := "Subrat"
	if num, found := phonebook[searchName]; found {
		fmt.Println(searchName, "'s number is:", num)

	} else {
		fmt.Println(searchName, "Not found any number")
	}
}
