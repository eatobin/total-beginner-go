package main

import (
	"bytes"
	"io/ioutil"
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

func TestWriteJSONStringToFile(t *testing.T) {
	cases := []struct {
		js string
		f1 string
		f2 string
	}{
		{brsJ, "../borrowers-before2.json", "../borrowers-before.json"},
		{bksJ, "../books-before2.json", "../books-before.json"},
	}
	for _, c := range cases {
		WriteJSONStringToFile(c.js, c.f1)
		f1, err1 := ioutil.ReadFile(c.f1)
		if err1 != nil {
			t.Error(err1)
		}
		f2, err2 := ioutil.ReadFile(c.f2)
		if err2 != nil {
			t.Error(err2)
		}
		if !bytes.Equal(f1, f2) {
			t.Errorf("WriteJSONStringToFile(\n%s, %s)\n\n!=\n\n%s",
				c.js, c.f1, f2)
		}
	}
}