package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("report.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File contents:\n", string(data))

}
