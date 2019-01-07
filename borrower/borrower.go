package borrower

import "strconv"

// A Borrower has a Name and a max books
type Borrower struct {
	Name     string `json:"name"`
	MaxBooks int    `json:"max-books"`
}

// NewBorrower needs a Name and a max books to create
func NewBorrower(n string, mb int) Borrower {
	br := Borrower{}
	br.Name = n
	br.MaxBooks = mb
	return br
}

func (br Borrower) GetName() string {
	return br.Name
}

// SetName sets a Name for a Borrower
func (br Borrower) SetName(n string) Borrower {
	br.Name = n
	return br
}

func (br Borrower) GetMaxBooks() int {
	return br.MaxBooks
}

// SetMaxBooks sets a max books for a Borrower
func (br Borrower) SetMaxBooks(mb int) Borrower {
	br.MaxBooks = mb
	return br
}

// BorrowerToString makes a Borrower into a string
func (br Borrower) BorrowerToString() string {
	return br.Name + " (" + strconv.Itoa(br.MaxBooks) + " books)"
}
