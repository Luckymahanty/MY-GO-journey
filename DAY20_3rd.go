package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Println("HII  I AM ", p.Name)

}

type Employee struct {
	Person
	Company string
}

func (e Employee) Greet() {
	fmt.Println("HII , I am ", e.Name, "and i work at ", e.Company)

}

func main() {

	e := Employee{Person: Person{Name: "LUCKY"}, Company: "Microsoft"}
	e.Greet()

}
