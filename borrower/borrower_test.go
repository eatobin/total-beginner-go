package borrower

import (
	"testing"
)

var badName, _ = JsonStringToBorrower("{\"name\":\"Borrower1X\",\"maxBooks\":1}")
var badMB, _ = JsonStringToBorrower("{\"name\":\"Borrower1\",\"maxBooks\":19}")
var wantBr = "Borrower1 [1 books]"
var jsonString = "{\"name\":\"Borrower1X\",\"maxBooks\":1}"

func TestSetName(t *testing.T) {
	goodName := "Borrower1"
	gotBrN := badName.SetName(goodName).String()
	if gotBrN != wantBr {
		t.Fatalf("(%v) SetName(%v) == %v, want %v", badName, goodName, gotBrN, wantBr)
	}
}

func TestSetMaxBooks(t *testing.T) {
	maxBooks := 1
	gotBrMB := Borrower.String(Borrower.SetMaxBooks(badMB, maxBooks))
	if gotBrMB != wantBr {
		t.Fatalf("SetMaxBooks(%v, %v) == %v, want %v", badMB, maxBooks, gotBrMB, wantBr)
	}
}

func TestBrToJsonString(t *testing.T) {
	gotJsonString, _ := Borrower.BrToJsonString(badName)
	if gotJsonString != jsonString {
		t.Fatalf("BrToJsonString(%v) == %v, want %v", badName, gotJsonString, jsonString)
	}
}
