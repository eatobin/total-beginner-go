package book

import "eatobin.com/totalbeginnergo/borrower"

// A Book has a Title and an Author
type Book struct {
	Title    string            `json:"title"`
	Author   string            `json:"author"`
	Borrower borrower.Borrower `json:"borrower"`
}

var zeroBorrower = borrower.Borrower{}

func NewBook(title string, author string) Book {
	return Book{Title: title, Author: author, Borrower: zeroBorrower}
}

// setTitle sets a Title for a Book
func setTitle(bk Book, title string) Book {
	bk.Title = title
	return bk
}

// setAuthor sets a Author for a Book
func setAuthor(bk Book, author string) Book {
	bk.Author = author
	return bk
}

// SetBorrower takes a BorrowerPtr and sets it for a Book
func SetBorrower(bk Book, borrower borrower.Borrower) Book {
	bk.Borrower = borrower
	return bk
}

func availableString(bk Book) string {
	if bk.Borrower == zeroBorrower {
		return "Available"
	}
	return "Checked out to " +
		bk.Borrower.Name
}

// BkToString makes a description of a Book
func BkToString(bk Book) string {
	return bk.Title +
		" by " + bk.Author +
		"; " + availableString(bk)
}
