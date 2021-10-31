package main

import (
	"encoding/json"
	"fmt"
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
	rec, err := json.Marshal(&myRecord)
	if err != nil {
		fmt.Println(err)
		return
	}
	unRec := new(Record)
	fmt.Println(string(rec))
	json.Unmarshal(rec, &unRec)
	fmt.Println(unRec)
}
