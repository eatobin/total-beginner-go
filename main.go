package main

import (
	"eatobin.com/totalbeginnergo/book"
	"eatobin.com/totalbeginnergo/borrower"
	"eatobin.com/totalbeginnergo/library"
	"fmt"
)

func main() {
	var borrowers []borrower.Borrower
	var books []book.Book

	borrowers = library.AddBorrower(borrowers, borrower.NewBorrower("Jim", 3))
	borrowers = library.AddBorrower(borrowers, borrower.NewBorrower("Sue", 3))
	books = library.AddBook(books, book.NewBook("War And Peace", "Tolstoy"))
	books = library.AddBook(books, book.NewBook("Great Expectations", "Dickens"))
	fmt.Println("\nJust created new library")
	fmt.Println(library.StatusToString(books, borrowers))

	fmt.Println("Check out War And Peace to Sue")
	books = library.CheckOut("Sue", "War And Peace", borrowers, books)
	fmt.Println(library.StatusToString(books, borrowers))
}
