package borrower

import (
	"encoding/json"
	"fmt"
)

// A Borrower has a Name and a MaxBooks
type Borrower struct {
	Name     string `json:"name"`
	MaxBooks int    `json:"maxBooks"`
}

func (br Borrower) Equal(b Borrower) bool {
	if &br == &b {
		return true
	}
	if br.Name != b.Name || br.MaxBooks != b.MaxBooks {
		return false
	}
	return true
}

// JsonStringToBorrower turns a Borrower JSON string into a Borrower
func JsonStringToBorrower(borrowerString string) (Borrower, error) {
	var borrower Borrower
	err := json.Unmarshal([]byte(borrowerString), &borrower)
	return borrower, err
}

// SetName sets a Name for a Borrower
func (br Borrower) SetName(name string) Borrower {
	br.Name = name
	return br
}

//SetMaxBooks sets a MaxBooks for a Borrower
func (br Borrower) SetMaxBooks(maxBooks int) Borrower {
	br.MaxBooks = maxBooks
	return br
}

// String makes a Borrower into a string
func (br Borrower) String() string {
	return fmt.Sprintf("%s [%d books]", br.Name, br.MaxBooks)
}

// BrToJsonString turns a Borrower into a Borrower JSON string
func (br Borrower) BrToJsonString() (string, error) {
	borrowerByte, err := json.Marshal(br)
	return string(borrowerByte), err
}
