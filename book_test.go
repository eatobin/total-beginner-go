package main

import (
	"testing"
)

var jsonStringBk1 = "{\"title\":\"Title11\",\"author\":\"Author11\",\"borrower\":null}"
var jsonStringBk2 = "{\"title\":\"Title1\",\"author\":\"Author1\",\"borrower\":{\"name\":\"Borrower2\",\"maxBooks\":2}}"
var br2Ptr = NewBorrower("Borrower2", 2)
var wantAvailS1 = "Title1 by Author11; Available"
var wantAvailS2 = "Title1 by Author1; Available"

func TestSetBookValues(t *testing.T) {
	badBkPtrNbS := NewBook("Title11", "Author11").BkToString()
	badBkPtr, _ := JsonStringToBook(jsonStringBk1)
	badBkPtrS := badBkPtr.BkToString()
	if badBkPtrNbS != badBkPtrS {
		t.Fatalf("NewBook (%v) and JsonStringToBook (%v) are not equal", badBkPtrNbS, badBkPtrS)
	}

	title := "Title1"
	badBkPtr.SetTitle(title)
	gotBkPtrS := badBkPtr.BkToString()
	if gotBkPtrS != wantAvailS1 {
		t.Fatalf("SetTitle(%v) == %v, want %v", title, gotBkPtrS, wantAvailS1)
	}
	author := "Author1"
	badBkPtr.SetAuthor(author)
	gotBkPtrS = badBkPtr.BkToString()
	if gotBkPtrS != wantAvailS2 {
		t.Fatalf("SetAuthor(%v) == %v, want %v", author, gotBkPtrS, wantAvailS2)
	}
	goodBkPtr, _ := JsonStringToBook(jsonStringBk2)
	goodBkPtrS := goodBkPtr.BkToString()
	badBkPtr.SetBorrower(br2Ptr)
	gotBkPtrS = badBkPtr.BkToString()
	if gotBkPtrS != goodBkPtrS {
		t.Fatalf("SetBorrower(%v) == %v, want %v", br2Ptr, gotBkPtrS, goodBkPtrS)
	}
}
