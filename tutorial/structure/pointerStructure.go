package main

import "fmt"

type myS struct {
	Name    string
	Surname string
	Height  int32
}

func createStruct(n, s string, h int32) *myS {
	if h > 300 {
		h = 0
	}
	return &myS{n, s, h}
}

func refStruct(n, s string, h int32) myS {
	if h > 300 {
		h = 0
	}
	return myS{n, s, h}
}

func main() {
	n1 := new(myS)
	s1 := createStruct("tae", "ho", 123)
	s3 := createStruct("bbak", "taeho", 400)
	s2 := refStruct("tae", "ho", 123)
	fmt.Println(n1)
	fmt.Println((*s1).Name)
	fmt.Println((*s3).Name)
	fmt.Println(s2.Name)
	fmt.Println(s1)
	fmt.Println(s3)
	fmt.Println(s2)
}
