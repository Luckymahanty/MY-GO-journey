package main

import "fmt"

type Student struct {
	name  string
	age   int
	grade string
}

func main() {
	s1 := Student{name: "lucky", age: 21, grade: "B"}
	fmt.Println("Name:", s1.name)
	fmt.Println("AGE:", s1.age)
	fmt.Println("Grade:", s1.grade)

}
