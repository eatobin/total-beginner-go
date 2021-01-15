package library

import (
	"eatobin.com/totalbeginnergo/book"
	"eatobin.com/totalbeginnergo/borrower"
	"errors"
	"reflect"
	"testing"
)

var br1libPtr = borrower.NewBorrower("Borrower1", 1)
var br2libPtr = borrower.NewBorrower("Borrower2", 2)
var br3libPtr = borrower.NewBorrower("Borrower3", 3)

var brs1 = []*borrower.Borrower{br1libPtr, br2libPtr}
var brs2 = []*borrower.Borrower{br1libPtr, br2libPtr, br3libPtr}

var bk1libPtr = &book.Book{Title: "Title1", Author: "Author1", Borrower: br1libPtr}
var bk2libPtr = book.NewBook("Title2", "Author2")
var bk3libPtr = &book.Book{Title: "Title3", Author: "Author3", Borrower: br3libPtr}

var bk4libPtr = &book.Book{Title: "Title4", Author: "Author4", Borrower: br3libPtr}

var bks1 = []*book.Book{bk1libPtr, bk2libPtr}
var bks2 = []*book.Book{bk1libPtr, bk2libPtr, bk3libPtr}

var bks3 = []*book.Book{bk1libPtr, bk2libPtr, bk3libPtr, bk4libPtr}

var jsonStringBorrowers = "[{\"name\":\"Borrower1\",\"maxBooks\":1},{\"name\":\"Borrower2\",\"maxBooks\":2}]"
var jsonStringBorrowersBadParse = `[{"name""Borrower1","maxBooks":1},{"name":"Borrower2","maxBooks":2}]`
var jsonStringBorrowersBadNameField = `[{"noName":"Borrower1","maxBooks":1},{"name":"Borrower2","maxBooks":2}]`
var jsonStringBorrowersBadMaxBooksField = `[{"name":"Borrower1","noMaxBooks":1},{"name":"Borrower2","maxBooks":2}]`

var jsonStringBooks = `[{"title":"Title1","author":"Author1","borrower":{"name":"Borrower1","maxBooks":1}},{"title":"Title2","author":"Author2","borrower":null}]`
var jsonStringBooksBadParse = `[{"title""Title2","author":"Author22","borrower":{"name":"NoName","maxBooks":-1}},{"title":"Title99","author":"Author99","borrower":{"name":"Borrower1","maxBooks":1}}]`
var jsonStringBooksBadTitleField = `[{"noTitle":"Title2","author":"Author22","borrower":{"name":"NoName","maxBooks":-1}},{"title":"Title99","author":"Author99","borrower":{"name":"Borrower1","maxBooks":1}}]`
var jsonStringBooksBadBorrowerField = `[{"title":"Title1","author":"Author1","borrower":{"noName":"Borrower1","maxBooks":1}},{"title":"Title2","author":"Author2","borrower":{"name":"Borrower2","maxBooks":2}}]`
var jsonStringBooks2 = "[{\"title\":\"Title1\",\"author\":\"Author1\",\"borrower\":null},{\"title\":\"Title2\",\"author\":\"Author2\",\"borrower\":null}]"
var ss = "\n--- Status Report of Test Library ---\n\nTest Library: 3 books; 3 borrowers.\n\nTitle1 by Author1; Checked out to Borrower1\nTitle2 by Author2; Available\nTitle3 by Author3; Checked out to Borrower3\n\nBorrower1 (1 books)\nBorrower2 (2 books)\nBorrower3 (3 books)\n\n--- End of Status Report ---\n"

func TestAddBorrower(t *testing.T) {
	cases := []struct {
		brs     []*borrower.Borrower
		br      *borrower.Borrower
		wantBrs []*borrower.Borrower
	}{
		{brs1, br3libPtr, brs2},
		{brs1, br2libPtr, brs1},
	}
	for _, c := range cases {
		gotBrs := AddBorrower(c.brs, c.br)
		if !reflect.DeepEqual(gotBrs, c.wantBrs) {
			t.Errorf("AddBorrower(%v, %v) ==\n%v want\n%v",
				c.brs, c.br, gotBrs, c.wantBrs)
		}
	}
}

func TestAddBook(t *testing.T) {
	cases := []struct {
		bks     []*book.Book
		bk      *book.Book
		wantBks []*book.Book
	}{
		{bks1, bk3libPtr, bks2},
		{bks1, bk2libPtr, bks1},
	}
	for _, c := range cases {
		gotBks := AddBook(c.bks, c.bk)
		if !reflect.DeepEqual(gotBks, c.wantBks) {
			t.Errorf("AddBook(%v, %v) ==\n%v want\n%v",
				c.bks, c.bk, gotBks, c.wantBks)
		}
	}
}

