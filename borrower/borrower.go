package borrower

import (
	"encoding/json"
	"fmt"
)

// A Borrower has a Name and a max books
type Borrower struct {
	Name     string `json:"name"`
	MaxBooks int    `json:"maxBooks"`
}

var ZeroBorrower = Borrower{}

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

// String makes a Borrower into a string
func String(br Borrower) string {
	return fmt.Sprintf("%s (%d books)", br.Name, br.MaxBooks)
}

func JsonStringToBorrower(borrowerString string) (Borrower, error) {
	var borrower Borrower
	err := json.Unmarshal([]byte(borrowerString), &borrower)
	return borrower, err
}

func ToJsonString(borrower Borrower) (string, error) {
	borrowerByte, err := json.Marshal(borrower)
	return string(borrowerByte), err
}
