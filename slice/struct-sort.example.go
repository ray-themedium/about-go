package main

import (
	"fmt"
	"sort"
)

type A struct {
	a string
	b int
}

type B struct {
	a string
	b int
}

type As []A

func (as As) Len() int           { return len(as) }
func (as As) Less(i, j int) bool { return as[i].b < as[j].b }
func (as As) Swap(i, j int)      { as[i], as[j] = as[j], as[i] }

// sort.Interface 타입을 지원하는 구조체 정렬 예제
func ExampleSortInterfaceSort() {
	fmt.Println("start ExampleSortInterfaceSort()")
	defer fmt.Println("end ExampleSortInterfaceSort()")
	as := []A{
		{"aa", 5},
		{"bb", 9},
		{"cc", 1},
		{"dd", 8},
	}

	sort.Sort(As(as))
	fmt.Println(as)
}

// sort.Interface 타입을 지원하지 않는 구조체 정렬 예제
func ExampleDefaultStructSort() {
	fmt.Println("start ExampleDefaultStructSort()")
	defer fmt.Println("end ExampleDefaultStructSort()")
	bs := []B{
		{"aa", 5},
		{"bb", 9},
		{"cc", 1},
		{"dd", 8},
	}

	sort.Slice(bs, func(i, j int) bool { return bs[i].b < bs[j].b })
	fmt.Println(bs)
}

func main() {
	ExampleDefaultStructSort()
	ExampleSortInterfaceSort()
}
