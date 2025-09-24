package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("report.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("Read %d bytes:\n%s\n", count, string(data[:count]))

}
