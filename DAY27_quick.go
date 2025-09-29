package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[len(arr)/2]
	left, right := []int{}, []int{}

	for i, val := range arr {
		if i == len(arr)/2 {
			continue
		}
		if val <= pivot {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}
	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, pivot), right...)
}
func main() {
	arr := []int{18, 198, 1, 76, 565, 45, 0, 89, 56765, 46}
	fmt.Println("Before using the sorting:", arr)
	sorted := quickSort(arr)
	fmt.Println("After using the sorting :", sorted)

}
