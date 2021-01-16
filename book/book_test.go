package book

import (
	"eatobin.com/totalbeginnergo/borrower"
	"testing"
)

var jsonStringBk1 = "{\"title\":\"Title11\",\"author\":\"Author11\",\"borrower\":null}"
var jsonStringBk2 = "{\"title\":\"Title1\",\"author\":\"Author1\",\"borrower\":{\"name\":\"Borrower2\",\"maxBooks\":2}}"
var br2Ptr = borrower.NewBorrower("Borrower2", 2)
var wantAvailS1 = "Title1 by Author11; Available"
var wantAvailS2 = "Title1 by Author1; Available"
var wantAvailS3 = "Title1 by Author1; Checked out to Borrower2"

func TestSetBookValues(t *testing.T) {
	title := "Title1"
	badBkT := NewBook("NoTitle", "Author1")
	gotBkT := String(SetTitle(badBkT, title))
	if gotBkT != wantAvailS {
		t.Fatalf("SetTitle(%v, %v) == %v, want %v", badBkT, title, gotBkT, wantAvailS)
	}
	author := "Author1"
	badBkA := NewBook("Title1", "NoAuthor")
	gotBkA := String(SetAuthor(badBkA, author))
	if gotBkA != wantAvailS {
		t.Fatalf("SetAuthor(%v, %v) == %v, want %v", badBkA, author, gotBkA, wantAvailS)
	}
	newBorrower := borrower.NewBorrower("Borrower2", 2)
	badBkB := Book{"Title1", "Author1", ZeroBorrower}
	gotBkB := String(SetBorrower(badBkB, newBorrower))
	if gotBkB != wantNotAvailS {
		t.Fatalf("SetBorrower(%v, %v) == %v, want %v", badBkB, newBorrower, gotBkB, wantNotAvailS)
	}
}
