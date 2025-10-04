package main

import "fmt"

func knapSackRecursive(W int, wt []int, val []int, n int) int {
	if n == 0 || W == 0 {
		return 0
	}
	if wt[n-1] > W {
		return knapSackRecursive(W, wt, val, n-1)
	}

	include := val[n-1] + knapSackRecursive(W-wt[n-1], wt, val, n-1)
	exclude := knapSackRecursive(W, wt, val, n-1)

	if include > exclude {
		return include
	}
	return exclude
}

func main() {
	val := []int{60, 100, 120}
	wt := []int{10, 20, 30}
	W := 50
	n := len(val)

	fmt.Println("Max value in knapsack(Recursive):", knapSackRecursive(W, wt, val, n))

}
