package borrower

import "strconv"

// A Borrower has a Name and a max books
type Borrower struct {
	Name     string `json:"name"`
	MaxBooks int    `json:"max-books"`
}

// NewBorrower needs a Name and a max books to create
func NewBorrower(name string, maxBooks int) *Borrower {
	return &Borrower{Name: name, MaxBooks: maxBooks}
}

// SetName sets a Name for a Borrower
func (b *Borrower) SetName(name string) {
	b.Name = name
}

// SetMaxBooks sets a max books for a Borrower
func (b *Borrower) SetMaxBooks(maxBooks int) {
	b.MaxBooks = maxBooks
}

// BorrowerToString makes a Borrower into a string
func (b *Borrower) BorrowerToString() string {
	return b.Name + " (" + strconv.Itoa(b.MaxBooks) + " books)"
}