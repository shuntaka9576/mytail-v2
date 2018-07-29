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
			}
		}
	}
}
