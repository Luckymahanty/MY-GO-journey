package main

import (
	"fmt"
)

func printAnything(i interface{}) {
	fmt.Println(i)

}

func main() {
	printAnything("Hello")

	printAnything("user 123432")

}
