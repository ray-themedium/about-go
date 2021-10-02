package main

import (
	"fmt"
	"sort"
)

func main() {
	SortForInts()
	SortForStrings()
	SortReverse1()
	SortReverse2()
	SortForStructs()
}

// 기본 타입 슬라이스 정렬
func SortForInts() {
	slice := []int{5, 7, 9, 4, 2, 3, 1, 6, 8}
	sort.Ints(slice)
	fmt.Println(slice, sort.IntsAreSorted(slice))
}
func SortForStrings() {
	slice := []string{"b", "bb", "c", "ab", "aa", "abb", "acc"}
	sort.Strings(slice)
	fmt.Println(slice, sort.StringsAreSorted(slice))
}
func SortReverse1() {
	slice := []int{5, 7, 9, 4, 2, 3, 1, 6, 8}
	sort.Sort(sort.Reverse(sort.IntSlice(slice))) // 정렬 인터페이스에 맞게 메서드를 붙여줌
	fmt.Println(slice, sort.IntsAreSorted(slice))
}
func SortReverse2() {
	slice := []int{5, 7, 9, 4, 2, 3, 1, 6, 8}
	sort.Slice(slice, func(i, j int) bool { // 익명함수를 활용한 정렬 방법
		return slice[i] > slice[j]
	})
	fmt.Println(slice, sort.IntsAreSorted(slice))
}

type St struct {
	Name string
	Age  int
}
type Sts []St

func (a Sts) Len() int           { return len(a) }
func (a Sts) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Sts) Less(i, j int) bool { return a[i].Age < a[j].Age } // sort.Interface

// 구조체 정렬
func SortForStructs() {
	sts := []St{
		{Name: "ss1", Age: 10},
		{Name: "ss2", Age: 50},
		{Name: "ss3", Age: 20},
		{Name: "ss4", Age: 33},
		{Name: "ss5", Age: 5},
		{Name: "ss6", Age: 11},
	}
	sort.Sort(Sts(sts))
	fmt.Printf("%+v\n", sts)
}
