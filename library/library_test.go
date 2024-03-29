package library

import (
	"errors"
	"reflect"
	"testing"
)

var br1lib = Borrower{Name: "Borrower1", MaxBooks: 1}
var br2lib = Borrower{Name: "Borrower2", MaxBooks: 2}
var br3lib = Borrower{Name: "Borrower3", MaxBooks: 3}

var brs1 = Borrowers{br1lib, br2lib}
var brs2 = Borrowers{br1lib, br2lib, br3lib}

var bk1lib = Book{Title: "Title1", Author: "Author1", Borrower: &br1lib}
var bk2lib = Book{Title: "Title2", Author: "Author2"}
var bk3lib = Book{Title: "Title3", Author: "Author3", Borrower: &br3lib}
var bk4lib = Book{Title: "Title4", Author: "Author4", Borrower: &br3lib}

var bks1 = Books{bk1lib, bk2lib}
var bks2 = Books{bk1lib, bk2lib, bk3lib}
var bks3 = Books{bk1lib, bk2lib, bk3lib, bk4lib}

var jsonStringBorrowers = "[{\"name\":\"Borrower1\",\"maxBooks\":1},{\"name\":\"Borrower2\",\"maxBooks\":2}]"
var jsonStringBorrowersBadParse = `[{"name""Borrower1","maxBooks":1},{"name":"Borrower2","maxBooks":2}]`
var jsonStringBorrowersBadNameField = `[{"noName":"Borrower1","maxBooks":1},{"name":"Borrower2","maxBooks":2}]`
var jsonStringBorrowersBadMaxBooksField = `[{"name":"Borrower1","noMaxBooks":1},{"name":"Borrower2","maxBooks":2}]`

var jsonStringBooks = `[{"title":"Title1","author":"Author1","borrower":{"name":"Borrower1","maxBooks":1}},{"title":"Title2","author":"Author2"}]`
var jsonStringBooksBadParse = `[{"title""Title2","author":"Author22","borrower":{"name":"NoName","maxBooks":-1}},{"title":"Title99","author":"Author99","borrower":{"name":"Borrower1","maxBooks":1}}]`
var jsonStringBooksBadTitleField = `[{"noTitle":"Title2","author":"Author22","borrower":{"name":"NoName","maxBooks":-1}},{"title":"Title99","author":"Author99","borrower":{"name":"Borrower1","maxBooks":1}}]`
var jsonStringBooksBadBorrowerField = `[{"title":"Title1","author":"Author1","borrower":{"noName":"Borrower1","maxBooks":1}},{"title":"Title2","author":"Author2","borrower":{"name":"Borrower2","maxBooks":2}}]`

var ss = "\n--- Status Report of Test Library ---\n\nTest Library: 3 books; 3 borrowers.\n\nTitle1 by Author1; Checked out to Borrower1\nTitle2 by Author2; Available\nTitle3 by Author3; Checked out to Borrower3\n\nBorrower1 [1 books]\nBorrower2 [2 books]\nBorrower3 [3 books]\n\n--- End of Status Report ---\n"

func TestAddBorrower(t *testing.T) {
	cases := []struct {
		brs     Borrowers
		br      Borrower
		wantBrs Borrowers
	}{
		{brs1, br3lib, brs2},
		{brs1, br2lib, brs1},
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
		bks     Books
		bk      Book
		wantBks Books
	}{
		{bks1, bk3lib, bks2},
		{bks1, bk2lib, bks1},
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
		brs     Borrowers
		wantBr  Borrower
		wantErr error
	}{
		{"Borrower1", brs2, br1lib, nil},
		{"Borrower11", brs2, ZeroBorrower, ErrNoBorrowerFound},
	}
	for _, c := range cases {
		gotBr, gotErr := findBorrower(c.n, c.brs)
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
		bks     Books
		wantBk  Book
		wantErr error
	}{
		{"Title1", bks2, bk1lib, nil},
		{"Title11", bks2, ZeroBook, ErrNoBookFound},
	}
	for _, c := range cases {
		gotBk, gotErr := findBook(c.t, c.bks)
		if gotBk != c.wantBk ||
			!reflect.DeepEqual(gotBk, c.wantBk) ||
			!reflect.DeepEqual(gotErr, c.wantErr) {
			t.Errorf("findBook(%s, %v) ==\n%v\nwant\n%v\n%v\nwant\n%v",
				c.t, c.bks, gotBk, c.wantBk, gotErr, c.wantErr)
		}
	}
}

