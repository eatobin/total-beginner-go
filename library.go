package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func (br *Borrower) String() string {
	return fmt.Sprintf("Name = %s, MaxBooks = %d\n", br.Name, br.MaxBooks)
}

func (bk *Book) String() string {
	return fmt.Sprintf("Title = %s, Author = %s, Borrower = %+v\n", bk.Title, bk.Author, bk.Borrower)
}

func containsBorrower(brs []*Borrower, br *Borrower) bool {
	for _, b := range brs {
		if *b == *br {
			return true
		}
	}
	return false
}

func containsBook(bks []*Book, bk *Book) bool {
	for _, b := range bks {
		if *b == *bk {
			return true
		}
	}
	return false
}

// AddBorrower adds a Borrower to a slice of Borrowers
func AddBorrower(brs []*Borrower, br *Borrower) []*Borrower {
	if containsBorrower(brs, br) {
		return brs
	}
	return append(brs, br)
}

// AddBook adds a book to a slice of Books
func AddBook(bks []*Book, bk *Book) []*Book {
	if containsBook(bks, bk) {
		return bks
	}
	return append(bks, bk)
}

// findBorrower finds a Borrower given a Name
func findBorrower(n string, brs []*Borrower) (error, *Borrower) {
	for _, br := range brs {
		if br.Name == n {
			return nil, br
		}
	}
	return errors.New("did not find the requested borrower"), &Borrower{}
}

// findBook finds a Book given a Title
func findBook(t string, bks []*Book) (int, error, *Book) {
	for i, bk := range bks {
		if bk.Title == t {
			return i, nil, bk
		}
	}
	return 0, errors.New("did not find the requested book"), &Book{}
}

// GetBooksForBorrower will find books given a Borrower and a slice of Books
func getBooksForBorrower(br *Borrower, bks []*Book) []*Book {
	nBks := make([]*Book, 0)
	for _, bk := range bks {
		if bk.Borrower == br {
			nBks = append(nBks, bk)
		}
	}
	return nBks
}

// numberBooksOut returns the # Books checked out to a Borrower
func numberBooksOut(br *Borrower, bks []*Book) int {
	return len(getBooksForBorrower(br, bks))
}

// notMaxedOut returns True if books out < max books
func notMaxedOut(br *Borrower, bks []*Book) bool {
	return numberBooksOut(br, bks) < br.MaxBooks
}

func bookNotOut(bk *Book) bool {
	return bk.Borrower == nil
}

func bookOut(bk *Book) bool {
	return bk.Borrower != nil
}

func CheckOut(n string, t string, brs []*Borrower, bks []*Book) []*Book {
	errBr, mbr := findBorrower(n, brs)
	i, errBk, mbk := findBook(t, bks)
	if errBr == nil && errBk == nil && notMaxedOut(mbr, bks) && bookNotOut(mbk) {
		bks[i].SetBorrower(mbr)
		return bks
	}
	return bks
}

func CheckIn(t string, bks []*Book) []*Book {
	i, errBk, mbk := findBook(t, bks)
	if errBk == nil && bookOut(mbk) {
		bks[i].SetBorrower(nil)
		return bks
	}
	return bks
}

func JsonStringToBorrowers(borrowersString string) ([]*Borrower, error) {
	var borrowers []*Borrower
	err := json.Unmarshal([]byte(borrowersString), &borrowers)
	if err != nil {
		return []*Borrower{}, err
	}
	for _, br := range borrowers {
		if br.Name == "" || br.MaxBooks == 0 {
			err = errors.New("missing Borrower field value - borrowers list is empty")
			return []*Borrower{}, err
		}
	}
	return borrowers, err
}

//FIXME
func JsonStringToBooks(booksString string) ([]*Book, error) {
	var books []*Book
	err := json.Unmarshal([]byte(booksString), &books)
	return books, err
}

func BorrowersToJSONSting(brs []*Borrower) (string, error) {
	bytes, err := json.Marshal(brs)
	return string(bytes), err
}

func BooksToJSONSting(bks []*Book) (string, error) {
	bytes, err := json.Marshal(bks)
	return string(bytes), err
}

func libraryToString(bks []*Book, brs []*Borrower) string {
	return "Test Library: " +
		strconv.Itoa(len(bks)) + " books; " +
		strconv.Itoa(len(brs)) + " borrowers."
}

func StatusToString(bks []*Book, brs []*Borrower) string {
	var sb strings.Builder
	sb.WriteString("\n--- Status Report of Test Library ---\n\n")
	sb.WriteString(libraryToString(bks, brs) + "\n\n")
	for _, bk := range bks {
		sb.WriteString(bk.BkToString() + "\n")
	}
	sb.WriteString("\n")
	for _, br := range brs {
		sb.WriteString(br.BrToString() + "\n")
	}
	sb.WriteString("\n--- End of Status Report ---\n")
	return sb.String()
}
