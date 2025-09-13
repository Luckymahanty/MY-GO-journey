package main

import (
	"fmt"
)

func getGrade(score int) (string, error) {
	if score < 0 || score > 100 {

		return "", fmt.Errorf("Invalid score : %d ", score)

	}
	if score >= 50 {
		return "pass", nil
	}
	return "Fail", nil
}
func main() {
	grade, err := getGrade(126)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Your :", grade)
	}

}
