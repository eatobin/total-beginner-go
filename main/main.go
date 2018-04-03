package main

import (
	"fmt"
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
	tvBorrowers = library.AddBorrower(tvBorrowers, borrower.MakeBorrower("Jim", 3))
	tvBorrowers = library.AddBorrower(tvBorrowers, borrower.MakeBorrower("Sue", 3))
	tvBooks = library.AddBook(tvBooks, book.MakeBook("War And Peace", "Tolstoy"))
	tvBooks = library.AddBook(tvBooks, book.MakeBook("Great Expectations", "Dickens"))
	println("\nJust created new library")
	println(library.StatusToString(tvBooks, tvBorrowers))

	println("Check out War And Peace to Sue")
	tvBooks = library.CheckOut("Sue", "War And Peace", tvBorrowers, tvBooks)
	println(library.StatusToString(tvBooks, tvBorrowers))

	println("Now check in War And Peace from Sue...")
	tvBooks = library.CheckIn("War And Peace", tvBooks)
	println("...and check out Great Expectations to Jim")
	tvBooks = library.CheckOut("Jim", "Great Expectations", tvBorrowers, tvBooks)
	println(library.StatusToString(tvBooks, tvBorrowers))

	println("Add Eric and The Cat In The Hat")
	tvBorrowers = library.AddBorrower(tvBorrowers, borrower.MakeBorrower("Eric", 1))
	tvBooks = library.AddBook(tvBooks, book.MakeBook("The Cat In The Hat", "Dr. Seuss"))
	println("Check Out Dr. Seuss to Eric")
	tvBooks = library.CheckOut("Eric", "The Cat In The Hat", tvBorrowers, tvBooks)
	println(library.StatusToString(tvBooks, tvBorrowers))

	println("Now let's do some BAD stuff...\n")

	println("Add a borrower that already exists (borrower.Borrower{\"Jim\", 3})")
	tvBorrowers = library.AddBorrower(tvBorrowers, borrower.MakeBorrower("Jim", 3))
	println("No change to Test Library:")
	println(library.StatusToString(tvBooks, tvBorrowers))

	println("Add a book that already exists (book.Book{\"War And Peace\", \"Tolstoy\"})")
	tvBooks = library.AddBook(tvBooks, book.MakeBook("War And Peace", "Tolstoy"))
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
	newV(jsonBorrowersFileBefore, jsonBooksFile)
	println("Add... a new borrower:")
	tvBorrowers = library.AddBorrower(tvBorrowers, borrower.MakeBorrower("BorrowerNew", 300))
	println(library.StatusToString(tvBooks, tvBorrowers))

	println("Save the revised borrowers to \"borrowers-after.json\"")
	nBrsJ := library.BorrowersToJSONSting(tvBorrowers)
	WriteJSONStringToFile(nBrsJ, jsonBorrowersFileAfter)

	println("Clear the whole library again:")
	newEmpty()

	println("Then read in the revised library from \"borrowers-after.json\" and \"books-before.json\":")
	newV(jsonBorrowersFileAfter, jsonBooksFile)
}

func ReadFileIntoJsonString(f string) string {
	raw, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println("File read error. Library is empty.")
		return ""
	}
	return string(raw)
}

func WriteJSONStringToFile(js string, fp string) {
	f, err := os.Create(fp)
	if err != nil {
		fmt.Println("File write error. Library was not saved.")
	}
	f.WriteString(js)
	defer f.Close()
}

func newEmpty() {
	tvBooks = []book.Book{}
	tvBorrowers = []borrower.Borrower{}
	println(library.StatusToString(tvBooks, tvBorrowers))
}

func newV(brsFile string, bksFile string) {
	brsJ := ReadFileIntoJsonString(brsFile)
	bksJ := ReadFileIntoJsonString(bksFile)
	tvBorrowers = library.JSONStringToBorrowers(brsJ)
	tvBooks = library.JSONStringToBooks(bksJ)
	println(library.StatusToString(tvBooks, tvBorrowers))
}

//func main() {
//	brsJ := ReadFileIntoJsonString("borrowers-before.json")
//	fmt.Printf("%q\n", brsJ)
//	bksJ := ReadFileIntoJsonString("books-before.json")
//	fmt.Printf("%q\n", bksJ)
//	brs := library.JSONStringToBorrowers(brsJ)
//	bks := library.JSONStringToBooks(bksJ)
//	nBks := library.CheckOut("Borrower200", "Book200", brs, bks)
//	nBksJ := library.BooksToJSONSting(nBks)
//	WriteJSONStringToFile(nBksJ, "books-after.json")
//}
