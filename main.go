package main

func main() {
	var borrowers []Borrower
	var books []Book

	borrowers = AddBorrower(borrowers, NewBorrower("Jim", 3))
	borrowers = AddBorrower(borrowers, NewBorrower("Sue", 3))
	books = AddBook(books, NewBook("War And Peace", "Tolstoy"))
	books = AddBook(books, NewBook("Great Expectations", "Dickens"))
	println("\nJust created new library")
	println(StatusToString(books, borrowers))
}
