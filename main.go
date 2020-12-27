package main

//import (
//	"fmt"
//	"io/ioutil"
//)
//
//var borrowers []*Borrower
//var books []*Book
//var jsonBorrowersFileBefore = "resources/borrowers-before.json"
//var jsonBooksFile = "resources/books-before.json"
//
////var jsonBorrowersFileAfter = "resources/borrowers-after.json"
////var jsonBorrowersFileBad = "resources/bad-borrowers.json"
////var emptyFile = "resources/resources/empty.json"
//
//func main() {
//	borrowers = AddBorrower(borrowers, NewBorrower("Jim", 3))
//	borrowers = AddBorrower(borrowers, NewBorrower("Sue", 3))
//	books = AddBook(books, NewBook("War And Peace", "Tolstoy"))
//	books = AddBook(books, NewBook("Great Expectations", "Dickens"))
//	fmt.Println("\nJust created new library")
//	fmt.Println(StatusToString(books, borrowers))
//
//	fmt.Println("Check out War And Peace to Sue")
//	books = CheckOut("Sue", "War And Peace", borrowers, books)
//	fmt.Println(StatusToString(books, borrowers))
//
//	fmt.Println("Now check in War And Peace from Sue...")
//	books = CheckIn("War And Peace", books)
//	fmt.Println("...and check out Great Expectations to Jim")
//	books = CheckOut("Jim", "Great Expectations", borrowers, books)
//	fmt.Println(StatusToString(books, borrowers))
//
//	fmt.Println("Add Eric and The Cat In The Hat")
//	borrowers = AddBorrower(borrowers, NewBorrower("Eric", 1))
//	books = AddBook(books, NewBook("The Cat In The Hat", "Dr. Seuss"))
//	fmt.Println("Check Out Dr. Seuss to Eric")
//	books = CheckOut("Eric", "The Cat In The Hat", borrowers, books)
//	fmt.Println(StatusToString(books, borrowers))
//
//	fmt.Println("Now let's do some BAD stuff...")
//
//	fmt.Println("Add a borrower that already exists (Borrower('Jim', 3))")
//	borrowers = AddBorrower(borrowers, NewBorrower("Jim", 3))
//	fmt.Println("No change to Test Library:")
//	fmt.Println(StatusToString(books, borrowers))
//
//	fmt.Println("Add a book that already exists (Book('War And Peace', 'Tolstoy'))")
//	books = AddBook(books, NewBook("War And Peace", "Tolstoy"))
//	fmt.Println("No change to Test Library:")
//	fmt.Println(StatusToString(books, borrowers))
//
//	fmt.Println("Check out a valid book to an invalid person (checkOut('JoJo', 'War And Peace'))")
//	books = CheckOut("JoJo", "War And Peace", borrowers, books)
//	fmt.Println("No change to Test Library:")
//	fmt.Println(StatusToString(books, borrowers))
//
//	fmt.Println("Check out an invalid book to an valid person (checkOut('Sue', 'Not A Book'))")
//	books = CheckOut("Sue", "Not A Book", borrowers, books)
//	fmt.Println("No change to Test Library:")
//	fmt.Println(StatusToString(books, borrowers))
//
//	fmt.Println("Last - check in a book not checked out (checkIn('War And Peace'))")
//	books = CheckIn("War And Peace", books)
//	fmt.Println("No change to Test Library:")
//	fmt.Println(StatusToString(books, borrowers))
//
//	fmt.Println("Okay... let's finish with some persistence. First clear the whole library:")
//	newEmptyV()
//
//	fmt.Println("Lets read in a new library from \"borrowers-before.json\" and \"books-before.json\":")
//	newVError := newV(jsonBorrowersFileBefore, jsonBooksFile)
//	if newVError != nil {
//		panic(newVError)
//	}
//	fmt.Println("Add... a new borrower:")
//	borrowers = AddBorrower(borrowers, NewBorrower("BorrowerNew", 300))
//	fmt.Println(StatusToString(books, borrowers))
//}
//
//func newEmptyV() {
//	borrowers = []*Borrower{}
//	books = []*Book{}
//	fmt.Println(StatusToString(books, borrowers))
//}
//
//func readFileIntoJsonString(fp string) (string, error) {
//	dat, err := ioutil.ReadFile(fp)
//	return string(dat), err
//}
//
//func newV(brsFp string, bksFp string) error {
//	jsonBrsStr, brFileErr := readFileIntoJsonString(brsFp)
//	if brFileErr != nil {
//		return brFileErr
//	}
//	jsonBksStr, bkFileErr := readFileIntoJsonString(bksFp)
//	if bkFileErr != nil {
//		return bkFileErr
//	}
//	brs, brParseErr := jsonStringToBorrowers(jsonBrsStr)
//	if brParseErr != nil {
//		return brParseErr
//	}
//	bks, bkParseErr := jsonStringToBooks(jsonBksStr)
//	if bkParseErr != nil {
//		return bkParseErr
//	}
//	borrowers = brs
//	books = bks
//	fmt.Println(StatusToString(books, borrowers))
//	return nil
//}
