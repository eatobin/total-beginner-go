package main

import (
	"testing"
)

var badBrPtr = NewBorrower("Jack", 11)
var wantS1 = "Borrower1 (11 books)"
var wantS2 = "Borrower1 (1 books)"
var jsonStringBr = "{\"name\":\"Borrower1\",\"maxBooks\":1}"

func TestSetBorrowerValues(t *testing.T) {
	n := "Borrower1"
	badBrPtr.SetName(n)
	gotBrPtr := badBrPtr.BrToString()
	if gotBrPtr != wantS1 {
		t.Fatalf("SetName(%v) == %v, want %v", n, gotBrPtr, wantS1)
	}
	mb := 1
	badBrPtr.SetMaxBooks(mb)
	gotBrPtr = badBrPtr.BrToString()
	if gotBrPtr != wantS2 {
		t.Fatalf("SetMaxBooks(%v) == %v, want %v", mb, gotBrPtr, wantS2)
	}
}

func TestConvertFromJSON(t *testing.T) {
	gotBrPtr, _ := JsonStringToBorrower(jsonStringBr)
	if gotBrPtr.BrToString() != badBrPtr.BrToString() {
		t.Fatalf("JsonStringToBorrower(%v) == %v, want %v", jsonStringBr, gotBrPtr, badBrPtr)
	}
}
