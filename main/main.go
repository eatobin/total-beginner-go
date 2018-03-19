package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Tentacle a character from Day of Tentacles
type Tentacle struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (t Tentacle) toString() string {
	bytes, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(bytes)
}

func toTentacles(ts string) []Tentacle {
	var tentacles []Tentacle
	json.Unmarshal([]byte(ts), &tentacles)
	return tentacles
}

func readFileIntoJsonString(f string) string {
	raw, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(raw)
}

func main() {
	tentaclesString := readFileIntoJsonString("books-before.json")
	fmt.Printf("%q\n", tentaclesString)
	//tentacles := toTentacles(tentaclesString)
	//for _, te := range tentacles {
	//	fmt.Println(te.toString())
	//}
}
