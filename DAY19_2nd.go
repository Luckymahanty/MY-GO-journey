package main

import "fmt"

func printType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Println("This is an String ", v)

	case int:
		fmt.Println("This is an Integer", v)

	case float64:
		fmt.Println("This is an float ", v)

	default:
		fmt.Println("Unknown data type")

	}

}

func main() {
	printType("Hello")
	printType(33)
	printType(56.6)
	printType(true)

}
