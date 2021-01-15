package book

import (
	"eatobin.com/totalbeginnergo/borrower"
	"encoding/json"
	"fmt"
)

// A Book has a Title and an Author
type Book struct {
	Title    string             `json:"title"`
	Author   string             `json:"author"`
	Borrower *borrower.Borrower `json:"borrower"`
}

func NewBook(title string, author string) *Book {
	return &Book{Title: title, Author: author, Borrower: nil}
}

// SetTitle sets a Title for a Book
func (bk *Book) SetTitle(title string) {
	bk.Title = title
}

// SetAuthor sets a Author for a Book
func (bk *Book) SetAuthor(author string) {
	bk.Author = author
}

// SetBorrower takes a BorrowerPtr and sets it for a Book
func (bk *Book) SetBorrower(borrower *borrower.Borrower) {
	bk.Borrower = borrower
}

func (bk *Book) availableString() string {
	if bk.Borrower == nil {
		return "Available"
	}
	return "Checked out to " +
		bk.Borrower.Name
}

// String makes a description of a Book
func (bk *Book) String() string {
	return fmt.Sprintf("%s by %s; %s", bk.Title, bk.Author, bk.availableString())
}

func JsonStringToBook(bookString string) (*Book, error) {
	var book *Book
	err := json.Unmarshal([]byte(bookString), &book)
	return book, err
}
