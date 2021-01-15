package main

import (
	"eatobin.com/totalbeginnergo/book"
	"eatobin.com/totalbeginnergo/borrower"
	"eatobin.com/totalbeginnergo/library"
	"fmt"
	"io/ioutil"
	"os"
)

var borrowers []*borrower.Borrower
var books []*book.Book
var jsonBorrowersFileBefore = "resources/borrowers-before.json"
var jsonBooksFile = "resources/books-before.json"

var jsonBorrowersFileAfter = "resources/borrowers-after.json"

var jsonBorrowersFileBad = "resources/bad-borrowers.json"

var emptyFile = "resources/resources/empty.json"

func main() {
	borrowers = library.AddBorrower(borrowers, borrower.NewBorrower("Jim", 3))
	borrowers = library.AddBorrower(borrowers, borrower.NewBorrower("Sue", 3))
	books = library.AddBook(books, book.NewBook("War And Peace", "Tolstoy"))
	books = library.AddBook(books, book.NewBook("Great Expectations", "Dickens"))
	fmt.Println("\nJust created new library")
	fmt.Println(library.StatusToString(books, borrowers))

	fmt.Println("Check out War And Peace to Sue")
	books = library.CheckOut("Sue", "War And Peace", borrowers, books)
	fmt.Println(library.StatusToString(books, borrowers))

	fmt.Println("Now check in War And Peace from Sue...")
	books = library.CheckIn("War And Peace", books)
	fmt.Println("...and check out Great Expectations to Jim")
	books = library.CheckOut("Jim", "Great Expectations", borrowers, books)
	fmt.Println(library.StatusToString(books, borrowers))

	fmt.Println("Add Eric and The Cat In The Hat")
	borrowers = library.AddBorrower(borrowers, borrower.NewBorrower("Eric", 1))
	books = library.AddBook(books, book.NewBook("The Cat In The Hat", "Dr. Seuss"))
	fmt.Println("Check Out Dr. Seuss to Eric")
	books = library.CheckOut("Eric", "The Cat In The Hat", borrowers, books)
	fmt.Println(library.StatusToString(books, borrowers))

	fmt.Println("Now let's do some BAD stuff...")

	fmt.Println("Add a borrower that already exists (Borrower('Jim', 3))")
	borrowers = library.AddBorrower(borrowers, borrower.NewBorrower("Jim", 3))
	fmt.Println("No change to Test Library:")
	fmt.Println(library.StatusToString(books, borrowers))

	fmt.Println("Add a book that already exists (Book('War And Peace', 'Tolstoy'))")
	books = library.AddBook(books, book.NewBook("War And Peace", "Tolstoy"))
	fmt.Println("No change to Test Library:")
	fmt.Println(library.StatusToString(books, borrowers))

	fmt.Println("Check out a valid book to an invalid person (checkOut('JoJo', 'War And Peace'))")
	books = library.CheckOut("JoJo", "War And Peace", borrowers, books)
	fmt.Println("No change to Test Library:")
	fmt.Println(library.StatusToString(books, borrowers))

	fmt.Println("Check out an invalid book to an valid person (checkOut('Sue', 'Not A Book'))")
	books = library.CheckOut("Sue", "Not A Book", borrowers, books)
	fmt.Println("No change to Test Library:")
	fmt.Println(library.StatusToString(books, borrowers))

	fmt.Println("Last - check in a book not checked out (checkIn('War And Peace'))")
	books = library.CheckIn("War And Peace", books)
	fmt.Println("No change to Test Library:")
	fmt.Println(library.StatusToString(books, borrowers))

	fmt.Println("Okay... let's finish with some persistence. First clear the whole library:")
	newEmptyV()

	fmt.Println("Lets read in a new library from \"borrowers-before.json\" and \"books-before.json\":")
	newVError := newV(jsonBorrowersFileBefore, jsonBooksFile)
	if newVError != nil {
		panic(newVError)
	}
	fmt.Println("Add... a new borrower:")
	borrowers = library.AddBorrower(borrowers, borrower.NewBorrower("BorrowerNew", 300))
	fmt.Println(library.StatusToString(books, borrowers))

	fmt.Println("Save the revised borrowers to \"borrowers-after.json\"")
	var jsonBrsStr, _ = library.BorrowersToJSONSting(borrowers)
	writeJsonStringToFile(jsonBrsStr)

	fmt.Println("Clear the whole library again:")
	newEmptyV()

	fmt.Println("Then read in the revised library from \"borrowers-after.json\" and \"books-before.json\":")
	newVError = newV(jsonBorrowersFileAfter, jsonBooksFile)
	if newVError != nil {
		panic(newVError)
	}

	fmt.Println("Last... delete the file \"borrowers-after.json\"")
	err := os.Remove(jsonBorrowersFileAfter)
	if err != nil {
		panic(newVError)
	}
	newEmptyV()

	fmt.Println("Then try to make a library using the deleted \"borrowers-after.json\" and \"books-before.json\":")
	newVError = newV(jsonBorrowersFileAfter, jsonBooksFile)
	if newVError != nil {
		fmt.Println(newVError.Error())
	}

	fmt.Println("\nAnd if we read in a file with mal-formed json content... like \"bad-borrowers.json\" and \"books-before.json\":")
	newVError = newV(jsonBorrowersFileBad, jsonBooksFile)
	if newVError != nil {
		fmt.Println(newVError.Error())
	}

	fmt.Println("\nOr how about reading in an empty file... \"empty.json\" (for borrowers and books):")
	newVError = newV(emptyFile, emptyFile)
	if newVError != nil {
		fmt.Println(newVError.Error())
	}

	fmt.Println("\nAnd... that's all...")
	fmt.Println("Thanks - bye!")
}

func newEmptyV() {
	borrowers = []*borrower.Borrower{}
	books = []*book.Book{}
	fmt.Println(library.StatusToString(books, borrowers))
}

func readFileIntoJsonString(fp string) (string, error) {
	dat, err := ioutil.ReadFile(fp)
	return string(dat), err
}

func writeJsonStringToFile(js string) {
	var file = "resources/borrowers-after.json"
	err := ioutil.WriteFile(file, []byte(js), 0644)
	if err != nil {
		panic(err)
	}
}

func newV(brsFp string, bksFp string) error {
	jsonBrsStr, brFileErr := readFileIntoJsonString(brsFp)
	if brFileErr != nil {
		return brFileErr
	}
	jsonBksStr, bkFileErr := readFileIntoJsonString(bksFp)
	if bkFileErr != nil {
		return bkFileErr
	}
	brs, brParseErr := library.JsonStringToBorrowers(jsonBrsStr)
	if brParseErr != nil {
		return brParseErr
	}
	bks, bkParseErr := library.JsonStringToBooks(jsonBksStr)
	if bkParseErr != nil {
		return bkParseErr
	}
	borrowers = brs
	books = bks
	fmt.Println(library.StatusToString(books, borrowers))
	return nil
}
