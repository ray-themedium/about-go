package main

import (
	"fmt"
)

type Product struct {
	Name string
	Price int
}

func main() {
	m := make(map[int]Product)

	m[16] = Product{"볼팬", 500}
	m[46] = Product{"ㅎㄹ", 400}
	m[111] = Product{"ㅗㅇㅎㄹ", 100}
	m[5] = Product{"ㅇㄹㅎ", 200}
	m[324] = Product{"432", 1000}

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, 16)

	for k, v := range m {
		fmt.Println(k, v)
	}

	v, ok := m[16]
	fmt.Println(v, ok)
	// m := make(map[string]string) // 맵 생성
	// m["임태호1"] = "경기도 안양시1"
	// m["임태호2"] = "경기도 안양시2"
	// m["임태호3"] = "경기도 안양시3"
	// m["임태호4"] = "경기도 안양시4"

	// m["임태호3"] = "석수동"

	// fmt.Println(m["임태호1"])
	// fmt.Println(m["임태호3"])
}