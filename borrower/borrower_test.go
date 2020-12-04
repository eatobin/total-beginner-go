package borrower

import "testing"

var br1 = NewBorrower("Borrower1", 1)
var wantS = "Borrower1 (1 books)"

func TestBrToString(t *testing.T) {
	gotS := BrToString(br1)
	if gotS != wantS {
		t.Fatalf("BrToString(%v) == %v, want %v", br1, gotS, wantS)
	}
}

func TestSetBorrowerValues(t *testing.T) {
	n := "Borrower1"
	badBrN := Borrower{"Jack", 1}
	gotBrN := BrToString(SetName(badBrN, n))
	if gotBrN != wantS {
		t.Fatalf("SetName(%v, %v) == %v, want %v", badBrN, n, gotBrN, wantS)
	}
	mb := 1
	badBrMB := Borrower{"Borrower1", 11}
	gotBrMB := BrToString(SetMaxBooks(badBrMB, mb))
	if gotBrMB != wantS {
		t.Fatalf("SetMaxBooks(%v, %v) == %v, want %v", badBrMB, mb, gotBrMB, wantS)
	}
}
