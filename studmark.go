package main

import "fmt"

type Student struct {
	name  string
	roll  int
	marks int
}

func updateMarks(s *Student, newMarks int) {
	s.marks = newMarks

}
func main() {
	s1 := Student{name: "Laxmikanta", roll: 25, marks: 100}
	fmt.Println("Before:", s1)

	updateMarks(&s1, 56)
	fmt.Println("After update", s1)
}
