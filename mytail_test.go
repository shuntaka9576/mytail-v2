package main

import (
	"bytes"
	"testing"
)

func TestMyTail(t *testing.T) {
	type input struct {
		filename            string
		ignoreBlankLineFlag bool
		N                   int
	}
	var tests = []struct {
		in   input
		want string
	}{
		{input{"./testlog/tailtestLF.log", false, 3}, "abcdefghi\nabcdefghij\n\n"},
		{input{"./testlog/tailtestLF.log", false, 15}, "a\nab\nabcd\nabcde\nabcdef\nabcdefg\n\nabcdefgh\n\nabcdefghi\nabcdefghij\n\n"},
		{input{"./testlog/tailtestLF.log", true, 3}, "abcdefgh\nabcdefghi\nabcdefghij\n"},
		{input{"./testlog/tailtestLF.log", true, 4}, "abcdefg\nabcdefgh\nabcdefghi\nabcdefghij\n"},
		{input{"./testlog/tailtestLF.log", true, 15}, "a\nab\nabcd\nabcde\nabcdef\nabcdefg\nabcdefgh\nabcdefghi\nabcdefghij\n"},
		{input{"./testlog/tailtestCRLF.log", false, 3}, "abcdefghi\r\nabcdefghij\r\n\r\n"},
		{input{"./testlog/tailtestCRLF.log", false, 15}, "a\r\nab\r\nabcd\r\nabcde\r\nabcdef\r\nabcdefg\r\n\r\nabcdefgh\r\n\r\nabcdefghi\r\nabcdefghij\r\n\r\n"},
		{input{"./testlog/tailtestCRLF.log", true, 3}, "abcdefgh\nabcdefghi\nabcdefghij\n"},
		{input{"./testlog/tailtestCRLF.log", true, 4}, "abcdefg\nabcdefgh\nabcdefghi\nabcdefghij\n"},
		{input{"./testlog/tailtestCRLF.log", true, 15}, "a\nab\nabcd\nabcde\nabcdef\nabcdefg\nabcdefgh\nabcdefghi\nabcdefghij\n"},
	}

	bufsizePattern := []int64{10, 100, 1024, 2048, 10240}
	for no, test := range tests {
		for _, bufsize := range bufsizePattern {
			output := &bytes.Buffer{}
			mytail(test.in.filename, test.in.ignoreBlankLineFlag, test.in.N, output, bufsize)
			if output.String() != test.want {
				t.Errorf("no[%v] bufsize[%v] test error\n", no, bufsize)
				t.Logf("output[%v]\n", output.String())
				t.Logf("want[%v]\n", test.want)
			} else {
				t.Logf("Pass!!:no[%v] bufsize[%v]\n", no, bufsize)
			}
		}
	}
}

func benchmarkLightText(b *testing.B, bufsize int64) {
	output := &bytes.Buffer{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mytail("./testlog/tailtestLF.log", false, 5, output, bufsize)
	}
}
func BenchmarkTailBuf1LightText(b *testing.B)     { benchmarkLightText(b, 1) }
func BenchmarkTailBuf100LightText(b *testing.B)   { benchmarkLightText(b, 100) }
func BenchmarkTailBuf1024LightText(b *testing.B)  { benchmarkLightText(b, 1024) }
func BenchmarkTailBuf2048LightText(b *testing.B)  { benchmarkLightText(b, 2048) }
func BenchmarkTailBuf10240LightText(b *testing.B) { benchmarkLightText(b, 10240) }

func benchmarkHeavyText(b *testing.B, bufsize int64) {
	output := &bytes.Buffer{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mytail("./testlog/heavy.log", false, 5, output, bufsize)
	}
}
func BenchmarkTailBuf1HeavyText(b *testing.B)     { benchmarkHeavyText(b, 1) }
func BenchmarkTailBuf100HeavyText(b *testing.B)   { benchmarkHeavyText(b, 100) }
func BenchmarkTailBuf1024HeavyText(b *testing.B)  { benchmarkHeavyText(b, 1024) }
func BenchmarkTailBuf2048HeavyText(b *testing.B)  { benchmarkHeavyText(b, 2048) }
func BenchmarkTailBuf10240HeavyText(b *testing.B) { benchmarkHeavyText(b, 10240) }

// ----
func benchmarkHeavyTextManyLines(b *testing.B, bufsize int64) {
	output := &bytes.Buffer{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mytail("./testlog/heavy.log", false, 100, output, bufsize)
	}
}
func BenchmarkTailBuf1HeavyTextManyLines(b *testing.B)     { benchmarkHeavyTextManyLines(b, 1) }
func BenchmarkTailBuf100HeavyTextManyLines(b *testing.B)   { benchmarkHeavyTextManyLines(b, 100) }
func BenchmarkTailBuf1024HeavyTextManyLines(b *testing.B)  { benchmarkHeavyTextManyLines(b, 1024) }
func BenchmarkTailBuf2048HeavyTextManyLines(b *testing.B)  { benchmarkHeavyTextManyLines(b, 2048) }
func BenchmarkTailBuf10240HeavyTextManyLines(b *testing.B) { benchmarkHeavyTextManyLines(b, 10240) }
