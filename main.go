package main

import "fmt"

var br1libPtrX = NewBorrower("Borrower1", 1)

func main() {
	var runes []rune
	var runes2 = []rune("Hi")
	for _, r := range "Hello" {
		runes = append(runes, r)
	}
	var borrowers []*Borrower
	var books []*Book

	AddBorrower(borrowers, br1libPtrX)
	fmt.Printf("%q\n", runes)
	fmt.Printf("%q\n", runes2)
	//AddBorrower(borrowers, NewBorrower("Sue", 3))
	AddBook(books, NewBook("War And Peace", "Tolstoy"))
	//AddBook(books, NewBook("Great Expectations", "Dickens"))
	//println("\nJust created new library")
	println(StatusToString(books, borrowers))
}
