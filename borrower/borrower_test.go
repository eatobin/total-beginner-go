package borrower

import (
	"testing"
)

var br1Ptr = NewBorrower("Borrower1", 1)

func TestBorrowerToString(t *testing.T) {
	gotS := br1Ptr.BorrowerToString()
	wantS := "Borrower1 (1 books)"
	if gotS != wantS {
		t.Fatalf("br.BorrowerToString() == %v, want %v", gotS, wantS)
	}
}

func TestSetValues(t *testing.T) {
	n := "Borrower1"
	badBr := Borrower{"Jack", 1}
	badBr.SetName(n)
	wantS := "Borrower1 (1 books)"
	if badBr.BorrowerToString() != wantS {
		t.Fatalf("br.SetName(%q) == %v, want %v", n, badBr.BorrowerToString(), wantS)
	}
	mb := 1
	badBrMB := Borrower{"Borrower1", 11}
	badBrMB.SetMaxBooks(mb)
	if badBrMB.BorrowerToString() != wantS {
		t.Fatalf("br.SetMaxBooks(%v) == %v, want %v", mb, badBrMB.BorrowerToString(), wantS)
	}
}
