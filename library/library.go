package library

import (
	"encoding/json"
	"errors"

	"strconv"
	"strings"

	"github.com/eatobin/totalbeginnergo/book"
	"github.com/eatobin/totalbeginnergo/borrower"
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
	if !containsBorrower(brs, br) {
		return append(brs, br)
	}
	return brs
}

// AddBook adds a book to a slice of Books
func AddBook(bks []book.Book, bk book.Book) []book.Book {
	if !containsBook(bks, bk) {
		return append(bks, bk)
	}
	return bks
}

// FindBorrower finds a Borrower given a Name
func FindBorrower(n string, brs []borrower.Borrower) (borrower.Borrower, error) {
	for _, br := range brs {
		if br.Name == n {
			return br, nil
		}
	}
	return borrower.Borrower{}, errors.New("did not find the requested borrower")
}

// FindBook finds a Book given a Title
func FindBook(t string, bks []book.Book) (int, book.Book, error) {
	for i, bk := range bks {
		if bk.Title == t {
			return i, bk, nil
		}
	}
	return 0, book.Book{}, errors.New("did not find the requested book")
}

// GetBooksForBorrower will find books given a Borrower and a slice of Books
func GetBooksForBorrower(br borrower.Borrower, bks []book.Book) []book.Book {
	nBks := make([]book.Book, 0)
	for _, bk := range bks {
		if bk.Borrower == br {
			nBks = append(nBks, bk)
		}
	}
	return nBks
}

// NumberBooksOut returns the # Books checked out to a Borrower
func NumberBooksOut(br borrower.Borrower, bks []book.Book) int {
	return len(GetBooksForBorrower(br, bks))
}

// NotMaxedOut returns True if books out < max books
func NotMaxedOut(br borrower.Borrower, bks []book.Book) bool {
	return NumberBooksOut(br, bks) < br.MaxBooks
}

func BookNotOut(bk book.Book) bool {
	return bk.Borrower == borrower.Borrower{}
}

func BookOut(bk book.Book) bool {
	return bk.Borrower != borrower.Borrower{}
}

func CheckOut(n string, t string, brs []borrower.Borrower, bks []book.Book) []book.Book {
	mbr, errBr := FindBorrower(n, brs)
	i, mbk, errBk := FindBook(t, bks)
	if errBr == nil && errBk == nil && NotMaxedOut(mbr, bks) && BookNotOut(mbk) {
		bks[i].SetBorrower(mbr)
		return bks
	}
	return bks
}

func CheckIn(t string, bks []book.Book) []book.Book {
	i, mbk, errBk := FindBook(t, bks)
	if errBk == nil && BookOut(mbk) {
		bks[i].SetBorrower(borrower.Borrower{})
		return bks
	}
	return bks
}

func JSONStringToBorrowers(js string) ([]borrower.Borrower, error) {
	var res []borrower.Borrower
	err := json.Unmarshal([]byte(js), &res)
	if err != nil {
		return []borrower.Borrower{}, err
	}
	for _, br := range res {
		if br.Name == "" || br.MaxBooks == 0 {
			err = errors.New("missing Borrower field value - borrowers list is empty")
			return []borrower.Borrower{}, err
		}
	}
	return res, nil
}

func JSONStringToBooks(js string) ([]book.Book, error) {
	var res []book.Book
	err := json.Unmarshal([]byte(js), &res)
	if err != nil {
		return []book.Book{}, err
	}
	for _, bk := range res {
		if bk.Title == "" || bk.Author == "" || bk.Borrower.Name == "" || bk.Borrower.MaxBooks == 0 {
			err = errors.New("missing Book field value - book list is empty")
			return []book.Book{}, err
		}
	}
	return res, nil
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
		sb.WriteString(bk.BookToString() + "\n")
	}
	sb.WriteString("\n")
	for _, br := range brs {
		sb.WriteString(br.BorrowerToString() + "\n")
	}
	sb.WriteString("\n--- End of Status Report ---\n")
	return sb.String()
}
