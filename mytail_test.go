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

// 軽いテキストで5行出力するベンチマーク
func benchmarkLightText(b *testing.B, bufsize int64) {
	output := &bytes.Buffer{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mytail("./testlog/tailtestLF.log", false, 5, output, bufsize)
	}
}
func BenchmarkTailBuf1LightText(b *testing.B)     { benchmarkLightText(b, 1) }
func BenchmarkTailBuf100LightText(b *testing.B)   { benchmarkLightText(b, 100) }
func BenchmarkTailBuf1000LightText(b *testing.B)  { benchmarkLightText(b, 1000) }
func BenchmarkTailBuf2000LightText(b *testing.B)  { benchmarkLightText(b, 2000) }
func BenchmarkTailBuf3000LightText(b *testing.B)  { benchmarkLightText(b, 3000) }
func BenchmarkTailBuf4000LightText(b *testing.B)  { benchmarkLightText(b, 4000) }
func BenchmarkTailBuf5000LightText(b *testing.B)  { benchmarkLightText(b, 5000) }
func BenchmarkTailBuf6000LightText(b *testing.B)  { benchmarkLightText(b, 6000) }
func BenchmarkTailBuf7000LightText(b *testing.B)  { benchmarkLightText(b, 7000) }
func BenchmarkTailBuf8000LightText(b *testing.B)  { benchmarkLightText(b, 8000) }
func BenchmarkTailBuf9000LightText(b *testing.B)  { benchmarkLightText(b, 9000) }
func BenchmarkTailBuf10000LightText(b *testing.B) { benchmarkLightText(b, 10000) }

// 1G程度のテキストで5行出力するベンチマーク
func benchmarkHeavyText(b *testing.B, bufsize int64) {
	output := &bytes.Buffer{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mytail("./testlog/heavy.log", false, 5, output, bufsize)
	}
}
func BenchmarkTailBuf1HeavyText(b *testing.B)     { benchmarkHeavyText(b, 1) }
func BenchmarkTailBuf100HeavyText(b *testing.B)   { benchmarkHeavyText(b, 100) }
func BenchmarkTailBuf1000HeavyText(b *testing.B)  { benchmarkHeavyText(b, 1000) }
func BenchmarkTailBuf2000HeavyText(b *testing.B)  { benchmarkHeavyText(b, 2000) }
func BenchmarkTailBuf3000HeavyText(b *testing.B)  { benchmarkHeavyText(b, 3000) }
func BenchmarkTailBuf4000HeavyText(b *testing.B)  { benchmarkHeavyText(b, 4000) }
func BenchmarkTailBuf5000HeavyText(b *testing.B)  { benchmarkHeavyText(b, 5000) }
func BenchmarkTailBuf6000HeavyText(b *testing.B)  { benchmarkHeavyText(b, 6000) }
func BenchmarkTailBuf7000HeavyText(b *testing.B)  { benchmarkHeavyText(b, 7000) }
func BenchmarkTailBuf8000HeavyText(b *testing.B)  { benchmarkHeavyText(b, 8000) }
func BenchmarkTailBuf9000HeavyText(b *testing.B)  { benchmarkHeavyText(b, 9000) }
func BenchmarkTailBuf10000HeavyText(b *testing.B) { benchmarkHeavyText(b, 10000) }

// 1G程度のテキストで100行出力するベンチマーク
func benchmarkHeavyTextManyLines(b *testing.B, bufsize int64) {
	output := &bytes.Buffer{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mytail("./testlog/heavy.log", false, 100, output, bufsize)
	}
}
func BenchmarkTailBuf1HeavyTextManyLines(b *testing.B)     { benchmarkHeavyTextManyLines(b, 1) }
func BenchmarkTailBuf100HeavyTextManyLines(b *testing.B)   { benchmarkHeavyTextManyLines(b, 100) }
func BenchmarkTailBuf1000HeavyTextManyLines(b *testing.B)  { benchmarkHeavyTextManyLines(b, 1000) }
func BenchmarkTailBuf2000HeavyTextManyLines(b *testing.B)  { benchmarkHeavyTextManyLines(b, 2000) }
func BenchmarkTailBuf3000HeavyTextManyLines(b *testing.B)  { benchmarkHeavyTextManyLines(b, 3000) }
func BenchmarkTailBuf4000HeavyTextManyLines(b *testing.B)  { benchmarkHeavyTextManyLines(b, 4000) }
func BenchmarkTailBuf5000HeavyTextManyLines(b *testing.B)  { benchmarkHeavyTextManyLines(b, 5000) }
func BenchmarkTailBuf6000HeavyTextManyLines(b *testing.B)  { benchmarkHeavyTextManyLines(b, 6000) }
func BenchmarkTailBuf7000HeavyTextManyLines(b *testing.B)  { benchmarkHeavyTextManyLines(b, 7000) }
func BenchmarkTailBuf8000HeavyTextManyLines(b *testing.B)  { benchmarkHeavyTextManyLines(b, 8000) }
func BenchmarkTailBuf9000HeavyTextManyLines(b *testing.B)  { benchmarkHeavyTextManyLines(b, 9000) }
func BenchmarkTailBuf10000HeavyTextManyLines(b *testing.B) { benchmarkHeavyTextManyLines(b, 10000) }
