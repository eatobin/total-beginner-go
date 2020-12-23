package main

import (
	"testing"
)

var wantS1 = "Borrower1 (11 books)"
var wantS2 = "Borrower1 (1 books)"

func TestSetBorrowerValues(t *testing.T) {
	n := "Borrower1"
	badBrPtr := NewBorrower("Jack", 11)
	badBrPtr.SetName(n)
	gotBrPtr := badBrPtr.BrToString()
	if gotBrPtr != wantS1 {
		t.Fatalf("SetName(%q) == %v, want %v", n, gotBrPtr, wantS1)
	}
	mb := 1
	badBrPtr.SetMaxBooks(mb)
	gotBrPtr = badBrPtr.BrToString()
	if gotBrPtr != wantS2 {
		t.Fatalf("SetMaxBooks(%v) == %v, want %v", mb, gotBrPtr, wantS2)
	}
}
