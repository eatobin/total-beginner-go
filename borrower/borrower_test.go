package borrower

import "testing"

var br1 = MakeBorrower("Borrower1", 1)

func TestMakeBorrower(t *testing.T) {
	gotN := br1.Name
	wantN := "Borrower1"
	if gotN != wantN {
		t.Fatalf("br1lib.Name == %q, want %q", gotN, wantN)
	}
	gotMB := br1.MaxBooks
	wantMB := 1
	if gotMB != wantMB {
		t.Fatalf("br1lib.MaxBooks == %v, want %v", gotMB, wantMB)
	}
}

func TestSetValues(t *testing.T) {
	n := "Borrower1"
	gotBrN := Borrower{"Jack", 1}
	gotBrN.SetName(n)
	wantBrN := br1
	if gotBrN != wantBrN {
		t.Fatalf("br.SetName(%q) == %v, want %v", n, gotBrN, wantBrN)
	}
	mb := 1
	gotBrMB := Borrower{"Borrower1", 11}
	gotBrMB.SetMaxBooks(mb)
	wantBrMB := br1
	if gotBrMB != wantBrMB {
		t.Fatalf("br.SetMaxBooks(%v) == %v, want %v", mb, gotBrMB, wantBrMB)
	}
}

func TestBorrowerToString(t *testing.T) {
	gotS := br1.BorrowerToString()
	wantS := "Borrower1 (1 books)"
	if gotS != wantS {
		t.Fatalf("br.BorrowerToString() == %v, want %v", gotS, wantS)
	}
}
