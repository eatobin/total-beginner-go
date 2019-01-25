package library

import (
	"errors"
	"reflect"
	"testing"

	"github.com/eatobin/totalbeginnergo/book"
	"github.com/eatobin/totalbeginnergo/borrower"
)

var br1libPtr = borrower.NewBorrower("Borrower1", 1)
var br2libPtr = borrower.NewBorrower("Borrower2", 2)
var br3Ptr = borrower.NewBorrower("Borrower3", 3)

var brs1 = []borrower.Borrower{*br1libPtr, *br2libPtr}
var brs2 = []borrower.Borrower{*br1libPtr, *br2libPtr, *br3Ptr}

var bk1libPtr = &book.Book{Title: "Title1", Author: "Author1", Borrower: *br1libPtr}
var bk2Ptr = book.NewBook("Title2", "Author2")
var bk3Ptr = &book.Book{Title: "Title3", Author: "Author3", Borrower: *br3Ptr}
var bk4Ptr = &book.Book{Title: "Title4", Author: "Author4", Borrower: borrower.Borrower{Name: "Borrower3", MaxBooks: 3}}

var bks1 = []book.Book{*bk1libPtr, *bk2Ptr}
var bks2 = []book.Book{*bk1libPtr, *bk2Ptr, *bk3Ptr}
var bks3 = []book.Book{*bk1libPtr, *bk2Ptr, *bk3Ptr, *bk4Ptr}

var jsonStringBorrowers = "[\n  {\n    \"name\": \"Borrower1\",\n    \"max-books\": 1\n  },\n  {\n    \"name\": \"Borrower2\",\n    \"max-books\": 2\n  }\n]"
var jsonStringBooks = "[\n  {\n    \"title\": \"Title2\",\n    \"author\": \"Author22\",\n    \"borrower\": {\n      \"name\": \"NoName\",\n      \"max-books\": -1\n    }\n  }\n]"
var jsonStringBorrowersBadParse = "[{\"name\"\"Borrower1\",\"max-books\":1},{\"name\":\"Borrower2\",\"max-books\":2}]"
var jsonStringBorrowersBadNameField = "[{\"noName\":\"Borrower1\",\"max-books\":1},{\"name\":\"Borrower2\",\"max-books\":2}]"
var jsonStringBooksBadParse = "[{\"title\"\"Title2\",\"author\":\"Author22\",\"borrower\":{\"name\":\"NoName\",\"max-books\":-1}}, {\"title\":\"Title99\",\"author\":\"Author99\",\"borrower\":{\"name\":\"Borrower1\",\"max-books\":1}}]"
var jsonStringBooksBadTitleField = "[{\"noTitle\":\"Title2\",\"author\":\"Author22\",\"borrower\":{\"name\":\"NoName\",\"max-books\":-1}}, {\"title\":\"Title99\",\"author\":\"Author99\",\"borrower\":{\"name\":\"Borrower1\",\"max-books\":1}}]"
var jsonStringBooksBadBorrowerField = "[{\"title\":\"Title2\",\"author\":\"Author22\",\"borrower\":{\"noName\":\"NoName\",\"max-books\":-1}}, {\"title\":\"Title99\",\"author\":\"Author99\",\"borrower\":{\"name\":\"Borrower1\",\"max-books\":1}}]"
var ss = "\n--- Status Report of Test Library ---\n\nTest Library: 3 books; 3 borrowers.\n\nTitle1 by Author1; Checked out to Borrower1\nTitle2 by Author2; Available\nTitle3 by Author3; Checked out to Borrower3\n\nBorrower1 (1 books)\nBorrower2 (2 books)\nBorrower3 (3 books)\n\n--- End of Status Report ---\n"

func TestAddBorrower(t *testing.T) {
	cases := []struct {
		brs     []borrower.Borrower
		br      borrower.Borrower
		wantBrs []borrower.Borrower
	}{
		{brs1, *br3Ptr, brs2},
		{brs1, *br2libPtr, brs1},
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
		{bks1, *bk3Ptr, bks2},
		{bks1, *bk2Ptr, bks1},
	}
	for _, c := range cases {
		gotBks := AddBook(c.bks, c.bk)
		if !reflect.DeepEqual(gotBks, c.wantBks) {
			t.Errorf("AddBook(%v, %v) ==\n%v want\n%v",
				c.bks, c.bk, gotBks, c.wantBks)
		}
	}
}

func TestFindBorrower(t *testing.T) {
	var ErrNoBorrowerFound = errors.New("did not find the requested borrower")
	cases := []struct {
		n       string
		brs     []borrower.Borrower
		wantBr  borrower.Borrower
		wantErr error
	}{
		{"Borrower1", brs2, *br1libPtr, nil},
		{"Borrower11", brs2, borrower.Borrower{}, ErrNoBorrowerFound},
	}
	for _, c := range cases {
		gotBr, gotErr := FindBorrower(c.n, c.brs)
		if !reflect.DeepEqual(gotBr, c.wantBr) ||
			!reflect.DeepEqual(gotErr, c.wantErr) {
			t.Errorf("FindBorrower(%s, %v) ==\n%v\nwant\n%v\n%v\nwant\n%v",
				c.n, c.brs, gotBr, c.wantBr, gotErr, c.wantErr)
		}
	}
}