func Test_getBooksForBorrower(t *testing.T) {
	cases := []struct {
		br   Borrower
		bks  Books
		want Books
	}{
		{br2lib, bks1, []Book{}},
		{br1lib, bks1, []Book{bk1lib}},
		{br3lib, bks3, []Book{bk3lib, bk4lib}},
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
	var testbks = Books{bk1lib, {Title: "Title2", Author: "Author2", Borrower: &Borrower{Name: "Borrower2", MaxBooks: 2}}}
	cases := []struct {
		n    string
		t    string
		brs  Borrowers
		bks  Books
		want Books
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
	var testbks1 = Books{bk1lib, bk2lib}
	var testbks2 = Books{bk2lib, {Title: "Title1", Author: "Author1", Borrower: nil}}
	cases := []struct {
		t    string
		bks  Books
		want Books
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
	js := jsonStringBorrowers
	wantBrs := brs1
	wantError := error(nil)

	got, err := JsonStringToBorrowers(js)
	if !reflect.DeepEqual(got, wantBrs) || err != wantError {
		t.Errorf("JsonStringToBorrowers\n(%s)\n==\n%v and %v\nwant\n%v and %v",
			js, got, err, wantBrs, wantError)
	}
}

func Test_jsonStringToBorrowersFail(t *testing.T) {
	cases := []struct {
		js        string
		wantBrs   Borrowers
		wantError error
	}{
		{jsonStringBorrowersBadParse, ZeroBorrowers, errors.New("invalid character '\"' after object key")},
		{jsonStringBorrowersBadNameField, ZeroBorrowers, errors.New("missing Borrower field value - borrowers list is empty")},
		{jsonStringBorrowersBadMaxBooksField, ZeroBorrowers, errors.New("missing Borrower field value - borrowers list is empty")},
	}
	for _, c := range cases {
		got, err := JsonStringToBorrowers(c.js)
		if err != nil {
			if err.Error() != c.wantError.Error() {
				t.Errorf("JsonStringToBorrowers\n(%s)\n==\n%v and %v\nwant\n%v and %v",
					c.js, got, err, c.wantBrs, c.wantError)
			}
		}
	}
}

func Test_jsonStringToBooksPass(t *testing.T) {
	js := jsonStringBooks
	wantBks := Books{{Title: "Title1", Author: "Author1", Borrower: &Borrower{Name: "Borrower1", MaxBooks: 1}}, {Title: "Title2", Author: "Author2", Borrower: nil}}
	wantError := error(nil)

	got, err := JsonStringToBooks(js)

	if !reflect.DeepEqual(got, wantBks) || err != wantError {
		t.Errorf("JSONStringToBooks\n(%s)\n==\n%v and %v\nwant\n%v and %v",
			js, got, err, wantBks, wantError)
	}
}

func Test_jsonStringToBooksFail(t *testing.T) {
	cases := []struct {
		js        string
		wantBks   Books
		wantError error
	}{
		{jsonStringBooksBadParse, ZeroBooks, errors.New("invalid character '\"' after object key")},
		{jsonStringBooksBadTitleField, ZeroBooks, errors.New("missing Book field value - book list is empty")},
		{jsonStringBooksBadBorrowerField, ZeroBooks, errors.New("missing Borrower field value - book list is empty")},
	}
	for _, c := range cases {
		got, err := JsonStringToBooks(c.js)
		if err != nil {
			if err.Error() != c.wantError.Error() {
				t.Errorf("JSONStringToBooks\n(%s)\n==\n%v and %v\nwant\n%v and %v",
					c.js, got, err, c.wantBks, c.wantError)
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
	want := jsonStringBooks
	if got != want || err != nil {
		t.Errorf("BooksToJSONSting(%v) ==\n(%q) want \n(%q)",
			bks1, got, want)
	}
}

func TestStatusToString(t *testing.T) {
	br1libL := Borrower{Name: "Borrower1", MaxBooks: 1}
	br2libL := Borrower{Name: "Borrower2", MaxBooks: 2}
	br3libL := Borrower{Name: "Borrower3", MaxBooks: 3}
	brs2L := Borrowers{br1libL, br2libL, br3libL}

	bk1libL := Book{Title: "Title1", Author: "Author1", Borrower: &br1lib}
	bk2libL := Book{Title: "Title2", Author: "Author2"}
	bk3libL := Book{Title: "Title3", Author: "Author3", Borrower: &br3lib}
	bks2L := Books{bk1libL, bk2libL, bk3libL}

	bks := bks2L
	brs := brs2L
	got := StatusToString(bks, brs)
	want := ss
	if got != want {
		t.Errorf("StatusToString(%v, %v) ==\n(%q) want \n(%q)",
			bks, brs, got, want)
	}
}
