package main

import (
	"testing"
)

var br2Ptr = NewBorrower("Borrower2", 2)
var badBkPtr = NewBook("Title11", "Author11")
var wantAvailS1 = "Title1 by Author11; Available"
var wantAvailS2 = "Title1 by Author1; Available"
var wantNotAvail = "Title1 by Author1; Checked out to Borrower2"
var newBkPtr = NewBook("Title1", "Author1")
var jsonStringBk1 = "{\"title\":\"Title1\",\"author\":\"Author1\",\"borrower\":null}"
var jsonStringBk2 = "{\"title\":\"Title1\",\"author\":\"Author1\",\"borrower\":{\"name\":\"Borrower2\",\"maxBooks\":2}}"

func TestSetBookValues(t *testing.T) {
	title := "Title1"
	badBkPtr.SetTitle(title)
	gotBkPtr := badBkPtr.BkToString()
	if gotBkPtr != wantAvailS1 {
		t.Fatalf("SetTitle(%v) == %v, want %v", title, gotBkPtr, wantAvailS1)
	}
	author := "Author1"
	badBkPtr.SetAuthor(author)
	gotBkPtr = badBkPtr.BkToString()
	if gotBkPtr != wantAvailS2 {
		t.Fatalf("SetAuthor(%v) == %v, want %v", author, gotBkPtr, wantAvailS2)
	}
	badBkPtr.SetBorrower(br2Ptr)
	gotBkPtr = badBkPtr.BkToString()
	if gotBkPtr != wantNotAvail {
		t.Fatalf("SetBorrower(%v) == %v, want %v", br2Ptr, gotBkPtr, wantNotAvail)
	}
}

func TestConvertBookFromJSON(t *testing.T) {
	gotBkPtr, _ := JsonStringToBook(jsonStringBk1)
	if gotBkPtr.BkToString() != newBkPtr.BkToString() {
		t.Fatalf("JsonStringToBook(%v) == %v, want %v", jsonStringBk1, gotBkPtr.BkToString(), newBkPtr.BkToString())
	}
	gotBkPtr2, _ := JsonStringToBook(jsonStringBk2)
	if gotBkPtr2.BkToString() != badBkPtr.BkToString() {
		t.Fatalf("JsonStringToBook(%v) == %v, want %v", jsonStringBk2, gotBkPtr2.BkToString(), badBkPtr.BkToString())
	}
}
