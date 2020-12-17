package book

import (
	"eatobin.com/totalbeginnergo/borrower"
	"testing"
)

var wantAvailS = "Title1 by Author1; Available"
var wantNotAvailS = "Title1 by Author1; Checked out to Borrower2"

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
	newBorrower := borrower.NewBorrower("Borrower2", 2)
	badBkB := Book{"Title1", "Author1", nil}
	gotBkB := BkToString(SetBorrower(badBkB, &newBorrower))
	if gotBkB != wantNotAvailS {
		t.Fatalf("SetBorrower(%v, %v) == %v, want %v", badBkB, &newBorrower, gotBkB, wantNotAvailS)
	}
}
