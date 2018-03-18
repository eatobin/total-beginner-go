package totalbeginnergo

import (
	"encoding/json"

	"strconv"
	"strings"
)

func containsBorrower(brs []Borrower, br Borrower) bool {
	for _, b := range brs {
		if b == br {
			return true
		}
	}
	return false
}

func AddBorrower(brs []Borrower, br Borrower) []Borrower {
	if !containsBorrower(brs, br) {
		return append(brs, br)
	}
	return brs
}

func FindBorrower(n string, brs []Borrower) Borrower {
	for _, br := range brs {
		if br.Name == n {
			return br
		}
	}
	return Borrower{}
}

func FindBook(t string, bks []Book) (int, Book) {
	for i, bk := range bks {
		if bk.Title == t {
			return i, bk
		}
	}
	return -1, Book{}
}

func GetBooksForBorrower(br Borrower, bks []Book) []Book {
	nBks := make([]Book, 0)
	for _, bk := range bks {
		if bk.Borrower == br {
			nBks = append(nBks, bk)
		}
	}
	return nBks
}

func NumberBooksOut(br Borrower, bks []Book) int {
	return len(GetBooksForBorrower(br, bks))
}

func NotMaxedOut(br Borrower, bks []Book) bool {
	return NumberBooksOut(br, bks) < br.MaxBooks
}

func BookNotOut(bk Book) bool {
	return bk.Borrower == Borrower{Name: "NoName", MaxBooks: -1}
}

func BookOut(bk Book) bool {
	return bk.Borrower != Borrower{Name: "NoName", MaxBooks: -1}
}

func CheckOut(n string, t string, brs []Borrower, bks []Book) []Book {
	i, mbk := FindBook(t, bks)
	mbr := FindBorrower(n, brs)
	if mbk != (Book{}) && mbr != (Borrower{}) && NotMaxedOut(mbr, bks) && BookNotOut(mbk) {
		bks[i].SetBorrower(mbr)
		return bks
	}
	return bks
}

func CheckIn(t string, bks []Book) []Book {
	i, mbk := FindBook(t, bks)
	if mbk != (Book{}) && BookOut(mbk) {
		bks[i].SetBorrower(Borrower{Name: "NoName", MaxBooks: -1})
		return bks
	}
	return bks
}

func JSONStringToBorrowers(js string) ([]Borrower, string) {
	var res []Borrower
	err := json.Unmarshal([]byte(js), &res)
	if err != nil {
		return []Borrower{}, "JSON parse error."
	}
	for _, br := range res {
		if br.Name == "" || br.MaxBooks == 0 {
			return []Borrower{}, "Missing Borrower field value."
		}
	}
	return res, ""
}

func JSONStringToBooks(js string) ([]Book, string) {
	var res []Book
	err := json.Unmarshal([]byte(js), &res)
	if err != nil {
		return []Book{}, "JSON parse error."
	}
	for _, bk := range res {
		if bk.Title == "" || bk.Author == "" || bk.Borrower.Name == "" || bk.Borrower.MaxBooks == 0 {
			return []Book{}, "Missing Book field value."
		}
	}
	return res, ""
}

func BorrowersToJSONSting(brs []Borrower) string {
	bytes, _ := json.Marshal(brs)
	return string(bytes)
}

func BooksToJSONSting(bks []Book) string {
	bytes, _ := json.Marshal(bks)
	return string(bytes)
}

func libraryToString(bks []Book, brs []Borrower) string {
	return "Test Library: " +
		strconv.Itoa(len(bks)) + " books; " +
		strconv.Itoa(len(brs)) + " borrowers."
}

func StatusToString(bks []Book, brs []Borrower) string {
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
	sb.WriteString("\n--- End of Status Report ---\n\n")
	return sb.String()
}
