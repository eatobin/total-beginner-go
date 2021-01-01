package main

import (
	"testing"
)

var jsonStringBr = "{\"name\":\"Borrower1\",\"maxBooks\":11}"
var wantS1 = "Borrower1 (11 books)"
var wantS2 = "Borrower1 (1 books)"

func TestSetBorrowerValues(t *testing.T) {
	badBrPtr, _ := JsonStringToBorrower(jsonStringBr)
	n := "Borrower1"
	badBrPtr.SetName(n)
	gotBrPtrS := badBrPtr.String()
	if gotBrPtrS != wantS1 {
		t.Fatalf("SetName(%v) == %v, want %v", n, gotBrPtrS, wantS1)
	}
	mb := 1
	badBrPtr.SetMaxBooks(mb)
	gotBrPtrS = badBrPtr.String()
	if gotBrPtrS != wantS2 {
		t.Fatalf("SetMaxBooks(%v) == %v, want %v", mb, gotBrPtrS, wantS2)
	}
}
