package main

import "fmt"

func fibMemo(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val
	}
	memo[n] = fibMemo(n-1, memo) + fibMemo(n-2, memo)
	return memo[n]
}
func main() {
	memo := make(map[int]int)
	fmt.Println("Fibonacci of 10:", fibMemo(10, memo))

}
