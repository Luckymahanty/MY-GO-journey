package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Println("Hello , Iam ", p.Name)

}

type Employee struct {
	Person
	Company string
}

func main() {

	e := Employee{Person: Person{Name: "Lucky"}, Company: "CLOUD_CS"}
	e.Greet()

}
