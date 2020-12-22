package main

import "strconv"

// A Borrower has a Name and a max books
type Borrower struct {
	Name     string `json:"name"`
	MaxBooks int    `json:"max-books"`
}

// NewBorrower needs a Name and a max books to create
func NewBorrower(name string, maxBooks int) Borrower {
	return Borrower{Name: name, MaxBooks: maxBooks}
}

// SetName sets a Name for a Borrower
func SetName(br Borrower, name string) Borrower {
	br.Name = name
	return br
}

// SetMaxBooks sets a max books for a Borrower
func SetMaxBooks(br Borrower, maxBooks int) Borrower {
	br.MaxBooks = maxBooks
	return br
}

// BrToString makes a Borrower into a string
func BrToString(br Borrower) string {
	return br.Name + " (" + strconv.Itoa(br.MaxBooks) + " books)"
}
