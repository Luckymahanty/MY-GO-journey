package main

import "fmt"

func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1

}

func main() {
	arr := []int{10, 20, 30, 45, 34, 54, 24, 67, 75, 100}
	target := 75

	result := binarySearch(arr, target)
	if result != -1 {
		fmt.Printf("Found %d in  the index %d\n ", target, result)
	} else {
		fmt.Printf("Not found you Number %d in the array\n", target)
	}
}
