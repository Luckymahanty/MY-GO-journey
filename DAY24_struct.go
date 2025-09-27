package main

import (
	"fmt"
	"strings"
)

type Book struct {
	Title  string
	Author string
	Price  float64
}

func main() {

	library := []Book{
		{"Go Basic ", "Jhone", 344},
		{"DATA STRUCTURE", "Lucky", 56587},
		{"History", "Jhone deep", 460},
		{"Advance Go ", "Ram", 350},
	}
	searchTitle := "DATA STRUCTURE"
	found := false
	for _, book := range library {
		if strings.ToLower(book.Title) == strings.ToLower(searchTitle) {
			fmt.Println("Book Found:", book)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Book not Found:", searchTitle)
	}

	maxPrice := 400.0
	fmt.Println("\n Book cheaper than ", maxPrice, ":")
	for _, book := range library {
		if book.Price <= maxPrice {
			fmt.Println("-", book.Title, "by", book.Author, "Rs.", book.Price)
		}
	}

}
