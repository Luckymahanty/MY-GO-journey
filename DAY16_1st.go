package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Timer Started")
	timer := time.NewTimer(20 * time.Second)
	<-timer.C
	fmt.Println("Timer finshed ")

}
