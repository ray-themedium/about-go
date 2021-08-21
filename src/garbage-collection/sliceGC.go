// 포인터를 저장하는 방식에 따라 GC 성능이 크게 달라진다.
// 다루는 포인터가 많을수록 영향이 두드러진다.

package main

import "runtime"

type data struct {
	i, j int
}

func main() {
	var N = 40000000
	var structure []data
	for i := 0; i < N; i++ {
		value := int(i)
		structure = append(structure, data{value, value})
	}
	runtime.GC()
	_ = structure[0]
}
