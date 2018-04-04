package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"testing"
)

var brsJ = "[\n  {\n    \"name\": \"Borrower100\",\n    \"max-books\": 100\n  },\n  {\n    \"name\": \"Borrower200\",\n    \"max-books\": 200\n  }\n]\n"
var bksJ = "[\n  {\n    \"title\": \"Book100\",\n    \"author\": \"Author100\",\n    \"borrower\": {\n      \"name\": \"Borrower100\",\n      \"max-books\": 100\n    }\n  },\n  {\n    \"title\": \"Book200\",\n    \"author\": \"Author200\",\n    \"borrower\": {\n      \"name\": \"NoName\",\n      \"max-books\": -1\n    }\n  }\n]\n"

func TestReadFileIntoJSONStringPass(t *testing.T) {
	cases := []struct {
		f     string
		wantS string
		wantE error
	}{
		{"../borrowers-before.json", brsJ, nil},
		{"../books-before.json", bksJ, nil},
	}
	for _, c := range cases {
		got, err := ReadFileIntoJsonString(c.f)
		if got != c.wantS || err != c.wantE {
			t.Errorf("ReadFileIntoJsonString(%s) ==\n%s and %v\nwant\n%s and %v",
				c.f, got, err, c.wantS, c.wantE)
		}
	}
}

func TestReadFileIntoJSONStringFail(t *testing.T) {
	f := "../NoSuch.json"
	wantE := errors.New("open ../NoSuch.json: no such file or directory")

	_, err := ReadFileIntoJsonString(f)

	if err.Error() != wantE.Error() {
		t.Errorf("ReadFileIntoJsonString(%s) ==\n%v\nwant\n%v",
			f, err, wantE)
	}

}

func TestWriteJSONStringToFile(t *testing.T) {
	cases := []struct {
		js    string
		f1    string
		f2    string
		wantE error
	}{
		{brsJ, "../borrowers-before2.json", "../borrowers-before.json", nil},
		{bksJ, "../books-before2.json", "../books-before.json", nil},
	}
	for _, c := range cases {
		err := WriteJSONStringToFile(c.js, c.f1)
		f1, _ := ioutil.ReadFile(c.f1)
		f2, _ := ioutil.ReadFile(c.f2)
		if !bytes.Equal(f1, f2) || err != c.wantE {
			t.Errorf("WriteJSONStringToFile(\n%s, %s)\n\n(%v)\n\n!=\n\n%s\n(%v)",
				c.js, c.f1, err, f2, c.wantE)
		}
	}
}
