package main

import (
	"fmt"
	"sort"
)

type a struct {
	b string
	c int
	d int
}

func main() {
	sl := make([]a, 0)
	sl = append(sl, a{"1", 100, 200})
	sl = append(sl, a{"2", 50, 50})
	sl = append(sl, a{"3", 20, 10})
	sl = append(sl, a{"4", 106, 2040})
	sl = append(sl, a{"5", 3900, 20})

	sort.Slice(sl, func(i, j int) bool {
		return sl[i].c < sl[j].c
	})
	fmt.Println(sl)

	sort.Slice(sl, func(i, j int) bool {
		return sl[i].c > sl[j].c
	})

	fmt.Println(sl)
}
