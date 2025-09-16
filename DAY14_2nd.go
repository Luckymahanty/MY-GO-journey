package main

import (
	"fmt"
	"sync"
	"time"
)

func deliveryOrder(orderID int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf(" Your Order %d is Out for Delivery\n", orderID)
	time.Sleep(2 * time.Second)
	fmt.Printf("Your Order %d is Delived Susecefully\n", orderID)

}
func main() {

	var wg sync.WaitGroup
	wg.Add(3)

	go deliveryOrder(1, &wg)
	go deliveryOrder(2, &wg)
	go deliveryOrder(3, &wg)

	wg.Wait()

	fmt.Println("All Order Delivered")

}
