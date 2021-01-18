package borrower

import (
	"testing"
)

var jsonStringBr = "{\"name\":\"Borrower11\",\"maxBooks\":11}"
var wantS = String(NewBorrower("Borrower1", 11))
var wantS2 = String(NewBorrower("Borrower11", 1))

func TestSetBorrowerValues(t *testing.T) {
	badBr, _ := JsonStringToBorrower(jsonStringBr)
	n := "Borrower1"
	gotBrS := String(SetName(badBr, n))
	if gotBrS != wantS {
		t.Fatalf("SetName(%v, %v) == %v, want %v", badBr, n, gotBrS, wantS)
	}
	mb := 1
	gotBrS2 := String(SetMaxBooks(badBr, mb))
	if gotBrS2 != wantS2 {
		t.Fatalf("SetMaxBooks(%v) == %v, want %v", mb, gotBrS2, wantS2)
	}
	nb := NewBorrower("Borrower111", 111)
	gotBrString, _ := ToJsonString(nb)
	wantBrString := "{\"name\":\"Borrower111\",\"maxBooks\":111}"
	if gotBrString != wantBrString {
		t.Fatalf("ToJsonString(%v) == %v, want %v", nb, gotBrString, wantBrString)
	}
}
