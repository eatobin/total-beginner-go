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

func AddBorrower(brs []borrower.Borrower, br borrower.Borrower) []borrower.Borrower {
	if !containsBorrower(brs, br) {
		return append(brs, br)
	}
	return brs
}

func AddBook(bks []book.Book, bk book.Book) []book.Book {
	if !containsBook(bks, bk) {
		return append(bks, bk)
	}
	return bks
}

func FindBorrower(n string, brs []borrower.Borrower) borrower.Borrower {
	for _, br := range brs {
		if br.Name == n {
			return br
		}
	}
	return borrower.Borrower{}
}

func FindBook(t string, bks []book.Book) (int, book.Book) {
	for i, bk := range bks {
		if bk.Title == t {
			return i, bk
		}
	}
	return -1, book.Book{}
}

func GetBooksForBorrower(br borrower.Borrower, bks []book.Book) []book.Book {
	nBks := make([]book.Book, 0)
	for _, bk := range bks {
		if bk.Borrower == br {
			nBks = append(nBks, bk)
		}
	}
	return nBks
}

func NumberBooksOut(br borrower.Borrower, bks []book.Book) int {
	return len(GetBooksForBorrower(br, bks))
}

func NotMaxedOut(br borrower.Borrower, bks []book.Book) bool {
	return NumberBooksOut(br, bks) < br.MaxBooks
}

func BookNotOut(bk book.Book) bool {
	return bk.Borrower == borrower.Borrower{Name: "NoName", MaxBooks: -1}
}

func BookOut(bk book.Book) bool {
	return bk.Borrower != borrower.Borrower{Name: "NoName", MaxBooks: -1}
}

func CheckOut(n string, t string, brs []borrower.Borrower, bks []book.Book) []book.Book {
	i, mbk := FindBook(t, bks)
	mbr := FindBorrower(n, brs)
	if mbk != (book.Book{}) && mbr != (borrower.Borrower{}) && NotMaxedOut(mbr, bks) && BookNotOut(mbk) {
		bks[i].SetBorrower(mbr)
		return bks
	}
	return bks
}

func CheckIn(t string, bks []book.Book) []book.Book {
	i, mbk := FindBook(t, bks)
	if mbk != (book.Book{}) && BookOut(mbk) {
		bks[i].SetBorrower(borrower.Borrower{Name: "NoName", MaxBooks: -1})
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
