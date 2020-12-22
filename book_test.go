package main

import (
	"testing"
)

var wantAvailS = "Title1 by Author1; Available"
var wantNotAvailS = "Title1 by Author1; Checked out to Borrower2"

func TestSetBookValues(t *testing.T) {
	title := "Title1"
	badBkT := NewBook("NoTitle", "Author1")
	gotBkT := BkToString(SetTitle(badBkT, title))
	if gotBkT != wantAvailS {
		t.Fatalf("SetTitle(%v, %v) == %v, want %v", badBkT, title, gotBkT, wantAvailS)
	}
	author := "Author1"
	badBkA := NewBook("Title1", "NoAuthor")
	gotBkA := BkToString(SetAuthor(badBkA, author))
	if gotBkA != wantAvailS {
		t.Fatalf("SetAuthor(%v, %v) == %v, want %v", badBkA, author, gotBkA, wantAvailS)
	}
	newBorrower := NewBorrower("Borrower2", 2)
	badBkB := Book{"Title1", "Author1", ZeroBorrower}
	gotBkB := BkToString(SetBorrower(badBkB, newBorrower))
	if gotBkB != wantNotAvailS {
		t.Fatalf("SetBorrower(%v, %v) == %v, want %v", badBkB, newBorrower, gotBkB, wantNotAvailS)
	}
}
