package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["key"] = "value"
	m["key2"] = "value2"
	m["key3"] = "value3"
	m["key4"] = "value4"

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println()
	delete(m, "key3")

	for k, v := range m {
		fmt.Println(k, v)
	}
}
