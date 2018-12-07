package book

import (
	"fmt"
	"testing"

	"github.com/eatobin/totalbeginnergo/borrower"
)

var pBr2 = borrower.NewBorrower("Borrower2", 2)
var pBk1 = NewBook("Title1", "Author1")

func TestNewBook(t *testing.T) {
	fmt.Println(pBk1)
	fmt.Println(*pBk1)
	gotT := pBk1.Title
	wantT := "Title1"
	if gotT != wantT {
		t.Fatalf("bk1.Title == %q, want %q", gotT, wantT)
	}
	gotA := pBk1.Author
	wantA := "Author1"
	if gotA != wantA {
		t.Fatalf("bk1.Author == %q, want %q", gotA, wantA)
	}
	gotBr := pBk1.Borrower
	wantBr := borrower.Borrower{}
	if gotBr != wantBr {
		t.Fatalf("bk1.Borrower == %v, want %v", gotBr, wantBr)
	}
}

func TestSetBorrower(t *testing.T) {
	pBk1.SetBorrower(pBr2)
	gotBk := pBk1
	wantBk := Book{"Title1", "Author1",
		borrower.Borrower{Name: "Borrower2", MaxBooks: 2}}
	if *gotBk != wantBk {
		t.Fatalf("Book == %v, want %v", gotBk, wantBk)
	}
}

func TestBookToString(t *testing.T) {
	pBk1.SetBorrower(&borrower.Borrower{Name: "", MaxBooks: 0})
	gotAvail := pBk1.BookToString()
	wantAvail := "Title1 by Author1; Available"
	if gotAvail != wantAvail {
		t.Fatalf("bk.BookToString() == %q, want %q", gotAvail, wantAvail)
	}
	pBk1.SetBorrower(pBr2)
	gotNotAvail := pBk1.BookToString()
	wantNotAvail := "Title1 by Author1; Checked out to Borrower2"
	if gotNotAvail != wantNotAvail {
		t.Fatalf("bk.BookToString() == %q, want %q", gotNotAvail, wantNotAvail)
	}
}
