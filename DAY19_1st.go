package main

import "fmt"

func main() {
	var i interface{} = "HELLO LUCKY"

	str, ok := i.(string)
	if ok {
		fmt.Println("THE VALUE IS :", str)
	} else {
		fmt.Println("THIS IS NOT AN STRING")
	}

	num := i.(int)
	fmt.Println(num)
}
