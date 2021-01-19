package borrower

import (
	"testing"
)

//var br, _ = JsonStringToBorrower("{\"name\":\"Borrower1\",\"maxBooks\":1}")
var badName, _ = JsonStringToBorrower("{\"name\":\"Borrower1X\",\"maxBooks\":1}")

//var badMB, _ = JsonStringToBorrower("{\"name\":\"Borrower1\",\"maxBooks\":19}")
var wantBr = "Borrower1 (1 books)"

func TestSetName(t *testing.T) {
	name := "Borrower1"
	gotBrN := String(SetName(badName, name))
	if gotBrN != wantBr {
		t.Fatalf("SetName(%v, %v) == %v, want %v", badName, name, gotBrN, wantBr)
	}
}

//func TestSetBorrowerValues(t *testing.T) {
//
//	mb := 1
//	gotBrS2 := String(SetMaxBooks(badBr, mb))
//	if gotBrS2 != wantS2 {
//		t.Fatalf("SetMaxBooks(%v) == %v, want %v", mb, gotBrS2, wantS2)
//	}
//	nb := NewBorrower("Borrower11", 11)
//	gotBrString, _ := BrToJsonString(nb)
//	if gotBrString != jsonStringBr {
//		t.Fatalf("BrToJsonString(%v) == %v, want %v", nb, gotBrString, jsonStringBr)
//	}
//}
