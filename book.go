package main

// A Book has a Title and an Author
type Book struct {
	Title    string   `json:"title"`
	Author   string   `json:"author"`
	Borrower Borrower `json:"borrower"`
}

var ZeroBorrower = Borrower{}

func NewBook(title string, author string) Book {
	return Book{Title: title, Author: author, Borrower: ZeroBorrower}
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
func SetBorrower(bk Book, borrower Borrower) Book {
	bk.Borrower = borrower
	return bk
}

func availableString(bk Book) string {
	if bk.Borrower == ZeroBorrower {
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
