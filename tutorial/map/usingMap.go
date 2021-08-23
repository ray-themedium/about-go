package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 12
	m["k2"] = 13
	fmt.Println(m)

	mm := map[string]int{
		"k1": 12,
		"k2": 12,
	}
	fmt.Println(mm)

	delete(mm, "k1")
	delete(mm, "k1")
	delete(mm, "k1")
	fmt.Println(mm)

	// 맵을 접근할 때 원하는 정보가 있는지 알아날 때 사용하는 기법
	if _, ok := m["hi"]; ok {
		fmt.Println("Exists!")
	} else {
		fmt.Println("Does not exist")
	}

	for key, value := range m {
		fmt.Println(key, value)
	}
}
