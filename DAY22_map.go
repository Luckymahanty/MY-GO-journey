package main

import (
	"fmt"
)

func main() {

	Cart := make(map[string]int)

	Cart["Pc"] = 1
	Cart["Laptop"] = 3
	Cart["HeadPhone"] = 4

	fmt.Println("First Cart is:", Cart)

	Cart["Laptop"] = 4
	fmt.Println("Upated Cart is:", Cart)

	delete(Cart, "PC")
	fmt.Println("Now after removing the Iteam:", Cart)

	if qty, exists := Cart["Laptop"]; exists {
		fmt.Println("Laptop found quantity is :", qty)
	} else {
		fmt.Println("Not found Laptop")
	}

	fmt.Println("The final Cart is :")
	for product, qty := range Cart {
		fmt.Printf("%s -> %d pcs\n", product, qty)
	}

}
