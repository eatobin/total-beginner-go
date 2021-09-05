package library

import (
	"eatobin.com/totalbeginnergo/book"
	"eatobin.com/totalbeginnergo/borrower"
	"errors"
)

//var ZeroBorrower borrower.Borrower
//var ZeroBook book.Book
//var ZeroBorrowers []borrower.Borrower
//var ZeroBooks []book.Book

type Borrower = borrower.Borrower
type Borrowers = []*borrower.Borrower
type Book = book.Book
type Books = []*book.Book

func containsBorrower(brs Borrowers, br Borrower) bool {
	for _, b := range brs {
		if *b == br {
			return true
		}
	}
	return false
}

func containsBook(bks Books, bk Book) bool {
	for _, b := range bks {
		if *b == bk {
			return true
		}
	}
	return false
}

// AddBorrower adds a Borrower pointer to a slice of Borrower pointers
func AddBorrower(brs Borrowers, br Borrower) Borrowers {
	if containsBorrower(brs, br) {
		return brs
	}
	return append(brs, &br)
}

// AddBook adds a Book pointer to a slice of Book pointers
func AddBook(bks Books, bk Book) Books {
	if containsBook(bks, bk) {
		return bks
	}
	return append(bks, &bk)
}

// removeBook removes a Book pointer from a slice of Book pointers
func removeBook(bks Books, bk Book) Books {
	nBks := make(Books, 0)
	for _, nBk := range bks {
		if *nBk != bk {
			nBks = append(nBks, nBk)
		}
	}
	return nBks
}

// findBorrower finds a Borrower pointer given a Name
func findBorrower(n string, brs Borrowers) (*Borrower, error) {
	for _, br := range brs {
		if br.Name == n {
			return br, nil
		}
	}
	return nil, errors.New("did not find the requested borrower")
}

// findBook finds a Book pointer given a Title
func findBook(t string, bks Books) (*Book, error) {
	for _, bk := range bks {
		if bk.Title == t {
			return bk, nil
		}
	}
	return nil, errors.New("did not find the requested book")
}

// getBooksForBorrower will find books given a Borrower and a slice of Book pointers
func getBooksForBorrower(br Borrower, bks Books) Books {
	nBks := make(Books, 0)
	for _, bk := range bks {
		if bk.Borrower != nil {
			if *bk.Borrower == br {
				nBks = append(nBks, bk)
			}
		}
	}
	return nBks
}

// numberBooksOut returns the # Books checked out to a Borrower
func numberBooksOut(br Borrower, bks Books) int {
	return len(getBooksForBorrower(br, bks))
}

// notMaxedOut returns True if books out < max books
func notMaxedOut(br Borrower, bks Books) bool {
	return numberBooksOut(br, bks) < br.MaxBooks
}

func bookNotOut(bk Book) bool {
	return bk.Borrower == nil
}

func bookOut(bk Book) bool {
	return bk.Borrower != nil
}

func CheckOut(n string, t string, brs Borrowers, bks Books) Books {
	mbr, errBr := findBorrower(n, brs)
	mbk, errBk := findBook(t, bks)
	if errBr == nil && errBk == nil && notMaxedOut(*mbr, bks) && bookNotOut(*mbk) {
		newBook := mbk.SetBorrower(mbr)
		fewerBooks := removeBook(bks, *mbk)
		return AddBook(fewerBooks, newBook)
	}
	return bks
}

//func CheckIn(t string, bks []book.Book) []book.Book {
//	mbk, errBk := findBook(t, bks)
//	if errBk == nil && bookOut(mbk) {
//		newBook := book.Book.SetBorrower(mbk, nil)
//		fewerBooks := removeBook(mbk, bks)
//		return AddBook(fewerBooks, newBook)
//	}
//	return bks
//}
//
//func JsonStringToBorrowers(borrowersString string) ([]borrower.Borrower, error) {
//	borrowers := ZeroBorrowers
//	err := json.Unmarshal([]byte(borrowersString), &borrowers)
//	if err != nil {
//		return ZeroBorrowers, err
//	}
//	for _, br := range borrowers {
//		if br.Name == "" || br.MaxBooks == 0 {
//			err = errors.New("missing Borrower field value - borrowers list is empty")
//			return ZeroBorrowers, err
//		}
//	}
//	return borrowers, err
//}
//
//func JsonStringToBooks(bookString string) ([]book.Book, error) {
//	books := ZeroBooks
//	err := json.Unmarshal([]byte(bookString), &books)
//	if err != nil {
//		return ZeroBooks, err
//	}
//	for _, bk := range books {
//		if bk.Title == "" || bk.Author == "" {
//			err = errors.New("missing Book field value - book list is empty")
//			return ZeroBooks, err
//		}
//		if bk.Borrower != nil {
//			if bk.Borrower.Name == "" || bk.Borrower.MaxBooks == 0 {
//				err = errors.New("missing Borrower field value - book list is empty")
//				return ZeroBooks, err
//			}
//		}
//	}
//	return books, err
//}
//
//func BorrowersToJSONSting(brs []borrower.Borrower) (string, error) {
//	bytes, err := json.Marshal(brs)
//	return string(bytes), err
//}
//
//func BooksToJSONSting(bks []book.Book) (string, error) {
//	bytes, err := json.Marshal(bks)
//	return string(bytes), err
//}
//
//func libraryToString(bks []book.Book, brs []borrower.Borrower) string {
//	return "Test Library: " +
//		strconv.Itoa(len(bks)) + " books; " +
//		strconv.Itoa(len(brs)) + " borrowers."
//}
//
//func StatusToString(bks []book.Book, brs []borrower.Borrower) string {
//	var sb strings.Builder
//	sb.WriteString("\n--- Status Report of Test Library ---\n\n")
//	sb.WriteString(libraryToString(bks, brs) + "\n\n")
//	for _, bk := range bks {
//		sb.WriteString(book.Book.String(bk) + "\n")
//	}
//	sb.WriteString("\n")
//	for _, br := range brs {
//		sb.WriteString(borrower.Borrower.String(br) + "\n")
//	}
//	sb.WriteString("\n--- End of Status Report ---\n")
//	return sb.String()
//}
