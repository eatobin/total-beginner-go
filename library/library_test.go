package library

import (
	"eatobin.com/totalbeginnergo/book"
	"eatobin.com/totalbeginnergo/borrower"
	"errors"
	"reflect"
	"testing"
)

var br1lib = borrower.NewBorrower("Borrower1", 1)
var br2lib = borrower.NewBorrower("Borrower2", 2)
var br3lib = borrower.NewBorrower("Borrower3", 3)

var brs1 = []borrower.Borrower{br1lib, br2lib}
var brs2 = []borrower.Borrower{br1lib, br2lib, br3lib}

var bk1lib = book.Book{Title: "Title1", Author: "Author1", Borrower: br1lib}
var bk2lib = book.NewBook("Title2", "Author2")
var bk3lib = book.Book{Title: "Title3", Author: "Author3", Borrower: br3lib}

var bk4lib = book.Book{Title: "Title4", Author: "Author4", Borrower: br3lib}

var bks1 = []book.Book{bk1lib, bk2lib}
var bks2 = []book.Book{bk1lib, bk2lib, bk3lib}

var bks3 = []book.Book{bk1lib, bk2lib, bk3lib, bk4lib}

var jsonStringBorrowers = "[\n  {\n    \"name\": \"Borrower1\",\n    \"max-books\": 1\n  },\n  {\n    \"name\": \"Borrower2\",\n    \"max-books\": 2\n  }\n]"
var jsonStringBorrowersBadParse = `[{"name""Borrower1","max-books":1},{"name":"Borrower2","max-books":2}]`
var jsonStringBorrowersBadNameField = `[{"noName":"Borrower1","max-books":1},{"name":"Borrower2","max-books":2}]`
var jsonStringBorrowersBadMaxBooksField = `[{"name":"Borrower1","noMax-books":1},{"name":"Borrower2","max-books":2}]`

var jsonStringBooks = "[\n  {\n    \"title\": \"Title2\",\n    \"author\": \"Author22\",\n    \"borrower\": {\n      \"name\": \"NoName\",\n      \"max-books\": -1\n    }\n  }\n]"
var jsonStringBooksBadParse = `[{"title""Title2","author":"Author22","borrower":{"name":"NoName","max-books":-1}},{"title":"Title99","author":"Author99","borrower":{"name":"Borrower1","max-books":1}}]`
var jsonStringBooksBadTitleField = `[{"noTitle":"Title2","author":"Author22","borrower":{"name":"NoName","max-books":-1}},{"title":"Title99","author":"Author99","borrower":{"name":"Borrower1","max-books":1}}]`
var jsonStringBooksBadBorrowerField = `[{"title":"Title2","author":"Author22","borrower":{"noName":"NoName","max-books":-1}},{"title":"Title99","author":"Author99","borrower":{"name":"Borrower1","max-books":1}}]`

var ss = "\n--- Status Report of Test Library ---\n\nTest Library: 3 books; 3 borrowers.\n\nTitle1 by Author1; Checked out to Borrower1\nTitle2 by Author2; Available\nTitle3 by Author3; Checked out to Borrower3\n\nBorrower1 (1 books)\nBorrower2 (2 books)\nBorrower3 (3 books)\n\n--- End of Status Report ---\n"

