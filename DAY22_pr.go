package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("Student_report.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()

	f.WriteString("Student: Lxmi\nSubject: GoLang\nMarks: 95\n")
	f2, err := os.OpenFile("student_report.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f2.Close()
	f2.WriteString("Remarks: Excellent performance!\n")

	data, err := os.ReadFile("student_report.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Report Card:\n", string(data))
}
