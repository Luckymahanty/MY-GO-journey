package main

import "fmt"

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}
func main() {
	c1 := circle{radius: 5}
	fmt.Println("Area of an circle is :", c1.area())
}
