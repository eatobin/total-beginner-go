package totalbeginnergo

import (
	"reflect"
	"testing"
)

var br1lib = Borrower{"Borrower1", 1}
var br2lib = Borrower{"Borrower2", 2}
var br3 = Borrower{"Borrower3", 3}

var brs1 = []Borrower{br1lib, br2lib}
var brs2 = []Borrower{br1lib, br2lib, br3}

var bk1lib = Book{"Title1", "Author1", br1lib}
var bk2 = Book{"Title2", "Author2", Borrower{"NoName", -1}}
var bk3 = Book{"Title3", "Author3", br3}
var bk4 = Book{"Title4", "Author4", Borrower{"Borrower3", 3}}

var bks1 = []Book{bk1lib, bk2}
var bks2 = []Book{bk1lib, bk2, bk3}
var bks3 = []Book{bk1lib, bk2, bk3, bk4}

var jsonStringBorrowers = "[{\"name\":\"Borrower1\",\"max-books\":1},{\"name\":\"Borrower2\",\"max-books\":2}]"
var jsonStringBooks = "[{\"title\":\"Title2\",\"author\":\"Author22\",\"borrower\":{\"name\":\"NoName\",\"max-books\":-1}}]"
var jsonStringBorrowersBadParse = "[{\"name\"\"Borrower1\",\"max-books\":1},{\"name\":\"Borrower2\",\"max-books\":2}]"
var jsonStringBorrowersBadNameField = "[{\"noName\":\"Borrower1\",\"max-books\":1},{\"name\":\"Borrower2\",\"max-books\":2}]"
var jsonStringBooksBadParse = "[{\"title\"\"Title2\",\"author\":\"Author22\",\"borrower\":{\"name\":\"NoName\",\"max-books\":-1}}, {\"title\":\"Title99\",\"author\":\"Author99\",\"borrower\":{\"name\":\"Borrower1\",\"max-books\":1}}]"
var jsonStringBooksBadTitleField = "[{\"noTitle\":\"Title2\",\"author\":\"Author22\",\"borrower\":{\"name\":\"NoName\",\"max-books\":-1}}, {\"title\":\"Title99\",\"author\":\"Author99\",\"borrower\":{\"name\":\"Borrower1\",\"max-books\":1}}]"
var jsonStringBooksBadBorrowerField = "[{\"title\":\"Title2\",\"author\":\"Author22\",\"borrower\":{\"noName\":\"NoName\",\"max-books\":-1}}, {\"title\":\"Title99\",\"author\":\"Author99\",\"borrower\":{\"name\":\"Borrower1\",\"max-books\":1}}]"
var ss = "\n--- Status Report of Test Library ---\n\nTest Library: 3 books; 3 borrowers.\n\nTitle1 by Author1; Checked out to Borrower1\nTitle2 by Author2; Available\nTitle3 by Author3; Checked out to Borrower3\n\nBorrower1 (1 books)\nBorrower2 (2 books)\nBorrower3 (3 books)\n\n--- End of Status Report ---\n\n"

func TestAddBorrower(t *testing.T) {
	cases := []struct {
		brs     []Borrower
		br      Borrower
		wantBrs []Borrower
	}{
		{brs1, br3, brs2},
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

func TestFindBorrower(t *testing.T) {
	cases := []struct {
		n    string
		brs  []Borrower
		want Borrower
	}{
		{"Borrower1", brs2, br1lib},
		{"Borrower11", brs2, Borrower{}},
	}
	for _, c := range cases {
		got := FindBorrower(c.n, c.brs)
		if got != c.want {
			t.Errorf("FindBorrower(%s, %v) ==\n%v want \n%v",
				c.n, c.brs, got, c.want)
		}
	}
}

func TestFindBook(t *testing.T) {
	cases := []struct {
		t    string
		bks  []Book
		want Book
	}{
		{"Title1", bks2, bk1lib},
		{"Title11", bks2, Book{}},
	}
	for _, c := range cases {
		_, got := FindBook(c.t, c.bks)
		if got != c.want {
			t.Errorf("FindBook(%s, %v) ==\n%v want \n%v",
				c.t, c.bks, got, c.want)
		}
	}
}

func TestGetBooksForBorrower(t *testing.T) {
	cases := []struct {
		br   Borrower
		bks  []Book
		want []Book
	}{
		{br2lib, bks1, []Book{}},
		{br1lib, bks1, []Book{bk1lib}},
		{br3, bks3, []Book{bk3, bk4}},
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
	var bks2 = []Book{bk1lib, {Title: "Title2", Author: "Author2", Borrower: Borrower{Name: "Borrower2", MaxBooks: 2}}}
	cases := []struct {
		n    string
		t    string
		brs  []Borrower
		bks  []Book
		want []Book
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
	var bks1 = []Book{bk1lib, bk2}
	var bks2 = []Book{{Title: "Title1", Author: "Author1", Borrower: Borrower{Name: "NoName", MaxBooks: -1}}, bk2}
	cases := []struct {
		t    string
		bks  []Book
		want []Book
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

func TestJSONStringToBorrowers(t *testing.T) {
	cases := []struct {
		js      string
		wantBrs []Borrower
		wantErr string
	}{
		{jsonStringBorrowers, brs1, ""},
		{jsonStringBorrowersBadParse, []Borrower{}, "JSON parse error."},
		{jsonStringBorrowersBadNameField, []Borrower{}, "Missing Borrower field value."},
	}
	for _, c := range cases {
		got, err := JSONStringToBorrowers(c.js)
		if !reflect.DeepEqual(got, c.wantBrs) || err != c.wantErr {
			t.Errorf("JSONStringToBorrowers(%s) ==\n(%v, %v) want \n(%v, %v)",
				c.js, got, err, c.wantBrs, c.wantErr)
		}
	}
}

func TestJSONStringToBooks(t *testing.T) {
	cases := []struct {
		js      string
		wantBks []Book
		wantErr string
	}{
		{jsonStringBooks,
			[]Book{{Title: "Title2", Author: "Author22", Borrower: Borrower{Name: "NoName", MaxBooks: -1}}}, ""},
		{jsonStringBooksBadParse, []Book{}, "JSON parse error."},
		{jsonStringBooksBadTitleField, []Book{}, "Missing Book field value."},
		{jsonStringBooksBadBorrowerField, []Book{}, "Missing Book field value."},
	}
	for _, c := range cases {
		got, err := JSONStringToBooks(c.js)
		if !reflect.DeepEqual(got, c.wantBks) || err != c.wantErr {
			t.Errorf("JSONStringToBooks(%s) ==\n(%v, %v) want \n(%v, %v)",
				c.js, got, err, c.wantBks, c.wantErr)
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
	bks := []Book{{Title: "Title2", Author: "Author22", Borrower: Borrower{Name: "NoName", MaxBooks: -1}}}
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
