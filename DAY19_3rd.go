package main

import "fmt"

type vehicle interface {
	Move() string
}

type Car struct{}
type Bike struct{}

func (c Car) Move() string {
	return "Car Driver"

}

func (b Bike) Move() string {
	return "Bike Racer "

}

func VehicleType(v vehicle) {
	switch v.(type) {
	case Car:
		fmt.Println("This is a car :", v.Move())

	case Bike:
		fmt.Println("This is a Bike:", v.Move())

	default:
		fmt.Println("Unknow Type ")

	}

}

func main() {

	var v1 vehicle = Car{}
	var v2 vehicle = Bike{}

	VehicleType(v1)
	VehicleType(v2)

}
