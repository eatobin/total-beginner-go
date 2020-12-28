package main

import (
	"encoding/json"
	"strconv"
)

// A Borrower has a Name and a max books
type Borrower struct {
	Name     string `json:"name"`
	MaxBooks int    `json:"maxBooks"`
}

// NewBorrower needs a Name and a max books to create
func NewBorrower(name string, maxBooks int) *Borrower {
	return &Borrower{Name: name, MaxBooks: maxBooks}
}

// SetName sets a Name for a Borrower
func (br *Borrower) SetName(name string) {
	br.Name = name
}

// SetMaxBooks sets a max books for a Borrower
func (br *Borrower) SetMaxBooks(maxBooks int) {
	br.MaxBooks = maxBooks
}

// BrToString makes a Borrower into a string
func (br *Borrower) BrToString() string {
	return br.Name + " (" + strconv.Itoa(br.MaxBooks) + " books)"
}

func JsonStringToBorrower(borrowerString string) (*Borrower, error) {
	var borrower *Borrower
	err := json.Unmarshal([]byte(borrowerString), &borrower)
	return borrower, err
}
