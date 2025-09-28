package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}

}

func main() {
	arr := []int{545, 657, 43, 3, 1, 65, 45, 987, 45, 6, 65}
	fmt.Println("Before the Insertion Sort Use :", arr)
	insertionSort(arr)
	fmt.Println("After the Insetion Sort is Use:", arr)

}
