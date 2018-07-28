package main

import (
	"testing"
	"bytes"
)

const benchN = 100

func BenchmarkTailBuf1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mytail("testlog.log", benchN, &bytes.Buffer{}, 1)
	}
}

func BenchmarkTailBufbench10(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mytail("testlog.log", benchN, &bytes.Buffer{}, 10)
	}
}

func BenchmarkTailBuf1024(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mytail("testlog.log", benchN, &bytes.Buffer{}, 1024)
	}
}

func BenchmarkTailBuf2048(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mytail("testlog.log", benchN, &bytes.Buffer{}, 2048)
	}
}

func BenchmarkTailBuf4000(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mytail("testlog.log", benchN, &bytes.Buffer{}, 4000)
	}
}

func BenchmarkTailBuf10240(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mytail("testlog.log", benchN, &bytes.Buffer{}, 10240)
	}
}
