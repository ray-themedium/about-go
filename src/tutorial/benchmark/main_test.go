package main

import (
	"testing"
)

func BenchmarkFibo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibo1(20)
	}
}

func BenchmarkFibo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibo2(20)
	}
}
