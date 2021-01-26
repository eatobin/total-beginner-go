package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestReadFileIntoJsonString(t *testing.T) {
	cases := []struct {
		fp             string
		wantJsonString string
		wantError      error
	}{
		{"resources/noFile.txt", "'no-string'", errors.New("open resources/noFile.txt: no such file or directory")},
		//{brs1, br2lib, brs1},
	}
	for _, c := range cases {
		gotString, gotError := ReadFileIntoJsonString(c.fp)
		if gotError != nil {
			if gotError.Error() != c.wantError.Error() {
				t.Errorf("ReadFileIntoJsonString(%s)\n==\n'no-string' and %v\nwant\n%v and %v",
					c.fp, gotError, c.wantJsonString, c.wantError)
			}
		}
		fmt.Println(gotString)
	}
}
