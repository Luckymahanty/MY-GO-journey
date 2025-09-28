package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		miniIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[miniIndex] {
				miniIndex = j
			}
		}
		arr[i], arr[miniIndex] = arr[miniIndex], arr[i]
	}

}
func main() {
	arr := []int{29, 34, 5, 355, 54, 76, 9, 45}
	fmt.Println("Before Selection Sort Use:", arr)
	selectionSort(arr)
	fmt.Println("After Selction Sort  Use :", arr)

}
