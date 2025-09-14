package main

import (
	"fmt"

	"time"
)

func deliverFood(orderId int, timeToDelivery int) {
	fmt.Printf("Your order %d is prepared...\n", orderId)
	time.Sleep(time.Duration(timeToDelivery) * time.Second)
	fmt.Printf("Your Order %d is Deliver in %d second \n", orderId, timeToDelivery)

}
func main() {
	fmt.Println("FOOD DELIVERY BY LUCKY")

	go deliverFood(1, 3)
	go deliverFood(2, 5)
	go deliverFood(4, 6)

	time.Sleep(6 * time.Second)

	fmt.Println("ALL Order Deliverd")

}
