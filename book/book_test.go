package book

import (
	"testing"

	"github.com/eatobin/totalbeginnergo/borrower"
)

var br2 = borrower.NewBorrower("Borrower2", 2)
var bk1 = MakeBook("Title1", "Author1")

func TestMakeBook(t *testing.T) {
	gotT := bk1.Title
	wantT := "Title1"
	if gotT != wantT {
		t.Fatalf("bk1lib.Title == %q, want %q", gotT, wantT)
	}
	gotA := bk1.Author
	wantA := "Author1"
	if gotA != wantA {
		t.Fatalf("bk1lib.Author == %q, want %q", gotA, wantA)
	}
	gotBrIsNot := bk1.Borrower
	wantBrIsNot := borrower.Borrower{Name: "NoName", MaxBooks: -1}
	if gotBrIsNot != wantBrIsNot {
		t.Fatalf("bk1lib.Borrower == %v, want %v", gotBrIsNot, wantBrIsNot)
	}
	bk1.SetBorrower(br2)
	gotBrIs := bk1.Borrower
	wantBrIs := br2
	if gotBrIs != wantBrIs {
		t.Fatalf("bk1lib.Borrower == %v, want %v", gotBrIs, wantBrIs)
	}
}

func TestBookToString(t *testing.T) {
	bk1.SetBorrower(borrower.Borrower{Name: "NoName", MaxBooks: -1})
	gotAvail := bk1.BookToString()
	wantAvail := "Title1 by Author1; Available"
	if gotAvail != wantAvail {
		t.Fatalf("bk.BookToString() == %q, want %q", gotAvail, wantAvail)
	}
	bk1.SetBorrower(br2)
	gotNotAvail := bk1.BookToString()
	wantNotAvail := "Title1 by Author1; Checked out to Borrower2"
	if gotNotAvail != wantNotAvail {
		t.Fatalf("bk.BookToString() == %q, want %q", gotNotAvail, wantNotAvail)
	}
}
