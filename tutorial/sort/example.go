package main

import (
	"fmt"
	"sort"
)

func main() {
	// SortForInts()
	// SortForStrings()
	// SortReverse1()
	// SortReverse2()
	// SortForStructs()
	SortPlayers()
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

type Player struct {
	Name            string
	Age             int
	Score           int
	PassSuccessRate float64
}

type Players []Player

func (a Players) Len() int      { return len(a) }
func (a Players) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Players) Less(i, j int) bool {
	if a[i].Score > a[j].Score {
		return true
	} else if a[i].Score == a[j].Score { // 점수가 같을 때
		return a[i].Age > a[j].Age // 나이가 많은 사람이 먼저 오도록함
	} else {
		return false
	}
}
func SortPlayers() {
	slice := []Player{
		{Name: "a1", Age: 10, Score: 50, PassSuccessRate: 50.5},
		{Name: "a2", Age: 6, Score: 88, PassSuccessRate: 50.5},
		{Name: "a3", Age: 11, Score: 44, PassSuccessRate: 50.5},
		{Name: "a4", Age: 23, Score: 22, PassSuccessRate: 50.5},
		{Name: "a5", Age: 14, Score: 190, PassSuccessRate: 50.5},
		{Name: "a6", Age: 8, Score: 22, PassSuccessRate: 50.5},
		{Name: "a7", Age: 2, Score: 50, PassSuccessRate: 50.5},
	}
	sort.Sort(Players(slice))
	fmt.Println(slice)
}
