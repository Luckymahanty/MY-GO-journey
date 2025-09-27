package main

import "fmt"

func linearSearch(arr []int, target int) int {
	for i, num := range arr {
		if num == target {
			return i
		}
	}
	return -1

}

func main() {
	arr := []int{10, 20, 30, 25, 40, 35, 50, 55, 56, 57}
	target := 50

	result := linearSearch(arr, target)
	if result != -1 {
		fmt.Printf("Found the %d  at the index %d\n", target, result)
	} else {
		fmt.Printf("%d not found in the Array\n", target)
	}

}
