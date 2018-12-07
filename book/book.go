package book

import "github.com/eatobin/totalbeginnergo/borrower"

// A Book has a Title and an Author
type Book struct {
	Title    string            `json:"title"`
	Author   string            `json:"author"`
	Borrower borrower.Borrower `json:"borrower"`
}

// NewBook needs a Title and an Author to make a Book
func NewBook(t string, a string) *Book {
	bk := Book{}
	bk.Title = t
	bk.Author = a
	return &bk
}

// SetBorrower takes a Borrower and sets it for a Book
func (bk *Book) SetBorrower(br *borrower.Borrower) {
	bk.Borrower = *br
}

func (bk *Book) availableString() string {
	if bk.Borrower == (borrower.Borrower{Name: "", MaxBooks: 0}) {
		return "Available"
	}
	return "Checked out to " +
		bk.Borrower.Name
}

// BookToString makes a description of a Book
func (bk *Book) BookToString() string {
	return bk.Title +
		" by " + bk.Author +
		"; " + bk.availableString()
}
