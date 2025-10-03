package main

import "fmt"

func fibTab(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]

}

func main() {
	fmt.Println("Fibonacci of 10(DP Tabulation ):", fibTab(10))

}
