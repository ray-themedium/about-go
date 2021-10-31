package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func saveToJSON(filename *os.File, key interface{}) {
	encodeJSON := json.NewEncoder(filename)
	err := encodeJSON.Encode(key)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	myRecord := Record{
		"taeho",
		"bbaktaeho",
		[]Telephone{
			{true, "1234-1234"},
			{false, "1233-1234"},
			{true, "1232-1234"},
		},
	}
	file, _ := os.Create("saveJSON.json")
	saveToJSON(file, myRecord)
}
