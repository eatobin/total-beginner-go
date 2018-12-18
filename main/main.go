package main

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/eatobin/totalbeginnergo/book"
	"github.com/eatobin/totalbeginnergo/borrower"
	"github.com/eatobin/totalbeginnergo/library"
)

var tvBorrowers []borrower.Borrower
var tvBooks []book.Book

var jsonBorrowersFileBefore = "borrowers-before.json"
var jsonBooksFile = "books-before.json"
var jsonBorrowersFileAfter = "borrowers-after.json"

var jsonBorrowersFileBad = "bad-borrowers.json"
var emptyFile = "empty.json"

func main() {
	tvBorrowers = library.AddBorrower(tvBorrowers, borrower.NewBorrower("Jim", 3))
	tvBorrowers = library.AddBorrower(tvBorrowers, borrower.NewBorrower("Sue", 3))
	tvBooks = library.AddBook(tvBooks, book.NewBook("War And Peace", "Tolstoy"))
	tvBooks = library.AddBook(tvBooks, book.NewBook("Great Expectations", "Dickens"))
	println("\nJust created new library:")
	println(library.StatusToString(tvBooks, tvBorrowers))

	println("Check out War And Peace to Sue:")
	tvBooks = library.CheckOut("Sue", "War And Peace", tvBorrowers, tvBooks)
	println(library.StatusToString(tvBooks, tvBorrowers))

	println("Now check in War And Peace from Sue...")
	tvBooks = library.CheckIn("War And Peace", tvBooks)
	println("...and check out Great Expectations to Jim:")
	tvBooks = library.CheckOut("Jim", "Great Expectations", tvBorrowers, tvBooks)
	println(library.StatusToString(tvBooks, tvBorrowers))

	println("Add Eric and The Cat In The Hat\nand")
	tvBorrowers = library.AddBorrower(tvBorrowers, borrower.NewBorrower("Eric", 1))
	tvBooks = library.AddBook(tvBooks, book.NewBook("The Cat In The Hat", "Dr. Seuss"))
	println("Check Out Dr. Seuss to Eric:")
	tvBooks = library.CheckOut("Eric", "The Cat In The Hat", tvBorrowers, tvBooks)
	println(library.StatusToString(tvBooks, tvBorrowers))

	println("Now let's do some BAD stuff...\n")

	println("Add a borrower that already exists (borrower.Borrower{\"Jim\", 3})")
	tvBorrowers = library.AddBorrower(tvBorrowers, borrower.NewBorrower("Jim", 3))
	println("No change to Test Library:")
	println(library.StatusToString(tvBooks, tvBorrowers))

	println("Add a book that already exists (book.Book{\"War And Peace\", \"Tolstoy\"})")
	tvBooks = library.AddBook(tvBooks, book.NewBook("War And Peace", "Tolstoy"))
	println("No change to Test Library:")
	println(library.StatusToString(tvBooks, tvBorrowers))

	println("Check out a valid book to an invalid person (CheckOut('JoJo', 'War And Peace', borrowers, books))")
	tvBooks = library.CheckOut("JoJo", "War And Peace", tvBorrowers, tvBooks)
	println("No change to Test Library:")
	println(library.StatusToString(tvBooks, tvBorrowers))

	println("Check out an invalid book to an valid person (CheckOut(\"Sue\", \"Not A Book\", borrowers, books))")
	tvBooks = library.CheckOut("Sue", "Not A Book", tvBorrowers, tvBooks)
	println("No change to Test Library:")
	println(library.StatusToString(tvBooks, tvBorrowers))

	println("Last - check in a book not checked out (CheckIn(\"War And Peace\", books))")
	tvBooks = library.CheckIn("War And Peace", tvBooks)
	println("No change to Test Library:")
	println(library.StatusToString(tvBooks, tvBorrowers))

	println("Okay... let's finish with some persistence. First clear the whole library:")
	newEmpty()

	println("Lets read in a new library from \"borrowers-before.json\" and \"books-before.json\":")
	err := newV(jsonBorrowersFileBefore, jsonBooksFile)
	if err != nil {
		println(err.Error())
	}
	println("Add... a new borrower:")
	tvBorrowers = library.AddBorrower(tvBorrowers, borrower.NewBorrower("BorrowerNew", 300))
	println(library.StatusToString(tvBooks, tvBorrowers))

	println("Save the revised borrowers to \"borrowers-after.json\":")
	nBrsJ := library.BorrowersToJSONSting(tvBorrowers)
	WriteJSONStringToFile(nBrsJ, jsonBorrowersFileAfter)

	println("Clear the whole library again:")
	newEmpty()

	println("Then read in the revised library from \"borrowers-after.json\" and \"books-before.json\":")
	err = newV(jsonBorrowersFileAfter, jsonBooksFile)
	if err != nil {
		println(err.Error())
	}

	println("Last... delete the file \"borrowers-after.json\" and clear the library:")
	err = os.Remove(jsonBorrowersFileAfter)
	if err != nil {
		println(err.Error())
		return
	}
	newEmpty()

	println("Then try to make a library using the deleted \"borrowers-after.json\" and \"books-before.json\":")
	err = newV(jsonBorrowersFileAfter, jsonBooksFile)
	if err != nil {
		println(err.Error())
	}

	println("And if we read in a file with mal-formed json content... like \"bad-borrowers.json\" and \"books-before.json\":")
	err = newV(jsonBorrowersFileBad, jsonBooksFile)
	if err != nil {
		println(err.Error())
	}

	println("Or how about reading in an empty file... \"empty.json\" (for books):")
	err = newV(jsonBorrowersFileBefore, emptyFile)
	if err != nil {
		println(err.Error())
	}

	println("And... that's all...")
	println("Thanks - bye!\n")
}

func ReadFileIntoJSONString(f string) (string, error) {
	raw, err := ioutil.ReadFile(f)
	if err != nil {
		return "", err
	}
	return string(raw), err
}

func WriteJSONStringToFile(js string, fp string) error {
	f, err := os.Create(fp)
	if err != nil {
		return err
	}
	f.WriteString(js)
	defer f.Close()
	return nil
}

func newEmpty() {
	tvBooks = []book.Book{}
	tvBorrowers = []borrower.Borrower{}
	println(library.StatusToString(tvBooks, tvBorrowers))
}

func newV(brsFile string, bksFile string) error {
	brsPError := error(nil)
	bksPError := error(nil)
	brsJ, brsRError := ReadFileIntoJSONString(brsFile)
	bksJ, bksRError := ReadFileIntoJSONString(bksFile)
	tvBorrowers, brsPError = library.JSONStringToBorrowers(brsJ)
	tvBooks, bksPError = library.JSONStringToBooks(bksJ)
	switch {
	case brsRError != nil:
		return errors.New("borrowers file read error")
	case bksRError != nil:
		return errors.New("books file read error")
	case brsPError != nil:
		return errors.New("borrowers file parse error")
	case bksPError != nil:
		return errors.New("books file parse error")
	default:
		println(library.StatusToString(tvBooks, tvBorrowers))
		return nil
	}
}
