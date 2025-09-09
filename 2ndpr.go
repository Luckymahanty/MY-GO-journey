package main

import "fmt"

func main() {
	Phonebook := map[string]string{

		"Lucky": "456789876234",
		"AMit":  "234567987456",
	}

	Phonebook["Lucky"] = "11111111111"

	delete(Phonebook, "AMit")

	fmt.Println("The new Phonebook is :", Phonebook)

}
