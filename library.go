package main

//
//import (
//	"encoding/json"
//	"errors"
//	"strconv"
//	"strings"
//)
//
//func containsBorrower(brs []*main.Borrower, br *main.Borrower) bool {
//	for _, b := range brs {
//		if *b == *br {
//			return true
//		}
//	}
//	return false
//}
//
//func containsBook(bks []*main.Book, bk *main.Book) bool {
//	for _, b := range bks {
//		if *b == *bk {
//			return true
//		}
//	}
//	return false
//}
//
//// AddBorrower adds a Borrower to a slice of Borrowers
//func AddBorrower(brs []*main.Borrower, br *main.Borrower) []*main.Borrower {
//	if containsBorrower(brs, br) {
//		return brs
//	}
//	return append(brs, br)
//}
//
//// AddBook adds a book to a slice of Books
//func AddBook(bks []*main.Book, bk *main.Book) []*main.Book {
//	if containsBook(bks, bk) {
//		return bks
//	}
//	return append(bks, bk)
//}
//
//// findBorrower finds a Borrower given a Name
//func findBorrower(n string, brs []*main.Borrower) (error, *main.Borrower) {
//	for _, br := range brs {
//		if br.Name == n {
//			return nil, br
//		}
//	}
//	return errors.New("did not find the requested borrower"), &main.Borrower{}
//}
//
//// findBook finds a Book given a Title
//func findBook(t string, bks []*main.Book) (int, error, *main.Book) {
//	for i, bk := range bks {
//		if bk.Title == t {
//			return i, nil, bk
//		}
//	}
//	return 0, errors.New("did not find the requested book"), &main.Book{}
//}
//
//// GetBooksForBorrower will find books given a Borrower and a slice of Books
//func getBooksForBorrower(br *main.Borrower, bks []*main.Book) []*main.Book {
//	nBks := make([]*main.Book, 0)
//	for _, bk := range bks {
//		if bk.Borrower == br {
//			nBks = append(nBks, bk)
//		}
//	}
//	return nBks
//}
//
//// numberBooksOut returns the # Books checked out to a Borrower
//func numberBooksOut(br *main.Borrower, bks []*main.Book) int {
//	return len(getBooksForBorrower(br, bks))
//}
//
//// notMaxedOut returns True if books out < max books
//func notMaxedOut(br *main.Borrower, bks []*main.Book) bool {
//	return numberBooksOut(br, bks) < br.MaxBooks
//}
//
//func bookNotOut(bk *main.Book) bool {
//	return bk.Borrower == nil
//}
//
//func bookOut(bk *main.Book) bool {
//	return bk.Borrower != nil
//}
//
//func CheckOut(n string, t string, brs []*main.Borrower, bks []*main.Book) []*main.Book {
//	errBr, mbr := findBorrower(n, brs)
//	i, errBk, mbk := findBook(t, bks)
//	if errBr == nil && errBk == nil && notMaxedOut(mbr, bks) && bookNotOut(mbk) {
//		bks[i].SetBorrower(mbr)
//		return bks
//	}
//	return bks
//}
//
//func CheckIn(t string, bks []*main.Book) []*main.Book {
//	i, errBk, mbk := findBook(t, bks)
//	if errBk == nil && bookOut(mbk) {
//		bks[i].SetBorrower(nil)
//		return bks
//	}
//	return bks
//}
//
//func jsonStringToBorrowers(js string) (error, []*main.Borrower) {
//	var res []*main.Borrower
//	err := json.Unmarshal([]byte(js), &res)
//	if err != nil {
//		return err, []*main.Borrower{}
//	}
//	for _, br := range res {
//		if br.Name == "" || br.MaxBooks == 0 {
//			err = errors.New("missing Borrower field value - borrowers list is empty")
//			return err, []*main.Borrower{}
//		}
//	}
//	return err, res
//}
//
//func jsonStringToBooks(js string) (error, []*main.Book) {
//	var res []*main.Book
//	err := json.Unmarshal([]byte(js), &res)
//	if err != nil {
//		return err, []*main.Book{}
//	}
//	for _, bk := range res {
//		if bk.Title == "" || bk.Author == "" || bk.Borrower.Name == "" || bk.Borrower.MaxBooks == 0 {
//			err = errors.New("missing Book field value - book list is empty")
//			return err, []*main.Book{}
//		}
//	}
//	return nil, res
//}
//
//func BorrowersToJSONSting(brs []*main.Borrower) string {
//	bytes, _ := json.MarshalIndent(brs, "", "  ")
//	return string(bytes)
//}
//
//func BooksToJSONSting(bks []*main.Book) string {
//	bytes, _ := json.MarshalIndent(bks, "", "  ")
//	return string(bytes)
//}
//
//func libraryToString(bks []*main.Book, brs []*main.Borrower) string {
//	return "Test Library: " +
//		strconv.Itoa(len(bks)) + " books; " +
//		strconv.Itoa(len(brs)) + " borrowers."
//}
//
//func StatusToString(bks []*main.Book, brs []*main.Borrower) string {
//	var sb strings.Builder
//	sb.WriteString("\n--- Status Report of Test Library ---\n\n")
//	sb.WriteString(libraryToString(bks, brs) + "\n\n")
//	for _, bk := range bks {
//		sb.WriteString(bk.BkToString() + "\n")
//	}
//	sb.WriteString("\n")
//	for _, br := range brs {
//		sb.WriteString(br.BorrowerToString() + "\n")
//	}
//	sb.WriteString("\n--- End of Status Report ---\n")
//	return sb.String()
//}
