package test

import (
	"testing"
)

func Add(a, b int) int {
	return a + b
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ { // b.N，测试循环次数
		Add(4, 5)
	}
}
