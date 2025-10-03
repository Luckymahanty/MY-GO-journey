package main

import "fmt"

func fibOptimized(n int) int {
	if n <= 1 {
		return n
	}
	prev2, prev1 := 0, 1
	for i := 2; i <= n; i++ {
		curr := prev1 + prev2
		prev2, prev1 = prev1, curr
	}
	return prev1

}
func main() {
	fmt.Println("Fibonacci of 50:", fibOptimized(50))

}
