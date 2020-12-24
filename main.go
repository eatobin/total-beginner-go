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

	println(StatusToString(books, borrowers))
	println(StatusToString(books, borrowers))
	println(StatusToString(books, borrowers))
	//println(StatusToString(books, borrowers))

	//println("Check out War And Peace to Sue")
	//books = CheckOut("Sue", "War And Peace", borrowers, books)
	//println(StatusToString(books, borrowers))
	//
	//println("Now check in War And Peace from Sue...")
	//books = CheckIn("War And Peace", books)
	//println("...and check out Great Expectations to Jim")
	//books = CheckOut("Jim", "Great Expectations", borrowers, books)
	//println(StatusToString(books, borrowers))
	//
	////println("Add Eric and The Cat In The Hat")
	//borrowers = AddBorrower(borrowers, NewBorrower("Eric", 1))
	//books = AddBook(books, NewBook("The Cat In The Hat", "Dr. Seuss"))
	////println("Check Out Dr. Seuss to Eric")
	////books = CheckOut("Eric", "The Cat In The Hat", borrowers, books)
	//println(StatusToString(books, borrowers))
}
