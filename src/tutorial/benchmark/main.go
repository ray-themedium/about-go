package main

import "fmt"

func fibo1(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	return fibo1(n-1) + fibo1(n-2)
}

func fibo2(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	one := 1
	two := 0
	rst := 0
	for i := 2; i <= n; i++ {
		rst = one + two
		two = one
		one = rst
	}
	return rst
}

func main() {
	fmt.Println(fibo1(13))
	fmt.Println(fibo2(13))
}
