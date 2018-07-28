package main

import (
	"testing"
)

func BenchmarkTailBuf1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tail(1)
	}
}

func BenchmarkTailBuf100(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tail(100)
	}
}

func BenchmarkTailBuf1024(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tail(1024)
	}
}
