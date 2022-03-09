package main

import (
	"fmt"
	"sort"
)

type A struct {
	a string
	b int
}

type As []A

// sort.Interface 타입으로 맞춤
func (as As) Len() int           { return len(as) }
func (as As) Less(i, j int) bool { return as[i].b < as[j].b }
func (as As) Swap(i, j int)      { as[i], as[j] = as[j], as[i] }

func main() {
	as := []A{
		{"aa", 5},
		{"bb", 9},
		{"cc", 1},
		{"dd", 8},
	}

	sort.Sort(As(as))
	fmt.Println(as)
}
