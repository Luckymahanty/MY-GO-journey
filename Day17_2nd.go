package main

import "fmt"

func main() {

	fmt.Println("Then Sentence was Before Panic use.")
	panic("This is the panic Sentence")

	fmt.Println("THIS WAS NEVER PRINT TILL PANIC WAS STAY")

}
