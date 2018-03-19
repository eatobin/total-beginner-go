package borrower

import "strconv"

type Borrower struct {
	Name     string `json:"name"`
	MaxBooks int    `json:"max-books"`
}

func MakeBorrower(n string, mb int) Borrower {
	br := Borrower{
		Name:     n,
		MaxBooks: mb,
	}
	return br
}

func (br *Borrower) SetName(n string) {
	br.Name = n
	return
}

func (br *Borrower) SetMaxBooks(mb int) {
	br.MaxBooks = mb
	return
}

func (br *Borrower) BorrowerToString() string {
	return br.Name + " (" + strconv.Itoa(br.MaxBooks) + " books)"
}
