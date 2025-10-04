package main

import "fmt"

func knapSackMemo(W int, wt []int, val []int, n int, memo map[[2]int]int) int {
	if n == 0 || W == 0 {
		return 0
	}
	if v, ok := memo[[2]int{n, W}]; ok {
		return v
	}
	if wt[n-1] > W {
		memo[[2]int{n, W}] = knapSackMemo(W, wt, val, n-1, memo)
	} else {
		include := val[n-1] + knapSackMemo(W-wt[n-1], wt, val, n-1, memo)
		exclude := knapSackMemo(W, wt, val, n-1, memo)
		if include > exclude {
			memo[[2]int{n, W}] = include
		} else {
			memo[[2]int{n, W}] = exclude
		}
	}
	return memo[[2]int{n, W}]

}
func main() {
	val := []int{60, 100, 120}
	wt := []int{10, 20, 30}
	W := 50
	n := len(val)

	memo := make(map[[2]int]int)
	fmt.Println("Max Value in knapsack(Memoization):", knapSackMemo(W, wt, val, n, memo))

}
