package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide a filename!")
		return
	}
	filename := args[1]

	fileData, err := ioutil.ReadFile(filename) // 파일 전체를 한 번에 다 읽어옴
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(fileData))

	var parseData map[string]interface{}
	json.Unmarshal(fileData, &parseData) // json의 구조를 모르고 파싱할 땐 map으로 해야 함

	for key, value := range parseData {
		fmt.Println("key:", key, "value:", value)
	}
}
