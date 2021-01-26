package main

import (
	"errors"
	"testing"
)

func TestReadFileIntoJsonString(t *testing.T) {
	cases := []struct {
		fp             string
		wantJsonString string
		wantError      error
	}{
		{"resources/noFile.txt", "", errors.New("open resources/noFile.txt: no such file or directory")},
		{"resources/testText.txt", "This is test text\n", errors.New("")},
	}
	for _, c := range cases {
		gotString, gotError := ReadFileIntoJsonString(c.fp)
		if gotError != nil {
			if gotError.Error() != c.wantError.Error() {
				t.Errorf("ReadFileIntoJsonString(%s) ==\n%s\n%s",
					c.fp, gotError, c.wantError)
			}
		}
		if gotString != c.wantJsonString {
			t.Errorf("ReadFileIntoJsonString(%s) ==\n%s%s",
				c.fp, gotString, c.wantJsonString)
		}
	}
}
