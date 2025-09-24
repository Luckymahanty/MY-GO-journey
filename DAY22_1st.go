package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("report.txt")
	if err != nil {
		fmt.Println("Error while creating the file : ", err)
		return
	}
	defer file.Close()

	text := "Laxmi's Go Learning Report\nDay 22: File Handling"
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Error writing to file :", err)
		return
	}
	fmt.Println("File written successfully")
}
