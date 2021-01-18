package book

import (
	"eatobin.com/totalbeginnergo/borrower"
	"testing"
)

var jsonStringBk1 = "{\"title\":\"Title11\",\"author\":\"Author11\",\"borrower\":null}"
var jsonStringBk2 = "{\"title\":\"Title11\",\"author\":\"Author11\",\"borrower\":{\"name\":\"Borrower2\",\"maxBooks\":2}}"
var wantAvailS1 = "Title1 by Author11; Available"
var wantAvailS2 = "Title11 by Author1; Available"

func TestSetBookValues(t *testing.T) {
	badBkAvail, _ := JsonStringToBook(jsonStringBk1)
	title := "Title1"
	gotBkT := String(SetTitle(badBkAvail, title))
	if gotBkT != wantAvailS1 {
		t.Fatalf("SetTitle(%v, %v) == %v, want %v", badBkAvail, title, gotBkT, wantAvailS1)
	}
	author := "Author1"
	gotBkA := String(SetAuthor(badBkAvail, author))
	if gotBkA != wantAvailS2 {
		t.Fatalf("SetAuthor(%v, %v) == %v, want %v", badBkAvail, author, gotBkA, wantAvailS2)
	}
	bkNotAvail, _ := JsonStringToBook(jsonStringBk2)
	wantNotAvailS := String(bkNotAvail)
	br2 := borrower.NewBorrower("Borrower2", 2)
	gotBkB := String(SetBorrower(badBkAvail, &br2))
	if gotBkB != wantNotAvailS {
		t.Fatalf("SetBorrower(%v, %v) == %v, want %v", badBkAvail, br2, gotBkB, wantNotAvailS)
	}
}
