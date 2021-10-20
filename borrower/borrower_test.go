package borrower

import "testing"

var badName, _ = JsonStringToBorrower("{\"name\":\"Borrower1X\",\"maxBooks\":1}")
var badMB, _ = JsonStringToBorrower("{\"name\":\"Borrower1\",\"maxBooks\":19}")
var wantBr = "Borrower1 [1 books]"
var jsonString = "{\"name\":\"Borrower1X\",\"maxBooks\":1}"
var brA = Borrower{Name: "Borrower1", MaxBooks: 1}
var brB = Borrower{Name: "Borrower1", MaxBooks: 1}
var brC = Borrower{Name: "Nope", MaxBooks: 1}
var brD = Borrower{Name: "Borrower1", MaxBooks: 111}

func TestEqual(t *testing.T) {
	if !brA.Equal(brB) {
		t.Fatalf("(%v) Equal(%v) == %t, want %t", brA, brB, false, true)
	}
	if brB.Equal(brC) {
		t.Fatalf("(%v) Equal(%v) == %t, want %t", brB, brC, true, false)
	}
	if brB.Equal(brD) {
		t.Fatalf("(%v) Equal(%v) == %t, want %t", brB, brD, true, false)
	}
}

func TestSetName(t *testing.T) {
	goodName := "Borrower1"
	gotBr := badName.SetName(goodName)
	if !gotBr.Equal(brA) {
		t.Fatalf("(%v) SetName(%v) == %v, want %v", badName, goodName, gotBr, wantBr)
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
