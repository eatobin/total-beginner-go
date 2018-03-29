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

func main() {
	tvBorrowers = append(tvBorrowers, borrower.MakeBorrower("Jim", 3))
	tvBorrowers = append(tvBorrowers, borrower.MakeBorrower("Sue", 3))
	tvBooks = append(tvBooks, book.MakeBook("War And Peace", "Tolstoy"))
	tvBooks = append(tvBooks, book.MakeBook("Great Expectations", "Dickens"))
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
	tvBorrowers = append(tvBorrowers, borrower.MakeBorrower("Eric", 1))
	tvBooks = append(tvBooks, book.MakeBook("The Cat In The Hat", "Dr. Seuss"))
	println("Check Out Dr. Seuss to Eric")
	tvBooks = library.CheckOut("Eric", "The Cat In The Hat", tvBorrowers, tvBooks)
	println(library.StatusToString(tvBooks, tvBorrowers))

	println("Now let's do some BAD stuff...\n")

	println("Add a borrower that already exists (borrower.Borrower{\"Jim\", 3})")
	tvBorrowers = append(tvBorrowers, borrower.MakeBorrower("Jim", 3))
	println("No change to Test Library:")
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
