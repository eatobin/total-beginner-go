package book

import (
	"eatobin.com/totalbeginnergo/borrower"
	"encoding/json"
	"fmt"
)

// A Book has a Title and an Author - and maybe a Borrower
type Book struct {
	Title    string             `json:"title"`
	Author   string             `json:"author"`
	Borrower *borrower.Borrower `json:"borrower,omitempty"`
}

// JsonStringToBook turns a Book JSON string into a Book
func JsonStringToBook(bookString string) (Book, error) {
	var book Book
	err := json.Unmarshal([]byte(bookString), &book)
	return book, err
}

// SetTitle sets a Title for a Book
func (bk Book) SetTitle(title string) Book {
	bk.Title = title
	return bk
}

// SetAuthor sets a Author for a Book
func (bk Book) SetAuthor(author string) Book {
	bk.Author = author
	return bk
}

// SetBorrower takes a Borrower pointer and sets it for a Book
func (bk Book) SetBorrower(borrower *borrower.Borrower) Book {
	bk.Borrower = borrower
	return bk
}

func (bk Book) availableString() string {
	if bk.Borrower == nil {
		return "Available"
	}
	return fmt.Sprintf("Checked out to %s", bk.Borrower.Name)
}

// String makes a description of a Book
func (bk Book) String() string {
	return fmt.Sprintf("%s by %s; %s", bk.Title, bk.Author, Book.availableString(bk))
}

// BkToJsonString turns a Book into a Book JSON string
func BkToJsonString(book Book) (string, error) {
	bookByte, err := json.Marshal(book)
	return string(bookByte), err
}
