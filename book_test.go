package main

import (
	"testing"
)

//var br2Ptr = NewBorrower("Borrower2", 2)
var badBkPtr = NewBook("Title11", "Author11")
var wantAvailS1 = "Title1 by Author11; Available"
var wantAvailS2 = "Title1 by Author1; Available"

//var wantNotAvail = "Title1 by Author1; Checked out to Borrower2"

func TestSetBookValues(t *testing.T) {
	title := "Title1"
	badBkPtr.SetTitle(title)
	gotBkPtr := badBkPtr.BkToString()
	if gotBkPtr != wantAvailS1 {
		t.Fatalf("SetTitle(%q) == %v, want %v", title, gotBkPtr, wantAvailS1)
	}
	author := "Author1"
	badBkPtr.SetAuthor(author)
	gotBkPtr = badBkPtr.BkToString()
	if gotBkPtr != wantAvailS2 {
		t.Fatalf("SetAuthor(%q) == %v, want %v", author, gotBkPtr, wantAvailS2)
	}
}
