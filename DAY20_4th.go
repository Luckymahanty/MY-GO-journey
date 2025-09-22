package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Println("Hello , I am ", p.Name)

}

type Employee struct {
	Person
	Company string
	Role    string
}

func (e Employee) work() {
	fmt.Println(e.Name, "is Workin as", e.Role, "at", e.Company)

}

type Manager struct {
	Employee
	TeamSize int
}

func (m Manager) Manage() {
	fmt.Println(m.Name, "Manager a team of ", m.TeamSize)

}
func main() {
	Manager := Manager{
		Employee: Employee{
			Person:  Person{Name: "Laxmi", Age: 23},
			Company: "OpenAI",
			Role:    "Engineering Manager",
		},
		TeamSize: 5,
	}

	Manager.Greet()
	Manager.work()
	Manager.Manage()

}
