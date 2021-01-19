package book

import (
	"eatobin.com/totalbeginnergo/borrower"
	"testing"
)

var bk1, _ = JsonStringToBook("{\"title\":\"Title1\",\"author\":\"Author1\"}")
var newBorrower = borrower.NewBorrower("Borrower1", 1)
var bk2, _ = JsonStringToBook("{\"title\":\"Title1\",\"author\":\"Author1\",\"borrower\":null}")
var bk3, _ = JsonStringToBook("{\"title\":\"Title1\",\"author\":\"Author1\",\"borrower\":{\"name\":\"Borrower1\",\"maxBooks\":1}}")
var wantAvail = "Title1 by Author1; Available"
var wantCheckedOut = "Title1 by Author1; Checked out to Borrower1"

func TestSetBorrower(t *testing.T) {
	cases := []struct {
		bk           Book
		br           *borrower.Borrower
		wantBkString string
	}{
		{bk1, &newBorrower, wantCheckedOut},
		{bk2, &newBorrower, wantCheckedOut},
		{bk3, nil, wantAvail},
	}
	for _, c := range cases {
		gotBkString := String(SetBorrower(c.bk, c.br))
		if gotBkString != c.wantBkString {
			t.Fatalf("SetBorrower(%v, %v) == %v, want %v", c.bk, c.br, gotBkString, c.wantBkString)
		}
		//if !reflect.DeepEqual(gotBrs, c.wantBrs) {
		//	t.Errorf("AddBorrower(%v, %v) ==\n%v want\n%v",
		//		c.brs, c.br, gotBrs, c.wantBrs)
		//}
	}
}

//func TestSetBookValues(t *testing.T) {
//	badBkAvail, _ := JsonStringToBook(bk1)
//	title := "Title1"
//	gotBkT := String(SetTitle(badBkAvail, title))
//	if gotBkT != wantAvailS1 {
//		t.Fatalf("SetTitle(%v, %v) == %v, want %v", badBkAvail, title, gotBkT, wantAvailS1)
//	}
//	author := "Author1"
//	gotBkA := String(SetAuthor(badBkAvail, author))
//	if gotBkA != wantAvailS2 {
//		t.Fatalf("SetAuthor(%v, %v) == %v, want %v", badBkAvail, author, gotBkA, wantAvailS2)
//	}
//	bkNotAvail, _ := JsonStringToBook(bk2)
//	wantNotAvailS := String(bkNotAvail)
//	br2 := borrower.NewBorrower("Borrower2", 2)
//	gotBkB := String(SetBorrower(badBkAvail, &br2))
//	if gotBkB != wantNotAvailS {
//		t.Fatalf("SetBorrower(%v, %v) == %v, want %v", badBkAvail, br2, gotBkB, wantNotAvailS)
//	}
//	fmt.Println(BkToJsonString(NewBook("me", "you")))
//	fmt.Println(BkToJsonString(Book{
//		Title:  "meeToo",
//		Author: "youToo",
//		Borrower: &borrower.Borrower{
//			Name:     "br",
//			MaxBooks: 33,
//		},
//	}))
//	fmt.Println(BkToJsonString(Book{
//		Title:  "meeToo2",
//		Author: "youToo2",
//	}))
//	gotBk, _ := JsonStringToBook(jsonStringBk3)
//	wantBk := Book{Title: "Title11", Author: "Author11"}
//	if String(gotBk) != String(wantBk) {
//		t.Fatalf("JsonStringToBook(%v) == %v, want %v", jsonStringBk3, gotBk, wantBk)
//	}
//}
