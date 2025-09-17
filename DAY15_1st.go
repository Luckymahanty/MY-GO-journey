package main

import (
	"fmt"
	"time"
)

func orderFromResturant(name string, delay int, ch chan string) {
	time.Sleep(time.Duration(delay) * time.Second)
	ch <- fmt.Sprintf("Order from %s is delivered in %d second", name, delay)

}
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go orderFromResturant("Dominos", 3, ch1)
	go orderFromResturant("KFC", 5, ch2)
	go orderFromResturant("Pizaa hut", 2, ch3)

	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		case msg3 := <-ch3:
			fmt.Println(msg3)
		}
	}

	fmt.Println("All Order Recived")
}
