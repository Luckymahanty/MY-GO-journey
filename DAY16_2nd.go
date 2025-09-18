package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(2 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at:", t)
			}
		}

	}()

	time.Sleep(7 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker Stopped")

}
