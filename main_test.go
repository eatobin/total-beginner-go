package main

import (
	"eatobin.com/totalbeginnergo/library"
	"errors"
	"testing"
)

func TestReadFileIntoJsonString(t *testing.T) {
	cases := []struct {
		fp             string
		wantJsonString string
		wantError      error
	}{
		{"resources-test/noFile.txt", "", errors.New("open resources-test/noFile.txt: no such file or directory")},
		{"resources-test/testText.txt", "This is test text\n", errors.New("")},
	}
	for _, c := range cases {
		gotString, gotError := readFileIntoJsonString(c.fp)
		if gotError != nil {
			if gotError.Error() != c.wantError.Error() {
				t.Errorf("readFileIntoJsonString(%s) ==\n%s\n%s",
					c.fp, gotError, c.wantError)
			}
		}
		if gotString != c.wantJsonString {
			t.Errorf("readFileIntoJsonString(%s) ==\n%s%s",
				c.fp, gotString, c.wantJsonString)
		}
	}
}

func TestNewV(t *testing.T) {
	cases := []struct {
		brsfp   string
		bksfp   string
		wantLib string
	}{
		{"resources/borrowers-before.json", "resources/books-before.json",
			"\n--- Status Report of Test Library ---\n\nTest Library: 2 books; 2 borrowers.\n\nBook100 by Author100; Checked out to Borrower100\nBook200 by Author200; Available\n\nBorrower100 (100 books)\nBorrower200 (200 books)\n\n--- End of Status Report ---\n"},
		{"resources/bad-borrowers.json", "resources/books-before.json",
			"\n--- Status Report of Test Library ---\n\nTest Library: 0 books; 0 borrowers.\n\n\n\n--- End of Status Report ---\n"},
		{"resources/noFile.json", "resources/books-before.json",
			"\n--- Status Report of Test Library ---\n\nTest Library: 0 books; 0 borrowers.\n\n\n\n--- End of Status Report ---\n"},
		{"resources/empty.json", "resources/books-before.json",
			"\n--- Status Report of Test Library ---\n\nTest Library: 0 books; 0 borrowers.\n\n\n\n--- End of Status Report ---\n"},
	}
	for _, c := range cases {
		borrowers := library.ZeroBorrowers
		books := library.ZeroBooks
		borrowers, books = NewV(c.brsfp, c.bksfp)
		gotLib := library.StatusToString(books, borrowers)
		if gotLib != c.wantLib {
			t.Errorf("NewV(%s, %s) ==\n%s\nwant\n%s",
				c.brsfp, c.bksfp, gotLib, c.wantLib)
		}
	}
}

func TestWriteJsonStringToFile(t *testing.T) {
	cases := []struct {
		fp        string
		txt       string
		wantError error
	}{
		{"resourcesX/borrowers-after.json", "This is test text\n", errors.New("open resourcesX/borrowers-after.json: no such file or directory")},
		//{"resources-test/testText.txt", "This is test text\n", errors.New("")},
	}
	for _, c := range cases {
		gotError := writeJsonStringToFile(c.fp, c.txt)
		if gotError != nil {
			if gotError.Error() != c.wantError.Error() {
				t.Errorf("writeJsonStringToFile(%s, %s) ==\n%s\nwant\n%s",
					c.fp, c.txt, gotError, c.wantError)
			}
		}
	}
}
