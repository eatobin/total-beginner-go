package library

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"eatobin.com/totalbeginnergo/book"
	"eatobin.com/totalbeginnergo/borrower"
)

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
func findBorrower(n string, brs []borrower.Borrower) (error, borrower.Borrower) {
	for _, br := range brs {
		if br.Name == n {
			return nil, br
		}
	}
	return errors.New("did not find the requested borrower"), book.ZeroBorrower
}

// findBook finds a Book given a Title
func findBook(t string, bks []book.Book) (error, book.Book) {
	for _, bk := range bks {
		if bk.Title == t {
			return nil, bk
		}
	}
	return errors.New("did not find the requested book"), book.Book{}
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
	return bk.Borrower == book.ZeroBorrower
}

func bookOut(bk book.Book) bool {
	return bk.Borrower != book.ZeroBorrower
}

func CheckOut(n string, t string, brs []borrower.Borrower, bks []book.Book) []book.Book {
	errBr, mbr := findBorrower(n, brs)
	errBk, mbk := findBook(t, bks)
	if errBr == nil && errBk == nil && notMaxedOut(mbr, bks) && bookNotOut(mbk) {
		newBook := book.SetBorrower(mbk, mbr)
		fewerBooks := removeBook(mbk, bks)
		return AddBook(fewerBooks, newBook)
	}
	return bks
}

func CheckIn(t string, bks []book.Book) []book.Book {
	errBk, mbk := findBook(t, bks)
	if errBk == nil && bookOut(mbk) {
		newBook := book.SetBorrower(mbk, book.ZeroBorrower)
		fewerBooks := removeBook(mbk, bks)
		return AddBook(fewerBooks, newBook)
	}
	return bks
}

func jsonStringToBorrowers(js string) (error, []borrower.Borrower) {
	var res []borrower.Borrower
	err := json.Unmarshal([]byte(js), &res)
	if err != nil {
		return err, []borrower.Borrower{}
	}
	for _, br := range res {
		if br.Name == "" || br.MaxBooks == 0 {
			err = errors.New("missing Borrower field value - borrowers list is empty")
			return err, []borrower.Borrower{}
		}
	}
	return err, res
}

func jsonStringToBooks(js string) (error, []book.Book) {
	var res []book.Book
	err := json.Unmarshal([]byte(js), &res)
	if err != nil {
		return err, []book.Book{}
	}
	for _, bk := range res {
		if bk.Title == "" || bk.Author == "" || bk.Borrower.Name == "" || bk.Borrower.MaxBooks == 0 {
			err = errors.New("missing Book field value - book list is empty")
			return err, []book.Book{}
		}
	}
	return nil, res
}

func BorrowersToJSONSting(brs []borrower.Borrower) string {
	bytes, _ := json.MarshalIndent(brs, "", "  ")
	return string(bytes)
}

func BooksToJSONSting(bks []book.Book) string {
	bytes, _ := json.MarshalIndent(bks, "", "  ")
	return string(bytes)
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
		sb.WriteString(book.BkToString(bk) + "\n")
	}
	sb.WriteString("\n")
	for _, br := range brs {
		sb.WriteString(borrower.BrToString(br) + "\n")
	}
	sb.WriteString("\n--- End of Status Report ---\n")
	return sb.String()
}
