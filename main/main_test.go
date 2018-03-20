package main

import (
	"testing"
)

// var brsJ = readFileIntoJsonString("borrowers-before.json")
// var bksJ = readFileIntoJsonString("books-before.json")

func TestReadFileIntoJSONString(t *testing.T) {
	cases := []struct {
		f    string
		want string
	}{
		{"borrowers-before.json", ""},
		// 		{brs1, br2lib, brs1},
	}
	for _, c := range cases {
		got := ReadFileIntoJsonString(c.f)
		if got != c.want {
			t.Errorf("ReadFileIntoJsonString(%s) ==\n%s want\n%s",
				c.f, got, c.want)
		}
	}
}
