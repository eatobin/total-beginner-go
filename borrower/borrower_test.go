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
	goodMB := 1
	gotBrMB := badMB.SetMaxBooks(goodMB).String()
	if gotBrMB != wantBr {
		t.Fatalf("(%v) SetMaxBooks(%v) == %v, want %v", badMB, goodMB, gotBrMB, wantBr)
	}
}

func TestBrToJsonString(t *testing.T) {
	gotJsonString, _ := badName.BrToJsonString()
	if gotJsonString != jsonString {
		t.Fatalf("(%v) BrToJsonString() == %v, want %v", badName, gotJsonString, jsonString)
	}
}
