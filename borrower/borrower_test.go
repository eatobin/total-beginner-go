package borrower

import (
	"fmt"
	"testing"
)

var pBr1 = NewBorrower("Borrower1", 1)

func TestNewBorrower(t *testing.T) {
	fmt.Println(pBr1)
	fmt.Println(*pBr1)
	gotN := pBr1.Name
	wantN := "Borrower1"
	if gotN != wantN {
		t.Fatalf("br1.Name == %q, want %q", gotN, wantN)
	}
	gotMB := pBr1.MaxBooks
	wantMB := 1
	if gotMB != wantMB {
		t.Fatalf("br1.MaxBooks == %v, want %v", gotMB, wantMB)
	}
}

func TestSetValues(t *testing.T) {
	n := "Borrower1"
	gotBrN := Borrower{"Jack", 1}
	gotBrN.SetName(n)
	wantBrN := *pBr1
	if gotBrN != wantBrN {
		t.Fatalf("br.SetName(%q) == %v, want %v", n, gotBrN, wantBrN)
	}
	mb := 1
	gotBrMB := Borrower{"Borrower1", 11}
	gotBrMB.SetMaxBooks(mb)
	wantBrMB := *pBr1
	if gotBrMB != wantBrMB {
		t.Fatalf("br.SetMaxBooks(%v) == %v, want %v", mb, gotBrMB, wantBrMB)
	}
}

func TestBorrowerToString(t *testing.T) {
	gotS := pBr1.BorrowerToString()
	wantS := "Borrower1 (1 books)"
	if gotS != wantS {
		t.Fatalf("br.BorrowerToString() == %v, want %v", gotS, wantS)
	}
}
