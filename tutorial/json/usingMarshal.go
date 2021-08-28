package main

import (
	"encoding/json"
	"fmt"
)

type Record struct {
	Id   string
	Name string
}

func main() {
	record := Record{
		Id:   "1",
		Name: "taeho",
	}
	rec, err := json.Marshal(&record) // 마샬 함수는 포인터 변수를 받아서 json 포멧으로 변환한다.
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(rec))

	var unRec Record
	err1 := json.Unmarshal(rec, &unRec)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(unRec)
}
