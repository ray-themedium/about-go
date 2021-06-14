package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9 // 역 참조를 명시할 필요 없다.
	fmt.Println(v)
}