package main

import "fmt"

type shape interface {
	area() float64
}
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius

}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height

}
func main() {
	var s shape
	s = Circle{radius: 5}
	fmt.Println("Circle  Area :", s.area())

	s = Rectangle{width: 5, height: 10}
	fmt.Println("Area of Rectangle is : ", s.area())

}
