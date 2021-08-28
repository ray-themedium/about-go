package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Name    string `json:"id"`
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func loadFromJSON(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	decodeJSON := json.NewDecoder(in)
	err = decodeJSON.Decode(key)
	if err != nil {
		return err
	}
	in.Close()
	return nil
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
	// args := os.Args
	// if len(args) == 1 {
	// 	fmt.Println("Please provide a filename!")
	// 	return
	// }
	// filename := args[1]
	// var myRecord Record
	// err := loadFromJSON(filename, &myRecord)
	// if err == nil {
	// 	fmt.Println(myRecord)
	// } else {
	// 	fmt.Println(err)
	// }

	fmt.Println("\n----- save ------")
	writeRecord := Record{
		Name:    "test",
		Surname: "test2",
		Tel: []Telephone{
			{Mobile: true, Number: "11-11"},
			{Mobile: true, Number: "11-12"},
			{Mobile: true, Number: "11-13"},
		},
	}
	saveToJSON(os.Stdout, writeRecord)
}
