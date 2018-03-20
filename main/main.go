package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/eatobin/totalbeginnergo/library"
)

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
	brsJ := ReadFileIntoJsonString("borrowers-before.json")
	fmt.Printf("%q\n", brsJ)
	bksJ := ReadFileIntoJsonString("books-before.json")
	fmt.Printf("%q\n", bksJ)
	brs := library.JSONStringToBorrowers(brsJ)
	bks := library.JSONStringToBooks(bksJ)
	nBks := library.CheckOut("Borrower200", "Book200", brs, bks)
	nBksJ := library.BooksToJSONSting(nBks)
	WriteJSONStringToFile(nBksJ, "books-after.json")
}
