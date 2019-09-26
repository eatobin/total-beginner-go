package main

import (
	"errors"
	//"errors"
	"reflect"
	"testing"
)

var br1libPtr = NewBorrower("Borrower1", 1)
var br2libPtr = NewBorrower("Borrower2", 2)
var br3libPtr = NewBorrower("Borrower3", 3)

var brs1 = []*Borrower{br1libPtr, br2libPtr}
var brs2 = []*Borrower{br1libPtr, br2libPtr, br3libPtr}

var bk1libPtr = &Book{Title: "Title1", Author: "Author1", Borrower: br1libPtr}
var bk2libPtr = NewBook("Title2", "Author2")
var bk3libPtr = &Book{Title: "Title3", Author: "Author3", Borrower: br3libPtr}

var bk4libPtr = &Book{Title: "Title4", Author: "Author4", Borrower: br3libPtr}

var bks1 = []*Book{bk1libPtr, bk2libPtr}
var bks2 = []*Book{bk1libPtr, bk2libPtr, bk3libPtr}

var bks3 = []*Book{bk1libPtr, bk2libPtr, bk3libPtr, bk4libPtr}

//var b = `{"firstName":"John","lastName":"Dow"}`
//var jsonStringBorrowers = "[\n  {\n    \"name\": \"Borrower1\",\n    \"max-books\": 1\n  },\n  {\n    \"name\": \"Borrower2\",\n    \"max-books\": 2\n  }\n]"
//var jsonStringBooks = "[\n  {\n    \"title\": \"Title2\",\n    \"author\": \"Author22\",\n    \"borrower\": {\n      \"name\": \"NoName\",\n      \"max-books\": -1\n    }\n  }\n]"
//var jsonStringBorrowersBadParse = "[{\"name\"\"Borrower1\",\"max-books\":1},{\"name\":\"Borrower2\",\"max-books\":2}]"
//var jsonStringBorrowersBadNameField = "[{\"noName\":\"Borrower1\",\"max-books\":1},{\"name\":\"Borrower2\",\"max-books\":2}]"
//var jsonStringBooksBadParse = "[{\"title\"\"Title2\",\"author\":\"Author22\",\"borrower\":{\"name\":\"NoName\",\"max-books\":-1}}, {\"title\":\"Title99\",\"author\":\"Author99\",\"borrower\":{\"name\":\"Borrower1\",\"max-books\":1}}]"
//var jsonStringBooksBadTitleField = "[{\"noTitle\":\"Title2\",\"author\":\"Author22\",\"borrower\":{\"name\":\"NoName\",\"max-books\":-1}}, {\"title\":\"Title99\",\"author\":\"Author99\",\"borrower\":{\"name\":\"Borrower1\",\"max-books\":1}}]"
//var jsonStringBooksBadBorrowerField = "[{\"title\":\"Title2\",\"author\":\"Author22\",\"borrower\":{\"noName\":\"NoName\",\"max-books\":-1}}, {\"title\":\"Title99\",\"author\":\"Author99\",\"borrower\":{\"name\":\"Borrower1\",\"max-books\":1}}]"
//var ss = "\n--- Status Report of Test Library ---\n\nTest Library: 3 books; 3 borrowers.\n\nTitle1 by Author1; Checked out to Borrower1\nTitle2 by Author2; Available\nTitle3 by Author3; Checked out to Borrower3\n\nBorrower1 (1 books)\nBorrower2 (2 books)\nBorrower3 (3 books)\n\n--- End of Status Report ---\n"

