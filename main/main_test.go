package main

import (
	"testing"
)

var brsJ = "[\n  {\n    \"name\": \"Borrower100\",\n    \"max-books\": 100\n  },\n  {\n    \"name\": \"Borrower200\",\n    \"max-books\": 200\n  }\n]\n"
var bksJ = "[\n  {\n    \"title\": \"Book100\",\n    \"author\": \"Author100\",\n    \"borrower\": {\n      \"name\": \"Borrower100\",\n      \"max-books\": 100\n    }\n  },\n  {\n    \"title\": \"Book200\",\n    \"author\": \"Author200\",\n    \"borrower\": {\n      \"name\": \"NoName\",\n      \"max-books\": -1\n    }\n  }\n]\n"

func TestReadFileIntoJSONString(t *testing.T) {
	cases := []struct {
		f    string
		want string
	}{
		{"../borrowers-before.json", brsJ},
		{"../books-before.json", bksJ},
		{"../NoSuch.json", ""},
	}
	for _, c := range cases {
		got := ReadFileIntoJsonString(c.f)
		if got != c.want {
			t.Errorf("ReadFileIntoJsonString(%s) ==\n%s want\n%s",
				c.f, got, c.want)
		}
	}
}
