package main

import "fmt"

func main() {
	sci := 67
	math := 79
	english := 89
	history := 45

	total := sci + math + english + history

	totalper := float64(total) / 400 * 100
	fmt.Println(" your total percentage is  ", totalper, " 🚀")
}
