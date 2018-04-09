package borrower

import "strconv"

// A Borrower has a Name and a max books
type Borrower struct {
	Name     string `json:"name"`
	MaxBooks int    `json:"max-books"`
}

// MakeBorrower needs a Name and a max books to create
func MakeBorrower(n string, mb int) Borrower {
	br := Borrower{
		Name:     n,
		MaxBooks: mb,
	}
	return br
}

// SetName sets a Name for a Borrower
func (br *Borrower) SetName(n string) {
	br.Name = n
	return
}

// SetMaxBooks sets a max books for a Borrower
func (br *Borrower) SetMaxBooks(mb int) {
	br.MaxBooks = mb
	return
}

// BorrowerToString makes a Borrower into a string
func (br *Borrower) BorrowerToString() string {
	return br.Name + " (" + strconv.Itoa(br.MaxBooks) + " books)"
}
