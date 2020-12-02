package book

import (
	"eatobin.com/totalbeginnergo/borrower"
	"testing"
)

var br2Ptr = borrower.NewBorrower("Borrower2", 2)
var bk1Ptr = NewBook("Title1", "Author1")
var wantAvail = "Title1 by Author1; Available"
var wantNotAvail = "Title1 by Author1; Checked out to Borrower2"

func TestBookToString(t *testing.T) {
	gotAvail := bk1Ptr.BookToString()
	if gotAvail != wantAvail {
		t.Fatalf("bk.BookToString() == %q, want %q", gotAvail, wantAvail)
	}
	bk1Ptr.SetBorrower(br2Ptr)
	gotNotAvail := bk1Ptr.BookToString()
	if gotNotAvail != wantNotAvail {
		t.Fatalf("bk.BookToString() == %q, want %q", gotNotAvail, wantNotAvail)
	}
}

func TestSetBookValues(t *testing.T) {
	title := "Title1"
	author := "Author1"
	badTitle := Book{"NoTitle", "Author1", nil}
	badTitle.SetTitle(title)
	if badTitle.BookToString() != wantAvail {
		t.Fatalf("bk.SetTitle(%q) == %v, want %v", title, badTitle.BookToString(), wantAvail)
	}
	badAuthor := Book{"Title1", "NoAuthor", nil}
	badAuthor.SetAuthor(author)
	if badAuthor.BookToString() != wantAvail {
		t.Fatalf("bk.SetAuthor(%v) == %v, want %v", author, badAuthor.BookToString(), wantAvail)
	}
}
