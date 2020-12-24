package main

func main() {
	var borrowers []*Borrower
	var books []*Book

	borrowers = AddBorrower(borrowers, NewBorrower("Jim", 3))
	borrowers = AddBorrower(borrowers, NewBorrower("Sue", 3))
	books = AddBook(books, NewBook("War And Peace", "Tolstoy"))
	books = AddBook(books, NewBook("Great Expectations", "Dickens"))
	println("\nJust created new library")
	println(StatusToString(books, borrowers))

	println("Check out War And Peace to Sue")
	CheckOut("Sue", "War And Peace", borrowers, books)
	println(StatusToString(books, borrowers))

	println("Now check in War And Peace from Sue...")
	CheckIn("War And Peace", books)
	println("...and check out Great Expectations to Jim")
	CheckOut("Jim", "Great Expectations", borrowers, books)
	println(StatusToString(books, borrowers))
}
