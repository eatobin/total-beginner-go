package book

import "github.com/eatobin/totalbeginnergo/borrower"

// A Book has a Title and an Author
type Book struct {
	Title    string            `json:"title"`
	Author   string            `json:"author"`
	Borrower borrower.Borrower `json:"borrower"`
}

func NewBook(title string, author string) *Book {
	return &Book{Title: title, Author: author}
}

// SetTitle sets a Title for a Book
func (b *Book) SetTitle(title string) {
	b.Title = title
}

// SetAuthor sets a Author for a Book
func (b *Book) SetAuthor(author string) {
	b.Author = author
}

// SetBorrower takes a Borrower and sets it for a Book
func (b *Book) SetBorrower(borrower borrower.Borrower) {
	b.Borrower = borrower
}

func (b *Book) availableString() string {
	if b.Borrower == (borrower.Borrower{Name: "", MaxBooks: 0}) {
		return "Available"
	}
	return "Checked out to " +
		b.Borrower.Name
}

// BookToString makes a description of a Book
func (b *Book) BookToString() string {
	return b.Title +
		" by " + b.Author +
		"; " + b.availableString()
}
