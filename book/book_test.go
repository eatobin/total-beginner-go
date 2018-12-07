package book

import (
	"testing"

	"github.com/eatobin/totalbeginnergo/borrower"
)

var br2 = borrower.NewBorrower("Borrower2", 2)
var bk1 = NewBook("Title1", "Author1")

func TestNewBook(t *testing.T) {
	gotT := bk1.Title
	wantT := "Title1"
	if gotT != wantT {
		t.Fatalf("bk1.Title == %q, want %q", gotT, wantT)
	}
	gotA := bk1.Author
	wantA := "Author1"
	if gotA != wantA {
		t.Fatalf("bk1.Author == %q, want %q", gotA, wantA)
	}
	gotBr := bk1.Borrower
	wantBr := borrower.Borrower{"", 0}
	if gotBr != wantBr {
		t.Fatalf("bk1.Borrower == %v, want %v", gotBr, wantBr)
	}
	bk1.SetBorrower(br2)
	gotBk := bk1
	wantBk := Book{"Title1", "Author1", borrower.Borrower{"Borrower2", 2}}
	if gotBk != wantBk {
		t.Fatalf("Book == %v, want %v", gotBk, wantBk)
	}
}

//func TestBookToString(t *testing.T) {
//	bk1.SetBorrower(borrower.Borrower{Name: "NoName", MaxBooks: -1})
//	gotAvail := bk1.BookToString()
//	wantAvail := "Title1 by Author1; Available"
//	if gotAvail != wantAvail {
//		t.Fatalf("bk.BookToString() == %q, want %q", gotAvail, wantAvail)
//	}
//	bk1.SetBorrower(br2)
//	gotNotAvail := bk1.BookToString()
//	wantNotAvail := "Title1 by Author1; Checked out to Borrower2"
//	if gotNotAvail != wantNotAvail {
//		t.Fatalf("bk.BookToString() == %q, want %q", gotNotAvail, wantNotAvail)
//	}
//}
