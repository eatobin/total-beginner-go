package library

import (
	"eatobin.com/totalbeginnergo/book"
	"eatobin.com/totalbeginnergo/borrower"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

var ZeroBorrowers []borrower.Borrower
var ZeroBooks []book.Book

func containsBorrower(brs []borrower.Borrower, br borrower.Borrower) bool {
	for _, b := range brs {
		if b == br {
			return true
		}
	}
	return false
}

func containsBook(bks []book.Book, bk book.Book) bool {
	for _, b := range bks {
		if b == bk {
			return true
		}
	}
	return false
}

// AddBorrower adds a Borrower to a slice of Borrowers
func AddBorrower(brs []borrower.Borrower, br borrower.Borrower) []borrower.Borrower {
	if containsBorrower(brs, br) {
		return brs
	}
	return append(brs, br)
}

// AddBook adds a book to a slice of Books
func AddBook(bks []book.Book, bk book.Book) []book.Book {
	if containsBook(bks, bk) {
		return bks
	}
	return append(bks, bk)
}

// removeBook removes a book from a slice of Books
func removeBook(bk book.Book, bks []book.Book) []book.Book {
	nBks := make([]book.Book, 0)
	for _, nBk := range bks {
		if nBk != bk {
			nBks = append(nBks, nBk)
		}
	}
	return nBks
}

// findBorrower finds a Borrower given a Name
func findBorrower(n string, brs []borrower.Borrower) (borrower.Borrower, error) {
	for _, br := range brs {
		if br.Name == n {
			return br, nil
		}
	}
	return borrower.ZeroBorrower, errors.New("did not find the requested borrower")
}

// findBook finds a Book given a Title
func findBook(t string, bks []book.Book) (book.Book, error) {
	for _, bk := range bks {
		if bk.Title == t {
			return bk, nil
		}
	}
	return book.ZeroBook, errors.New("did not find the requested book")
}

// getBooksForBorrower will find books given a Borrower and a slice of Books
func getBooksForBorrower(br borrower.Borrower, bks []book.Book) []book.Book {
	nBks := make([]book.Book, 0)
	for _, bk := range bks {
		if bk.Borrower == br {
			nBks = append(nBks, bk)
		}
	}
	return nBks
}

// numberBooksOut returns the # Books checked out to a Borrower
func numberBooksOut(br borrower.Borrower, bks []book.Book) int {
	return len(getBooksForBorrower(br, bks))
}

// notMaxedOut returns True if books out < max books
func notMaxedOut(br borrower.Borrower, bks []book.Book) bool {
	return numberBooksOut(br, bks) < br.MaxBooks
}

func bookNotOut(bk book.Book) bool {
	return bk.Borrower == borrower.ZeroBorrower
}

func bookOut(bk book.Book) bool {
	return bk.Borrower != borrower.ZeroBorrower
}

func CheckOut(n string, t string, brs []borrower.Borrower, bks []book.Book) []book.Book {
	mbr, errBr := findBorrower(n, brs)
	mbk, errBk := findBook(t, bks)
	if errBr == nil && errBk == nil && notMaxedOut(mbr, bks) && bookNotOut(mbk) {
		newBook := book.SetBorrower(mbk, mbr)
		fewerBooks := removeBook(mbk, bks)
		return AddBook(fewerBooks, newBook)
	}
	return bks
}

func CheckIn(t string, bks []book.Book) []book.Book {
	mbk, errBk := findBook(t, bks)
	if errBk == nil && bookOut(mbk) {
		newBook := book.SetBorrower(mbk, borrower.ZeroBorrower)
		fewerBooks := removeBook(mbk, bks)
		return AddBook(fewerBooks, newBook)
	}
	return bks
}

func JsonStringToBorrowers(borrowersString string) ([]borrower.Borrower, error) {
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

func JsonStringToBooks(bookString string) ([]book.Book, error) {
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
		if bk.Borrower != borrower.ZeroBorrower {
			if bk.Borrower.Name == "" || bk.Borrower.MaxBooks == 0 {
				err = errors.New("missing Borrower field value - book list is empty")
				return ZeroBooks, err
			}
		}
	}
	return books, err
}

func BorrowersToJSONSting(brs []borrower.Borrower) (string, error) {
	bytes, err := json.Marshal(brs)
	return string(bytes), err
}

func BooksToJSONSting(bks []book.Book) (string, error) {
	bytes, err := json.Marshal(bks)
	return string(bytes), err
}

func libraryToString(bks []book.Book, brs []borrower.Borrower) string {
	return "Test Library: " +
		strconv.Itoa(len(bks)) + " books; " +
		strconv.Itoa(len(brs)) + " borrowers."
}

func StatusToString(bks []book.Book, brs []borrower.Borrower) string {
	var sb strings.Builder
	sb.WriteString("\n--- Status Report of Test Library ---\n\n")
	sb.WriteString(libraryToString(bks, brs) + "\n\n")
	for _, bk := range bks {
		sb.WriteString(book.String(bk) + "\n")
	}
	sb.WriteString("\n")
	for _, br := range brs {
		sb.WriteString(borrower.String(br) + "\n")
	}
	sb.WriteString("\n--- End of Status Report ---\n")
	return sb.String()
}
