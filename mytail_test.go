package main

import (
	"testing"
)

func BenchmarkTailBuf1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mytail(1)
	}
}

func BenchmarkTailBuf100(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mytail(100)
	}
}

func BenchmarkTailBuf1024(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mytail(1024)
	}
}

func BenchmarkTailBuf2048(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mytail(2048)
	}
}

func BenchmarkTailBuf10240(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mytail(10240)
	}
}
