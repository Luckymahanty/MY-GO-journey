package main

import "fmt"

func knapSackDP(W int, wt []int, val []int, n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}
	for i := 0; i <= n; i++ {
		for w := 0; w <= W; w++ {
			if i == 0 || w == 0 {
				dp[i][w] = 0
			} else if wt[i-1] <= w {
				include := val[i-1] + dp[i-1][w-wt[i-1]]
				exclude := dp[i-1][w]
				if include > exclude {
					dp[i][w] = include
				} else {
					dp[i][w] = exclude
				}

			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}
	return dp[n][W]

}
func main() {
	val := []int{60, 100, 120}
	wt := []int{10, 20, 30}
	W := 50
	n := len(val)

	fmt.Println("Max Value in knapsack(DP Tabulation):", knapSackDP(W, wt, val, n))

}
