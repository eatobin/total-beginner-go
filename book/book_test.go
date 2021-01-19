package book

import (
	"eatobin.com/totalbeginnergo/borrower"
	"testing"
)

var bk1, _ = JsonStringToBook("{\"title\":\"Title1\",\"author\":\"Author1\"}")
var newBorrower = borrower.NewBorrower("Borrower1", 1)
var bk2, _ = JsonStringToBook("{\"title\":\"Title1\",\"author\":\"Author1\",\"borrower\":null}")
var bk3, _ = JsonStringToBook("{\"title\":\"Title1\",\"author\":\"Author1\",\"borrower\":{\"name\":\"Borrower1\",\"maxBooks\":1}}")
var wantAvail = "Title1 by Author1; Available"
var wantCheckedOut = "Title1 by Author1; Checked out to Borrower1"
var badTitle, _ = JsonStringToBook("{\"title\":\"Title1X\",\"author\":\"Author1\"}")
var badAuthor, _ = JsonStringToBook("{\"title\":\"Title1\",\"author\":\"Author1X\"}")

func TestSetTitle(t *testing.T) {
	title := "Title1"
	gotBkT := String(SetTitle(badTitle, title))
	if gotBkT != wantAvail {
		t.Fatalf("SetTitle(%v, %v) == %v, want %v", badTitle, title, gotBkT, wantAvail)
	}
}

func TestSetAuthor(t *testing.T) {
	author := "Author1"
	gotBkA := String(SetAuthor(badAuthor, author))
	if gotBkA != wantAvail {
		t.Fatalf("SetAuthor(%v, %v) == %v, want %v", badAuthor, author, gotBkA, wantAvail)
	}
}

func TestSetBorrower(t *testing.T) {
	cases := []struct {
		bk           Book
		br           *borrower.Borrower
		wantBkString string
	}{
		{bk1, &newBorrower, wantCheckedOut},
		{bk2, &newBorrower, wantCheckedOut},
		{bk3, nil, wantAvail},
	}
	for _, c := range cases {
		gotBkString := String(SetBorrower(c.bk, c.br))
		if gotBkString != c.wantBkString {
			t.Fatalf("SetBorrower(%v, %v) == %v, want %v", c.bk, c.br, gotBkString, c.wantBkString)
		}
	}
}