func TestAddBorrower(t *testing.T) {
	cases := []struct {
		brs     []*Borrower
		br      *Borrower
		wantBrs []*Borrower
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
		bks     []*Book
		bk      *Book
		wantBks []*Book
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

func TestFindBorrower(t *testing.T) {
	var ErrNoBorrowerFound = errors.New("did not find the requested borrower")
	cases := []struct {
		n       string
		brs     []*Borrower
		wantBr  *Borrower
		wantErr error
	}{
		{"Borrower1", brs2, br1libPtr, nil},
		{"Borrower11", brs2, &Borrower{}, ErrNoBorrowerFound},
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

func TestFindBook(t *testing.T) {
	var ErrNoBookFound = errors.New("did not find the requested book")
	cases := []struct {
		t       string
		bks     []*Book
		wantI   int
		wantBk  *Book
		wantErr error
	}{
		{"Title1", bks2, 0, bk1libPtr, nil},
		{"Title11", bks2, 0, &Book{}, ErrNoBookFound},
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

func TestGetBooksForBorrower(t *testing.T) {
	cases := []struct {
		br   *Borrower
		bks  []*Book
		want []*Book
	}{
		{br2libPtr, bks1, []*Book{}},
		{br1libPtr, bks1, []*Book{bk1libPtr}},
		{br3libPtr, bks3, []*Book{bk3libPtr, bk4libPtr}},
	}
	for _, c := range cases {
		got := getBooksForBorrower(c.br, c.bks)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("GetBooksForborrower(%v, %v) ==\n%v want \n%v",
				c.br, c.bks, got, c.want)
		}
	}
}

//func TestCheckOut(t *testing.T) {
//	var bks2 = []Book{*bk1libPtr, {Title: "Title2", Author: "Author2", Borrower: &Borrower{Name: "Borrower2", MaxBooks: 2}}}
//	cases := []struct {
//		n    string
//		t    string
//		brs  []Borrower
//		bks  []Book
//		want []Book
//	}{
//		{"Borrower2", "Title1", brs1, bks1, bks1},
//		{"Borrower2", "NoTitle", brs1, bks1, bks1},
//		{"NoName", "Title1", brs1, bks1, bks1},
//		{"Borrower1", "Title2", brs1, bks1, bks1},
//		{"Borrower2", "Title2", brs1, bks1, bks2},
//	}
//	for _, c := range cases {
//		got := CheckOut(c.n, c.t, c.brs, c.bks)
//		if !reflect.DeepEqual(got, c.want) {
//			t.Errorf("CheckOut(%s, %s, %v, %v) ==\n%v want \n%v",
//				c.n, c.t, c.brs, c.bks, got, c.want)
//		}
//	}
//}
//
//func TestCheckIn(t *testing.T) {
//	var bks1 = []Book{*bk1libPtr, *bk2libPtr}
//	var bks2 = []Book{{Title: "Title1", Author: "Author1", Borrower: &Borrower{}}, *bk2libPtr}
//	cases := []struct {
//		t    string
//		bks  []Book
//		want []Book
//	}{
//		{"Title1", bks1, bks2},
//		{"Title2", bks1, bks1},
//		{"NoTitle", bks1, bks1},
//	}
//	for _, c := range cases {
//		got := CheckIn(c.t, c.bks)
//		if !reflect.DeepEqual(got, c.want) {
//			t.Errorf("CheckIn(%s, %v) ==\n%v want \n%v",
//				c.t, c.bks, got, c.want)
//		}
//	}
//}
//
//func TestJSONStringToBorrowersPass(t *testing.T) {
//	js := jsonStringBorrowers
//	wantBrs := brs1
//	wantE := error(nil)
//
//	got, err := JSONStringToBorrowers(js)
//	if !reflect.DeepEqual(got, wantBrs) || err != wantE {
//		t.Errorf("JSONStringToBorrowers\n(%s)\n==\n%v and %v\nwant\n%v and %v",
//			js, got, err, wantBrs, wantE)
//	}
//}
//
//func TestJSONStringToBorrowersFail(t *testing.T) {
//	cases := []struct {
//		js      string
//		wantBrs []Borrower
//		wantE   error
//	}{
//		{jsonStringBorrowersBadParse, []Borrower{}, errors.New("invalid character '\"' after object key")},
//		{jsonStringBorrowersBadNameField, []Borrower{}, errors.New("missing Borrower field value - borrowers list is empty")},
//	}
//	for _, c := range cases {
//		got, err := JSONStringToBorrowers(c.js)
//		if err.Error() != c.wantE.Error() {
//			t.Errorf("JSONStringToBorrowers\n(%s)\n==\n%v and %v\nwant\n%v and %v",
//				c.js, got, err, c.wantBrs, c.wantE)
//		}
//	}
//}
//
//func TestJSONStringToBooksPass(t *testing.T) {
//	js := jsonStringBooks
//	wantBks := []Book{{Title: "Title2", Author: "Author22", Borrower: &Borrower{Name: "NoName", MaxBooks: -1}}}
//	wantE := error(nil)
//
//	got, err := JSONStringToBooks(js)
//
//	if !reflect.DeepEqual(got, wantBks) || err != wantE {
//		t.Errorf("JSONStringToBooks\n(%s)\n==\n%v and %v\nwant\n%v and %v",
//			js, got, err, wantBks, wantE)
//	}
//}
//
//func TestJSONStringToBooksFail(t *testing.T) {
//	cases := []struct {
//		js      string
//		wantBks []Book
//		wantE   error
//	}{
//		{jsonStringBooksBadParse, []Book{}, errors.New("invalid character '\"' after object key")},
//		{jsonStringBooksBadTitleField, []Book{}, errors.New("missing Book field value - book list is empty")},
//		{jsonStringBooksBadBorrowerField, []Book{}, errors.New("missing Book field value - book list is empty")},
//	}
//	for _, c := range cases {
//		got, err := JSONStringToBooks(c.js)
//		if err.Error() != c.wantE.Error() {
//			t.Errorf("JSONStringToBooks\n(%s)\n==\n%v and %v\nwant\n%v and %v",
//				c.js, got, err, c.wantBks, c.wantE)
//		}
//	}
//}
//
//func TestBorrowersToJSONString(t *testing.T) {
//	brs := brs1
//	got := BorrowersToJSONSting(brs)
//	want := jsonStringBorrowers
//	if got != want {
//		t.Errorf("BorrwersToJSONSting(%v) ==\n(%q) want \n(%q)",
//			brs, got, want)
//	}
//}
//
//func TestBooksToJSONString(t *testing.T) {
//	bks := []Book{{Title: "Title2", Author: "Author22", Borrower: &Borrower{Name: "NoName", MaxBooks: -1}}}
//	got := BooksToJSONSting(bks)
//	want := jsonStringBooks
//	if got != want {
//		t.Errorf("BooksToJSONSting(%v) ==\n(%q) want \n(%q)",
//			bks, got, want)
//	}
//}
//
//func TestStatusToString(t *testing.T) {
//	bks := bks2
//	brs := brs2
//	got := StatusToString(bks, brs)
//	want := ss
//	if got != want {
//		t.Errorf("StatusToString(%v, %v) ==\n(%q) want \n(%q)",
//			bks, brs, got, want)
//	}
//}
