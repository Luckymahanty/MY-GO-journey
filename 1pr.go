package main

import "fmt"

func main() {
	phonebook := make(map[string]string)

	phonebook["Laxmi"] = "9987637823"
	phonebook["Lucky"] = "5676458758"
	phonebook["weekend"] = "4567898769"

	fmt.Println("Phonebook :", phonebook)

	fmt.Println("Lucky's phonr numb:", phonebook["Lucky"])

}
