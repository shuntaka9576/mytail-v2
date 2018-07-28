package main

import (
	"bytes"
	"os"
	"testing"
)

/*
	test
*/

func TestMytailStandardoutPut3(t *testing.T) {
	expect := `		size -= bufsize
	}
}
`
	output := new(bytes.Buffer)
	mytail("testlog.log", false, 3, output, 10)
	if output.String() != expect {
		t.Errorf("Error %v\n\n%v", expect, output.String())
	}
}

func TestMytailStandardOutput10(t *testing.T) {
	mytail("testlog.log", false, 10, os.Stdout, 10)
}

func TestMytailBlackLine(t *testing.T) {
	mytail("./testlog/blankline.log", true, 3, os.Stdout, 10)
}

// 3行しかないテキストに対して6行要求する
func TestMytailBlac3Line(t *testing.T) {
	expect := `aaa
aaaaa
aaaaaaaa
`
	output := new(bytes.Buffer)
	mytail("./testlog/3lines.log", true, 6, output, 10)
	if output.String() != expect {
		t.Errorf("Error [%v]\n\n[%v]", expect, output.String())
	}
}

/*
	benchmark
*/
const benchN = 100

func BenchmarkTailBuf1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mytail("testlog.log", false, benchN, &bytes.Buffer{}, 1)
	}
}

func BenchmarkTailBufbench10(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mytail("testlog.log", false, benchN, &bytes.Buffer{}, 10)
	}
}

func BenchmarkTailBuf1024(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mytail("testlog.log", false, benchN, &bytes.Buffer{}, 1024)
	}
}

func BenchmarkTailBuf2048(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mytail("testlog.log", false, benchN, &bytes.Buffer{}, 2048)
	}
}

func BenchmarkTailBuf4000(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mytail("testlog.log", false, benchN, &bytes.Buffer{}, 4000)
	}
}

func BenchmarkTailBuf10240(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mytail("testlog.log", false, benchN, &bytes.Buffer{}, 10240)
	}
}