func Test_findBorrower(t *testing.T) {
	var ErrNoBorrowerFound = errors.New("did not find the requested borrower")
	cases := []struct {
		n       string
		brs     []*borrower.Borrower
		wantBr  *borrower.Borrower
		wantErr error
	}{
		{"Borrower1", brs2, br1libPtr, nil},
		{"Borrower11", brs2, &borrower.Borrower{}, ErrNoBorrowerFound},
	}
	for _, c := range cases {
		gotErr, gotBr := findBorrower(c.n, c.brs)
		if !reflect.DeepEqual(gotBr, c.wantBr) ||
			!reflect.DeepEqual(gotErr, c.wantErr) {
			t.Errorf("findBorrower(%s, %v) ==\n%v\nwant\n%v\n%v\nwant\n%v",
				c.n, c.brs, gotBr, c.wantBr, gotErr, c.wantErr)
		}
	}
}

func Test_findBook(t *testing.T) {
	var ErrNoBookFound = errors.New("did not find the requested book")
	cases := []struct {
		t       string
		bks     []*book.Book
		wantI   int
		wantBk  *book.Book
		wantErr error
	}{
		{"Title1", bks2, 0, bk1libPtr, nil},
		{"Title11", bks2, 0, &book.Book{}, ErrNoBookFound},
	}
	for _, c := range cases {
		gotI, gotErr, gotBk := findBook(c.t, c.bks)
		if gotI != c.wantI ||
			!reflect.DeepEqual(gotBk, c.wantBk) ||
			!reflect.DeepEqual(gotErr, c.wantErr) {
			t.Errorf("findBook(%s, %v) ==\n%v\nwant\n%v\n%v\nwant\n%v\n%v\nwant\n%v",
				c.t, c.bks, gotI, c.wantI, gotBk, c.wantBk, gotErr, c.wantErr)
		}
	}
}

func Test_getBooksForBorrower(t *testing.T) {
	cases := []struct {
		br   *borrower.Borrower
		bks  []*book.Book
		want []*book.Book
	}{
		{br2libPtr, bks1, []*book.Book{}},
		{br1libPtr, bks1, []*book.Book{bk1libPtr}},
		{br3libPtr, bks3, []*book.Book{bk3libPtr, bk4libPtr}},
	}
	for _, c := range cases {
		got := getBooksForBorrower(c.br, c.bks)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("GetBooksForborrower(%v, %v) ==\n%v want \n%v",
				c.br, c.bks, got, c.want)
		}
	}
}

func TestCheckOut(t *testing.T) {
	var testbks = []*book.Book{bk1libPtr, {Title: "Title2", Author: "Author2", Borrower: &borrower.Borrower{Name: "Borrower2", MaxBooks: 2}}}
	cases := []struct {
		n    string
		t    string
		brs  []*borrower.Borrower
		bks  []*book.Book
		want []*book.Book
	}{
		{"Borrower2", "Title1", brs1, bks1, bks1},
		{"Borrower2", "NoTitle", brs1, bks1, bks1},
		{"NoName", "Title1", brs1, bks1, bks1},
		{"Borrower1", "Title2", brs1, bks1, bks1},
		{"Borrower2", "Title2", brs1, bks1, testbks},
	}
	for _, c := range cases {
		got := CheckOut(c.n, c.t, c.brs, c.bks)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("CheckOut(%s, %s, %v, %v) ==\n%v want \n%v",
				c.n, c.t, c.brs, c.bks, got, c.want)
		}
	}
}

func TestCheckIn(t *testing.T) {
	var testbks1 = []*book.Book{bk1libPtr, bk2libPtr}
	var testbks2 = []*book.Book{{Title: "Title1", Author: "Author1", Borrower: nil}, bk2libPtr}
	cases := []struct {
		t    string
		bks  []*book.Book
		want []*book.Book
	}{
		{"Title1", testbks1, testbks2},
		{"Title2", testbks1, testbks1},
		{"NoTitle", testbks1, testbks1},
	}
	for _, c := range cases {
		got := CheckIn(c.t, c.bks)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("CheckIn(%s, %v) ==\n%v want \n%v",
				c.t, c.bks, got, c.want)
		}
	}
}

