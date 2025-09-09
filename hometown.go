package main

import "fmt"

func main() {
	var name string
	var city string
	fmt.Printf("Enter your name bro/sis: ")
	fmt.Scanln(&name)

	fmt.Printf("Enter your Hometown:")
	fmt.Scanln(&city)

	fmt.Printf("Helllo %s so your from %s! Nice.", name, city)

}
