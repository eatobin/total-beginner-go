package library

import (
	"eatobin.com/totalbeginnergo/book"
	"eatobin.com/totalbeginnergo/borrower"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

type BorrowerPtr = *borrower.Borrower
type Borrowers = []BorrowerPtr
type BookPtr = *book.Book
type Books = []BookPtr

var ZeroBorrowers Borrowers
var ZeroBooks Books

func containsBorrower(brs Borrowers, br BorrowerPtr) bool {
	for _, b := range brs {
		if *b == *br {
			return true
		}
	}
	return false
}

func containsBook(bks Books, bk BookPtr) bool {
	for _, b := range bks {
		if *b == *bk {
			return true
		}
	}
	return false
}

// AddBorrower adds a Borrower pointer to a slice of Borrower pointers
func AddBorrower(brs Borrowers, br BorrowerPtr) Borrowers {
	if containsBorrower(brs, br) {
		return brs
	}
	return append(brs, br)
}

// AddBook adds a Book pointer to a slice of Book pointers
func AddBook(bks Books, bk BookPtr) Books {
	if containsBook(bks, bk) {
		return bks
	}
	return append(bks, bk)
}

// removeBook removes a Book pointer from a slice of Book pointers
func removeBook(bks Books, bk BookPtr) Books {
	nBks := make(Books, 0)
	for _, nBk := range bks {
		if *nBk != *bk {
			nBks = append(nBks, nBk)
		}
	}
	return nBks
}

// findBorrower finds a Borrower pointer given a Name
func findBorrower(n string, brs Borrowers) (BorrowerPtr, error) {
	for _, br := range brs {
		if br.Name == n {
			return br, nil
		}
	}
	return nil, errors.New("did not find the requested borrower")
}

// findBook finds a Book pointer given a Title
func findBook(t string, bks Books) (BookPtr, error) {
	for _, bk := range bks {
		if bk.Title == t {
			return bk, nil
		}
	}
	return nil, errors.New("did not find the requested book")
}

// getBooksForBorrower will find books given a Borrower and a slice of Book pointers
func getBooksForBorrower(br BorrowerPtr, bks Books) Books {
	nBks := make(Books, 0)
	for _, bk := range bks {
		if bk.Borrower != nil {
			if *bk.Borrower == *br {
				nBks = append(nBks, bk)
			}
		}
	}
	return nBks
}

// numberBooksOut returns the # Books checked out to a Borrower
func numberBooksOut(br BorrowerPtr, bks Books) int {
	return len(getBooksForBorrower(br, bks))
}

// notMaxedOut returns True if books out < max books
func notMaxedOut(br BorrowerPtr, bks Books) bool {
	return numberBooksOut(br, bks) < br.MaxBooks
}

func bookNotOut(bk BookPtr) bool {
	return bk.Borrower == nil
}

func bookOut(bk BookPtr) bool {
	return bk.Borrower != nil
}

func CheckOut(n string, t string, brs Borrowers, bks Books) Books {
	mbr, errBr := findBorrower(n, brs)
	mbk, errBk := findBook(t, bks)
	if errBr == nil && errBk == nil && notMaxedOut(mbr, bks) && bookNotOut(mbk) {
		newBook := mbk.SetBorrower(mbr)
		fewerBooks := removeBook(bks, mbk)
		return AddBook(fewerBooks, &newBook)
	}
	return bks
}

func CheckIn(t string, bks Books) Books {
	mbk, errBk := findBook(t, bks)
	if errBk == nil && bookOut(mbk) {
		newBook := book.Book.SetBorrower(mbk, nil)
		fewerBooks := removeBook(bks, mbk)
		return AddBook(fewerBooks, &newBook)
	}
	return bks
}

func JsonStringToBorrowers(borrowersString string) (Borrowers, error) {
	borrowers := ZeroBorrowers
	err := json.Unmarshal([]byte(borrowersString), &borrowers)
	if err != nil {
		return ZeroBorrowers, err
	}
	for _, br := range borrowers {
		if br.Name == "" || br.MaxBooks == 0 {
			err = errors.New("missing Borrower field value - borrowers list is empty")
			return ZeroBorrowers, err
		}
	}
	return borrowers, err
}

func JsonStringToBooks(bookString string) (Books, error) {
	books := ZeroBooks
	err := json.Unmarshal([]byte(bookString), &books)
	if err != nil {
		return ZeroBooks, err
	}
	for _, bk := range books {
		if bk.Title == "" || bk.Author == "" {
			err = errors.New("missing Book field value - book list is empty")
			return ZeroBooks, err
		}
		if bk.Borrower != nil {
			if bk.Borrower.Name == "" || bk.Borrower.MaxBooks == 0 {
				err = errors.New("missing Borrower field value - book list is empty")
				return ZeroBooks, err
			}
		}
	}
	return books, err
}

func BorrowersToJSONSting(brs Borrowers) (string, error) {
	bytes, err := json.Marshal(brs)
	return string(bytes), err
}

func BooksToJSONSting(bks Books) (string, error) {
	bytes, err := json.Marshal(bks)
	return string(bytes), err
}

func libraryToString(bks Books, brs Borrowers) string {
	return "Test Library: " +
		strconv.Itoa(len(bks)) + " books; " +
		strconv.Itoa(len(brs)) + " borrowers."
}

func StatusToString(bks Books, brs Borrowers) string {
	var sb strings.Builder
	sb.WriteString("\n--- Status Report of Test Library ---\n\n")
	sb.WriteString(libraryToString(bks, brs) + "\n\n")
	for _, bk := range bks {
		sb.WriteString(book.Book.String(bk) + "\n")
	}
	sb.WriteString("\n")
	for _, br := range brs {
		sb.WriteString(borrower.Borrower.String(br) + "\n")
	}
	sb.WriteString("\n--- End of Status Report ---\n")
	return sb.String()
}
