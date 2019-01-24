package borrower

import (
	"testing"
)

var br1 = NewBorrower("Borrower1", 1)

func TestNewBorrower(t *testing.T) {
	gotN := br1.Name
	wantN := "Borrower1"
	if gotN != wantN {
		t.Fatalf("br1.Name == %q, want %q", gotN, wantN)
	}
	gotMB := br1.MaxBooks
	wantMB := 1
	if gotMB != wantMB {
		t.Fatalf("br1.MaxBooks == %v, want %v", gotMB, wantMB)
	}
}

func TestSetValues(t *testing.T) {
	n := "Borrower1"
	badBr := Borrower{"Jack", 1}
	badBr.SetName(n)
	if badBr != *br1 {
		t.Fatalf("br.SetName(%q) == %v, want %v", n, badBr, *br1)
	}
	mb := 1
	badBrMB := Borrower{"Borrower1", 11}
	badBrMB.SetMaxBooks(mb)
	if badBrMB != *br1 {
		t.Fatalf("br.SetMaxBooks(%v) == %v, want %v", mb, badBrMB, *br1)
	}
}

func TestBorrowerToString(t *testing.T) {
	gotS := br1.BorrowerToString()
	wantS := "Borrower1 (1 books)"
	if gotS != wantS {
		t.Fatalf("br.BorrowerToString() == %v, want %v", gotS, wantS)
	}
}
