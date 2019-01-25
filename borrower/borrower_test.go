package borrower

import (
	"testing"
)

var br1Ptr = NewBorrower("Borrower1", 1)

func TestNewBorrower(t *testing.T) {
	gotN := br1Ptr.Name
	wantN := "Borrower1"
	if gotN != wantN {
		t.Fatalf("br1Ptr.Name == %q, want %q", gotN, wantN)
	}
	gotMB := br1Ptr.MaxBooks
	wantMB := 1
	if gotMB != wantMB {
		t.Fatalf("br1Ptr.MaxBooks == %v, want %v", gotMB, wantMB)
	}
}

func TestSetValues(t *testing.T) {
	n := "Borrower1"
	badBr := Borrower{"Jack", 1}
	badBr.SetName(n)
	if badBr != *br1Ptr {
		t.Fatalf("br.SetName(%q) == %v, want %v", n, badBr, *br1Ptr)
	}
	mb := 1
	badBrMB := Borrower{"Borrower1", 11}
	badBrMB.SetMaxBooks(mb)
	if badBrMB != *br1Ptr {
		t.Fatalf("br.SetMaxBooks(%v) == %v, want %v", mb, badBrMB, *br1Ptr)
	}
}

func TestBorrowerToString(t *testing.T) {
	gotS := br1Ptr.BorrowerToString()
	wantS := "Borrower1 (1 books)"
	if gotS != wantS {
		t.Fatalf("br.BorrowerToString() == %v, want %v", gotS, wantS)
	}
}