func Test_jsonStringToBorrowersPass(t *testing.T) {
	wantBrs := brs1
	wantE := error(nil)

	got, err := JsonStringToBorrowers(jsonStringBorrowers)
	if !reflect.DeepEqual(got, wantBrs) || err != wantE {
		t.Errorf("JsonStringToBorrowers\n(%s)\n==\n%v and %v\nwant\n%v and %v",
			jsonStringBorrowers, got, err, wantBrs, wantE)
	}
}

func Test_jsonStringToBorrowersFail(t *testing.T) {
	cases := []struct {
		js      string
		wantBrs []*borrower.Borrower
		wantE   error
	}{
		{jsonStringBorrowersBadParse, []*borrower.Borrower{}, errors.New("invalid character '\"' after object key")},
		{jsonStringBorrowersBadNameField, []*borrower.Borrower{}, errors.New("missing Borrower field value - borrowers list is empty")},
		{jsonStringBorrowersBadMaxBooksField, []*borrower.Borrower{}, errors.New("missing Borrower field value - borrowers list is empty")},
	}
	for _, c := range cases {
		got, err := JsonStringToBorrowers(c.js)
		if err != nil {
			if err.Error() != c.wantE.Error() {
				t.Errorf("JsonStringToBorrowers\n(%s)\n==\n%v and %v\nwant\n%v and %v",
					c.js, got, err, c.wantBrs, c.wantE)
			}
		}
	}
}

func Test_jsonStringToBooksPass(t *testing.T) {
	wantBks := []*book.Book{{Title: "Title1", Author: "Author1", Borrower: &borrower.Borrower{Name: "Borrower1", MaxBooks: 1}}, {Title: "Title2", Author: "Author2", Borrower: nil}}
	wantE := error(nil)

	got, err := JsonStringToBooks(jsonStringBooks)

	if !reflect.DeepEqual(got, wantBks) || err != wantE {
		t.Errorf("JSONStringToBooks\n(%s)\n==\n%v and %v\nwant\n%v and %v",
			jsonStringBooks, got, err, wantBks, wantE)
	}
}

func Test_jsonStringToBooksFail(t *testing.T) {
	cases := []struct {
		js      string
		wantBks []*book.Book
		wantE   error
	}{
		{jsonStringBooksBadParse, []*book.Book{}, errors.New("invalid character '\"' after object key")},
		{jsonStringBooksBadTitleField, []*book.Book{}, errors.New("missing Book field value - book list is empty")},
		{jsonStringBooksBadBorrowerField, []*book.Book{}, errors.New("missing Borrower field value - book list is empty")},
	}
	for _, c := range cases {
		got, err := JsonStringToBooks(c.js)
		if err != nil {
			if err.Error() != c.wantE.Error() {
				t.Errorf("JSONStringToBooks\n(%s)\n==\n%v and %v\nwant\n%v and %v",
					c.js, got, err, c.wantBks, c.wantE)
			}
		}
	}
}

func TestBorrowersToJSONString(t *testing.T) {
	got, err := BorrowersToJSONSting(brs1)
	want := jsonStringBorrowers
	if got != want || err != nil {
		t.Errorf("BorrwersToJSONSting(%v) ==\n(%q) want \n(%q)",
			brs1, got, want)
	}
}

func TestBooksToJSONString(t *testing.T) {
	got, err := BooksToJSONSting(bks1)
	want := jsonStringBooks2
	if got != want || err != nil {
		t.Errorf("BooksToJSONSting(%v) ==\n(%q) want \n(%q)",
			bks1, got, want)
	}
}

func TestStatusToString(t *testing.T) {
	br1libPtrL := borrower.NewBorrower("Borrower1", 1)
	br2libPtrL := borrower.NewBorrower("Borrower2", 2)
	br3libPtrL := borrower.NewBorrower("Borrower3", 3)
	brs2L := []*borrower.Borrower{br1libPtrL, br2libPtrL, br3libPtrL}

	bk1libPtrL := &book.Book{Title: "Title1", Author: "Author1", Borrower: br1libPtr}
	bk2libPtrL := book.NewBook("Title2", "Author2")
	bk3libPtrL := &book.Book{Title: "Title3", Author: "Author3", Borrower: br3libPtr}
	bks2L := []*book.Book{bk1libPtrL, bk2libPtrL, bk3libPtrL}

	bks := bks2L
	brs := brs2L
	got := StatusToString(bks, brs)
	want := ss
	if got != want {
		t.Errorf("StatusToString(%v, %v) ==\n(%q) want \n(%q)",
			bks, brs, got, want)
	}
}
