package main

import (
	"fmt"
	"time"
)

func cookDish(dish string, ch chan string) {
	fmt.Println("Cooking", dish, ".....")
	time.Sleep(2 * time.Second)
	ch <- dish + " is ready"

}
func main() {

	ch := make(chan string)

	go cookDish("Pizza", ch)
	go cookDish("Burger", ch)
	go cookDish("Pasta", ch)

	for i := 0; i < 3; i++ {
		message := <-ch
		fmt.Println("Waiter recived:", message)
	}

	fmt.Println("_All Dishes Server_")

}
