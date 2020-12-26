package main

import "fmt"

func main() {
	var borrowers []*Borrower
	var books []*Book

	borrowers = AddBorrower(borrowers, NewBorrower("Jim", 3))
	borrowers = AddBorrower(borrowers, NewBorrower("Sue", 3))
	books = AddBook(books, NewBook("War And Peace", "Tolstoy"))
	books = AddBook(books, NewBook("Great Expectations", "Dickens"))
	fmt.Println("\nJust created new library")
	fmt.Println(StatusToString(books, borrowers))

	fmt.Println("Check out War And Peace to Sue")
	books = CheckOut("Sue", "War And Peace", borrowers, books)
	fmt.Println(StatusToString(books, borrowers))

	fmt.Println("Now check in War And Peace from Sue...")
	books = CheckIn("War And Peace", books)
	fmt.Println("...and check out Great Expectations to Jim")
	books = CheckOut("Jim", "Great Expectations", borrowers, books)
	fmt.Println(StatusToString(books, borrowers))

	fmt.Println("Add Eric and The Cat In The Hat")
	borrowers = AddBorrower(borrowers, NewBorrower("Eric", 1))
	books = AddBook(books, NewBook("The Cat In The Hat", "Dr. Seuss"))
	fmt.Println("Check Out Dr. Seuss to Eric")
	books = CheckOut("Eric", "The Cat In The Hat", borrowers, books)
	fmt.Println(StatusToString(books, borrowers))

	fmt.Println("Now let's do some BAD stuff...")

	fmt.Println("Add a borrower that already exists (Borrower('Jim', 3))")
	borrowers = AddBorrower(borrowers, NewBorrower("Jim", 3))
	fmt.Println("No change to Test Library:")
	fmt.Println(StatusToString(books, borrowers))

	fmt.Println("Add a book that already exists (Book('War And Peace', 'Tolstoy'))")
	books = AddBook(books, NewBook("War And Peace", "Tolstoy"))
	fmt.Println("No change to Test Library:")
	fmt.Println(StatusToString(books, borrowers))
}
