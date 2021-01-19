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
	Borrower *borrower.Borrower `json:"borrower,omitempty"`
}

func NewBook(title string, author string) Book {
	return Book{Title: title, Author: author, Borrower: nil}
}

// SetTitle sets a Title for a Book
func SetTitle(bk Book, title string) Book {
	bk.Title = title
	return bk
}

// SetAuthor sets a Author for a Book
func SetAuthor(bk Book, author string) Book {
	bk.Author = author
	return bk
}

// SetBorrower takes a BorrowerPtr and sets it for a Book
func SetBorrower(bk Book, borrower *borrower.Borrower) Book {
	bk.Borrower = borrower
	return bk
}

func availableString(bk Book) string {
	if bk.Borrower == nil {
		return "Available"
	}
	return fmt.Sprintf("Checked out to %s", bk.Borrower.Name)
}

// String makes a description of a Book
func String(bk Book) string {
	return fmt.Sprintf("%s by %s; %s", bk.Title, bk.Author, availableString(bk))
}

// JsonStringToBook turns a Book JSON string into a Book
func JsonStringToBook(bookString string) (Book, error) {
	var book Book
	err := json.Unmarshal([]byte(bookString), &book)
	return book, err
}

// TODO - test this
// BkToJsonString turns a Book into a Book JSON string
func BkToJsonString(book Book) (string, error) {
	bookByte, err := json.Marshal(book)
	return string(bookByte), err
}