func TestFindBook(t *testing.T) {
	var ErrNoBookFound = errors.New("did not find the requested book")
	cases := []struct {
		t       string
		bks     []book.Book
		wantI   int
		wantBk  book.Book
		wantErr error
	}{
		{"Title1", bks2, 0, *bk1libPtr, nil},
		{"Title11", bks2, 0, book.Book{}, ErrNoBookFound},
	}
	for _, c := range cases {
		gotI, gotBk, gotErr := FindBook(c.t, c.bks)
		if gotI != c.wantI ||
			!reflect.DeepEqual(gotBk, c.wantBk) ||
			!reflect.DeepEqual(gotErr, c.wantErr) {
			t.Errorf("FindBook(%s, %v) ==\n%v\nwant\n%v\n%v\nwant\n%v\n%v\nwant\n%v",
				c.t, c.bks, gotI, c.wantI, gotBk, c.wantBk, gotErr, c.wantErr)
		}
	}
}

func TestGetBooksForBorrower(t *testing.T) {
	cases := []struct {
		br   borrower.Borrower
		bks  []book.Book
		want []book.Book
	}{
		{*br2libPtr, bks1, []book.Book{}},
		{*br1libPtr, bks1, []book.Book{*bk1libPtr}},
		{*br3Ptr, bks3, []book.Book{*bk3Ptr, *bk4Ptr}},
	}
	for _, c := range cases {
		got := GetBooksForBorrower(c.br, c.bks)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("GetBoksForborrower(%v, %v) ==\n%v want \n%v",
				c.br, c.bks, got, c.want)
		}
	}
}

func TestCheckOut(t *testing.T) {
	var bks2 = []book.Book{*bk1libPtr, {Title: "Title2", Author: "Author2", Borrower: borrower.Borrower{Name: "Borrower2", MaxBooks: 2}}}
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
		{"Borrower2", "Title2", brs1, bks1, bks2},
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
	var bks1 = []book.Book{*bk1libPtr, *bk2Ptr}
	var bks2 = []book.Book{{Title: "Title1", Author: "Author1", Borrower: borrower.Borrower{}}, *bk2Ptr}
	cases := []struct {
		t    string
		bks  []book.Book
		want []book.Book
	}{
		{"Title1", bks1, bks2},
		{"Title2", bks1, bks1},
		{"NoTitle", bks1, bks1},
	}
	for _, c := range cases {
		got := CheckIn(c.t, c.bks)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("CheckIn(%s, %v) ==\n%v want \n%v",
				c.t, c.bks, got, c.want)
		}
	}
}

func TestJSONStringToBorrowersPass(t *testing.T) {
	js := jsonStringBorrowers
	wantBrs := brs1
	wantE := error(nil)

	got, err := JSONStringToBorrowers(js)
	if !reflect.DeepEqual(got, wantBrs) || err != wantE {
		t.Errorf("JSONStringToBorrowers\n(%s)\n==\n%v and %v\nwant\n%v and %v",
			js, got, err, wantBrs, wantE)
	}
}

func TestJSONStringToBorrowersFail(t *testing.T) {
	cases := []struct {
		js      string
		wantBrs []borrower.Borrower
		wantE   error
	}{
		{jsonStringBorrowersBadParse, []borrower.Borrower{}, errors.New("invalid character '\"' after object key")},
		{jsonStringBorrowersBadNameField, []borrower.Borrower{}, errors.New("missing Borrower field value - borrowers list is empty")},
	}
	for _, c := range cases {
		got, err := JSONStringToBorrowers(c.js)
		if err.Error() != c.wantE.Error() {
			t.Errorf("JSONStringToBorrowers\n(%s)\n==\n%v and %v\nwant\n%v and %v",
				c.js, got, err, c.wantBrs, c.wantE)
		}
	}
}

func TestJSONStringToBooksPass(t *testing.T) {
	js := jsonStringBooks
	wantBks := []book.Book{{Title: "Title2", Author: "Author22", Borrower: borrower.Borrower{Name: "NoName", MaxBooks: -1}}}
	wantE := error(nil)

	got, err := JSONStringToBooks(js)

	if !reflect.DeepEqual(got, wantBks) || err != wantE {
		t.Errorf("JSONStringToBooks\n(%s)\n==\n%v and %v\nwant\n%v and %v",
			js, got, err, wantBks, wantE)
	}
}

func TestJSONStringToBooksFail(t *testing.T) {
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
		got, err := JSONStringToBooks(c.js)
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
	bks := bks2
	brs := brs2
	got := StatusToString(bks, brs)
	want := ss
	if got != want {
		t.Errorf("StatusToString(%v, %v) ==\n(%q) want \n(%q)",
			bks, brs, got, want)
	}
}
