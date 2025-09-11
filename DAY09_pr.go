package main

import "fmt"

type Book struct {
	tittle string
	author string
	stock  int
}

func (b *Book) borrow() {
	if b.stock > 0 {
		b.stock--
		fmt.Println("Borrowed:", b.tittle)
	} else {
		fmt.Println("BOOK ", b.tittle, "is Out of Stock")
	}

}

func (b *Book) returnBook() {
	b.stock++
	fmt.Println("Book is return: ", b.tittle)
}
func (b Book) info() {
	fmt.Printf("%s by %s (avalable: %d)\n", b.tittle, b.author, b.stock)

}
func main() {
	book1 := Book{tittle: "GO 0 to 100", author: "Lucky", stock: 4}

	book1.info()
	book1.borrow()
	book1.borrow()
	book1.borrow()
	book1.info()
	book1.borrow()
	book1.returnBook()
	book1.info()

}
