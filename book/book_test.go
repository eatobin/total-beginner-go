package book

import (
	"fmt"
	"testing"

	"eatobin.com/totalbeginnergo/borrower"
)

var br2Ptr = borrower.NewBorrower("Borrower2", 2)
var bk1Ptr = NewBook("Title1", "Author1")
var bk2 = Book{"Title2", "Author2", &br2Ptr}

func TestBorrower(t *testing.T) {
	fmt.Printf("%+v\n", br2Ptr)
	fmt.Printf("%+v\n", bk1Ptr)
	fmt.Printf("%+v\n", bk2)
	var br = bk2.Borrower
	fmt.Printf("%+v\n", br)
}

//func TestNewBook(t *testing.T) {
//	gotT := bk1Ptr.Title
//	wantT := "Title1"
//	if gotT != wantT {
//		t.Fatalf("bk1Ptr.Title == %q, want %q", gotT, wantT)
//	}
//	gotA := bk1Ptr.Author
//	wantA := "Author1"
//	if gotA != wantA {
//		t.Fatalf("bk1Ptr.Author == %q, want %q", gotA, wantA)
//	}
//	gotBr := bk1Ptr.Borrower
//	if gotBr != nil {
//		t.Fatalf("bk1Ptr.Borrower == %v, want %v", gotBr, nil)
//	}
//}
//
//func TestSetBorrower(t *testing.T) {
//	bk1Ptr.SetBorrower(br2Ptr)
//	wantBk := Book{"Title1", "Author1",
//		&borrower.Borrower{Name: "Borrower2", MaxBooks: 2}}
//	//if *bk1Ptr != wantBk {
//	//	t.Fatalf("Book == %v, want %v", bk1Ptr, wantBk)
//		fmt.Printf("%+v\n", wantBk)
//	}
////}

//func TestSetBorrower(t *testing.T) {
//	bk1Ptr.SetBorrower(br2Ptr)
//	wantBk := Book{"Title1", "Author1",
//		&borrower.Borrower{Name: "Borrower2", MaxBooks: 2}}
//	if *bk1Ptr != wantBk {
//		t.Fatalf("Book == %v, want %v", bk1Ptr, wantBk)
//	}
//}

//func TestBookToString(t *testing.T) {
//	bk1Ptr.SetBorrower(borrower.Borrower{Name: "", MaxBooks: 0})
//	gotAvail := bk1Ptr.BookToString()
//	wantAvail := "Title1 by Author1; Available"
//	if gotAvail != wantAvail {
//		t.Fatalf("bk.BookToString() == %q, want %q", gotAvail, wantAvail)
//	}
//	bk1Ptr.SetBorrower(*br2Ptr)
//	gotNotAvail := bk1Ptr.BookToString()
//	wantNotAvail := "Title1 by Author1; Checked out to Borrower2"
//	if gotNotAvail != wantNotAvail {
//		t.Fatalf("bk.BookToString() == %q, want %q", gotNotAvail, wantNotAvail)
//	}
//}
