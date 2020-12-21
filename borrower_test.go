package main

import "testing"

var wantS = "Borrower1 (1 books)"

func TestSetBorrowerValues(t *testing.T) {
	n := "Borrower1"
	badBrN := NewBorrower("Jack", 1)
	gotBrN := BrToString(setName(badBrN, n))
	if gotBrN != wantS {
		t.Fatalf("setName(%v, %v) == %v, want %v", badBrN, n, gotBrN, wantS)
	}
	mb := 1
	badBrMB := Borrower{"Borrower1", 11}
	gotBrMB := BrToString(setMaxBooks(badBrMB, mb))
	if gotBrMB != wantS {
		t.Fatalf("setMaxBooks(%v, %v) == %v, want %v", badBrMB, mb, gotBrMB, wantS)
	}
}
