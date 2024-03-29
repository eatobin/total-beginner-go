package main

import (
	"fmt"
	"github.com/eatobin/total-beginner-go/book"
	"github.com/eatobin/total-beginner-go/borrower"
	"github.com/eatobin/total-beginner-go/library"
	"io/ioutil"
	"os"
)

type Borrower = borrower.Borrower
type Book = book.Book

var borrowers library.Borrowers
var books library.Books

var jsonBorrowersFileBefore = "resources/borrowers-before.json"
var jsonBooksFile = "resources/books-before.json"
var jsonBorrowersFileAfter = "resources/borrowers-after.json"
var jsonBorrowersFileBad = "resources/bad-borrowers.json"
var emptyFile = "resources/resources/empty.json"

func main() {
	borrowers = library.AddBorrower(borrowers, Borrower{Name: "Jim", MaxBooks: 3})
	borrowers = library.AddBorrower(borrowers, Borrower{Name: "Sue", MaxBooks: 3})
	books = library.AddBook(books, Book{Title: "War And Peace", Author: "Tolstoy"})
	books = library.AddBook(books, Book{Title: "Great Expectations", Author: "Dickens"})
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
	borrowers = library.AddBorrower(borrowers, Borrower{Name: "Eric", MaxBooks: 1})
	books = library.AddBook(books, Book{Title: "The Cat In The Hat", Author: "Dr. Seuss"})
	fmt.Println("Check Out Dr. Seuss to Eric")
	books = library.CheckOut("Eric", "The Cat In The Hat", borrowers, books)
	fmt.Println(library.StatusToString(books, borrowers))

	fmt.Println("Now let's do some BAD stuff...")

	fmt.Println("Add a borrower that already exists (total.Borrower('Jim', 3))")
	borrowers = library.AddBorrower(borrowers, Borrower{Name: "Jim", MaxBooks: 3})
	fmt.Println("No change to Test Library:")
	fmt.Println(library.StatusToString(books, borrowers))

	fmt.Println("Add a book that already exists (total.Book('War And Peace', 'Tolstoy', None))")
	books = library.AddBook(books, Book{Title: "War And Peace", Author: "Tolstoy"})
	fmt.Println("No change to Test Library:")
	fmt.Println(library.StatusToString(books, borrowers))

	fmt.Println("Check out a valid book to an invalid person (checkOut('JoJo', 'War And Peace', borrowers))")
	books = library.CheckOut("JoJo", "War And Peace", borrowers, books)
	fmt.Println("No change to Test Library:")
	fmt.Println(library.StatusToString(books, borrowers))

	fmt.Println("Check out an invalid book to an valid person (checkOut('Sue', 'Not A total.Book', borrowers))")
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
	borrowers, books = newV(jsonBorrowersFileBefore, jsonBooksFile)
	fmt.Println("Add... a new borrower:")
	borrowers = library.AddBorrower(borrowers, Borrower{Name: "BorrowerNew", MaxBooks: 300})
	fmt.Println(library.StatusToString(books, borrowers))

	fmt.Println("Save the revised borrowers to \"borrowers-after.json\"")
	jsonBrsStr, _ := library.BorrowersToJSONSting(borrowers)
	_ = writeJsonStringToFile("resources/borrowers-after.json", jsonBrsStr)

	fmt.Println("Clear the whole library again:")
	newEmptyV()

	fmt.Println("Then read in the revised library from \"borrowers-after.json\" and \"books-before.json\":")
	borrowers, books = newV(jsonBorrowersFileAfter, jsonBooksFile)

	fmt.Println("Last... delete the file \"borrowers-after.json\"")
	_ = os.Remove(jsonBorrowersFileAfter)
	newEmptyV()

	fmt.Println("Then try to make a library using the deleted \"borrowers-after.json\" and \"books-before.json\":")
	borrowers, books = newV(jsonBorrowersFileAfter, jsonBooksFile)

	fmt.Println("\nAnd if we read in a file with mal-formed json content... like \"bad-borrowers.json\" and \"books-before.json\":")
	borrowers, books = newV(jsonBorrowersFileBad, jsonBooksFile)

	fmt.Println("\nOr how about reading in an empty file... \"empty.json\" (for borrowers and books):")
	borrowers, books = newV(emptyFile, emptyFile)

	fmt.Println("\nAnd... that's all...")
	fmt.Println("Thanks - bye!")
}

func newEmptyV() {
	borrowers = library.ZeroBorrowers
	books = library.ZeroBooks
	fmt.Println(library.StatusToString(books, borrowers))
}

func readFileIntoJsonString(fp string) (string, error) {
	jsonStr, err := ioutil.ReadFile(fp)
	return string(jsonStr), err
}

func writeJsonStringToFile(fp string, js string) error {
	f, createErr := os.Create(fp)
	if createErr != nil {
		return createErr
	}
	_, writeErr := f.WriteString(js)
	if writeErr != nil {
		return writeErr
	}
	return nil
}

func newV(brsfp string, bksfp string) (library.Borrowers, library.Books) {
	jsonBrsStr, brFileErr := readFileIntoJsonString(brsfp)
	if brFileErr != nil {
		fmt.Println(brFileErr.Error())
		return library.ZeroBorrowers, library.ZeroBooks
	}
	jsonBksStr, bkFileErr := readFileIntoJsonString(bksfp)
	if bkFileErr != nil {
		fmt.Println(bkFileErr.Error())
		return library.ZeroBorrowers, library.ZeroBooks
	}
	brs, brParseErr := library.JsonStringToBorrowers(jsonBrsStr)
	if brParseErr != nil {
		fmt.Println(brParseErr.Error())
		return library.ZeroBorrowers, library.ZeroBooks
	}
	bks, bkParseErr := library.JsonStringToBooks(jsonBksStr)
	if bkParseErr != nil {
		fmt.Println(bkParseErr.Error())
		return library.ZeroBorrowers, library.ZeroBooks
	}
	return brs, bks
}
