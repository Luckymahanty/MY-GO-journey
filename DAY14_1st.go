package main

import (
	"fmt"
	"time"
)

func pizzaMaking(order string, ch chan string) {
	fmt.Println("Cooking :", order)
	time.Sleep(2 * time.Second)
	ch <- order + " is ready"

}
func main() {
	ch := make(chan string, 3)

	pizzaMaking("Margerita ", ch)
	pizzaMaking("Farmhouse", ch)
	pizzaMaking("Peproni", ch)
	for i := 0; i < 3; i++ {
		fmt.Println("Waiter got the ", <-ch)

	}

}
