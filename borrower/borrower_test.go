package borrower

import "testing"

var wantS = "Borrower1 (1 books)"

func TestSetBorrowerValues(t *testing.T) {
	n := "Borrower1"
	badBrN := NewBorrower("Jack", 1)
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
