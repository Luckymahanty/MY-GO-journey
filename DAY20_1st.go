package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	Company string
}

func main() {

	e := Employee{
		Person:  Person{Name: "Lucky", Age: 21},
		Company: "Google",
	}

	fmt.Println("Empolyee Name is :", e.Name)
	fmt.Println("Empolyee Age is :", e.Age)
	fmt.Println("Company Name is :", e.Company)

}
