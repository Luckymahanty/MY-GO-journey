package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("report.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString("\nThis line is appended!")
	if err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}
	fmt.Println("Text appended successfully!")
}
