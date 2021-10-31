// 구조를 모르는 json 데이터를 읽고 파싱하기
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args
	filename := args[1]

	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	var parseData map[string]interface{}
	json.Unmarshal([]byte(fileData), &parseData)

	for key, value := range parseData {
		fmt.Println("key:", key, "value:", value)
	}

}
