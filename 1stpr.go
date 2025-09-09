package main

import "fmt"

func main() {
	x := 10
	p := &x

	fmt.Println("Print the value of x:", x)
	fmt.Println("Print the adress of x:", p)
	fmt.Println("print the point value of x;", *p)

}