func TestAddBorrower(t *testing.T) {
	cases := []struct {
		brs     []borrower.Borrower
		br      borrower.Borrower
		wantBrs []borrower.Borrower
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
		bks     []book.Book
		bk      book.Book
		wantBks []book.Book
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
		brs     []borrower.Borrower
		wantBr  borrower.Borrower
		wantErr error
	}{
		{"Borrower1", brs2, br1lib, nil},
		{"Borrower11", brs2, borrower.Borrower{}, ErrNoBorrowerFound},
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
		bks     []book.Book
		wantBk  book.Book
		wantErr error
	}{
		{"Title1", bks2, bk1lib, nil},
		{"Title11", bks2, book.Book{}, ErrNoBookFound},
	}
	for _, c := range cases {
		gotErr, gotBk := findBook(c.t, c.bks)
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
		br   borrower.Borrower
		bks  []book.Book
		want []book.Book
	}{
		{br2lib, bks1, []book.Book{}},
		{br1lib, bks1, []book.Book{bk1lib}},
		{br3lib, bks3, []book.Book{bk3lib, bk4lib}},
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
	var testbks = []book.Book{bk1lib, {Title: "Title2", Author: "Author2", Borrower: borrower.Borrower{Name: "Borrower2", MaxBooks: 2}}}
	cases := []struct {
		n    string
		t    string
		brs  []borrower.Borrower
		bks  []book.Book
		want []book.Book
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
	var testbks1 = []book.Book{bk1lib, bk2lib}
	var testbks2 = []book.Book{bk2lib, {Title: "Title1", Author: "Author1", Borrower: borrower.ZeroBorrower}}
	cases := []struct {
		t    string
		bks  []book.Book
		want []book.Book
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
	wantE := error(nil)

	err, got := JsonStringToBorrowers(js)
	if !reflect.DeepEqual(got, wantBrs) || err != wantE {
		t.Errorf("JsonStringToBorrowers\n(%s)\n==\n%v and %v\nwant\n%v and %v",
			js, got, err, wantBrs, wantE)
	}
}

func Test_jsonStringToBorrowersFail(t *testing.T) {
	cases := []struct {
		js      string
		wantBrs []borrower.Borrower
		wantE   error
	}{
		{jsonStringBorrowersBadParse, []borrower.Borrower{}, errors.New("invalid character '\"' after object key")},
		{jsonStringBorrowersBadNameField, []borrower.Borrower{}, errors.New("missing Borrower field value - borrowers list is empty")},
		{jsonStringBorrowersBadMaxBooksField, []borrower.Borrower{}, errors.New("missing Borrower field value - borrowers list is empty")},
	}
	for _, c := range cases {
		err, got := JsonStringToBorrowers(c.js)
		if err.Error() != c.wantE.Error() {
			t.Errorf("JsonStringToBorrowers\n(%s)\n==\n%v and %v\nwant\n%v and %v",
				c.js, got, err, c.wantBrs, c.wantE)
		}
	}
}

func Test_jsonStringToBooksPass(t *testing.T) {
	js := jsonStringBooks
	wantBks := []book.Book{{Title: "Title2", Author: "Author22", Borrower: borrower.Borrower{Name: "NoName", MaxBooks: -1}}}
	wantE := error(nil)

	err, got := JsonStringToBooks(js)

	if !reflect.DeepEqual(got, wantBks) || err != wantE {
		t.Errorf("JSONStringToBooks\n(%s)\n==\n%v and %v\nwant\n%v and %v",
			js, got, err, wantBks, wantE)
	}
}

func Test_jsonStringToBooksFail(t *testing.T) {
	cases := []struct {
		js      string
		wantBks []book.Book
		wantE   error
	}{
		{jsonStringBooksBadParse, []book.Book{}, errors.New("invalid character '\"' after object key")},
		{jsonStringBooksBadTitleField, []book.Book{}, errors.New("missing Book field value - book list is empty")},
		{jsonStringBooksBadBorrowerField, []book.Book{}, errors.New("missing Book field value - book list is empty")},
	}
	for _, c := range cases {
		err, got := JsonStringToBooks(c.js)
		if err.Error() != c.wantE.Error() {
			t.Errorf("JSONStringToBooks\n(%s)\n==\n%v and %v\nwant\n%v and %v",
				c.js, got, err, c.wantBks, c.wantE)
		}
	}
}

func TestBorrowersToJSONString(t *testing.T) {
	brs := brs1
	got := BorrowersToJSONSting(brs)
	want := jsonStringBorrowers
	if got != want {
		t.Errorf("BorrwersToJSONSting(%v) ==\n(%q) want \n(%q)",
			brs, got, want)
	}
}

func TestBooksToJSONString(t *testing.T) {
	bks := []book.Book{{Title: "Title2", Author: "Author22", Borrower: borrower.Borrower{Name: "NoName", MaxBooks: -1}}}
	got := BooksToJSONSting(bks)
	want := jsonStringBooks
	if got != want {
		t.Errorf("BooksToJSONSting(%v) ==\n(%q) want \n(%q)",
			bks, got, want)
	}
}

func TestStatusToString(t *testing.T) {
	br1libL := borrower.NewBorrower("Borrower1", 1)
	br2libL := borrower.NewBorrower("Borrower2", 2)
	br3libL := borrower.NewBorrower("Borrower3", 3)
	brs2L := []borrower.Borrower{br1libL, br2libL, br3libL}

	bk1libL := book.Book{Title: "Title1", Author: "Author1", Borrower: br1lib}
	bk2libL := book.NewBook("Title2", "Author2")
	bk3libL := book.Book{Title: "Title3", Author: "Author3", Borrower: br3lib}
	bks2L := []book.Book{bk1libL, bk2libL, bk3libL}

	bks := bks2L
	brs := brs2L
	got := StatusToString(bks, brs)
	want := ss
	if got != want {
		t.Errorf("StatusToString(%v, %v) ==\n(%q) want \n(%q)",
			bks, brs, got, want)
	}
}
