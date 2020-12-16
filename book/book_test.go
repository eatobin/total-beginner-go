package book

import (
	"testing"
)

//var br2Ptr = borrower.NewBorrower("Borrower2", 2)
//var bk1Ptr = NewBook("Title1", "Author1")
var wantAvailS = "Title1 by Author1; Available"
var wantNotAvailS = "Title1 by Author1; Checked out to Borrower2"

//func TestBookToString(t *testing.T) {
//	gotAvail := bk1Ptr.BookToString()
//	if gotAvail != wantAvail {
//		t.Fatalf("bk.BookToString() == %q, want %q", gotAvail, wantAvail)
//	}
//	bk1Ptr.SetBorrower(br2Ptr)
//	gotNotAvail := bk1Ptr.BookToString()
//	if gotNotAvail != wantNotAvail {
//		t.Fatalf("bk.BookToString() == %q, want %q", gotNotAvail, wantNotAvail)
//	}
//}

func TestSetBookValues(t *testing.T) {
	title := "Title1"
	badBkT := Book{"NoTitle", "Author1", nil}
	gotBkT := BkToString(setTitle(badBkT, title))
	if gotBkT != wantAvailS {
		t.Fatalf("setTitle(%v, %v) == %v, want %v", badBkT, title, gotBkT, wantAvailS)
	}
	author := "Author1"
	badBkA := Book{"Title1", "NoAuthor", nil}
	gotBkA := BkToString(setAuthor(badBkA, author))
	if gotBkA != wantAvailS {
		t.Fatalf("setAuthor(%v, %v) == %v, want %v", badBkA, author, gotBkA, wantAvailS)
	}

	//badTitle.SetTitle(title)

	//badAuthor := Book{"Title1", "NoAuthor", nil}
	//badAuthor.SetAuthor(author)

}
