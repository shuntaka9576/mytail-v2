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
		bufsize             int64
	}
	var tests = []struct {
		in   input
		want string
	}{
		{input{"./testlog/tailtestLF.log", false, 3, 10}, "abcdefghi\nabcdefghij\n\n"},
		{input{"./testlog/tailtestLF.log", false, 15, 10}, "a\nab\nabcd\nabcde\nabcdef\nabcdefg\n\nabcdefgh\n\nabcdefghi\nabcdefghij\n\n"},
		{input{"./testlog/tailtestLF.log", true, 3, 10}, "abcdefgh\nabcdefghi\nabcdefghij\n"},
		{input{"./testlog/tailtestLF.log", true, 4, 10}, "abcdefg\nabcdefgh\nabcdefghi\nabcdefghij\n"},
		{input{"./testlog/tailtestLF.log", true, 15, 10}, "a\nab\nabcd\nabcde\nabcdef\nabcdefg\nabcdefgh\nabcdefghi\nabcdefghij\n"},
		{input{"./testlog/tailtestLF.log", false, 3, 1024}, "abcdefghi\nabcdefghij\n\n"},
		{input{"./testlog/tailtestLF.log", false, 15, 1024}, "a\nab\nabcd\nabcde\nabcdef\nabcdefg\n\nabcdefgh\n\nabcdefghi\nabcdefghij\n\n"},
		{input{"./testlog/tailtestLF.log", true, 3, 1024}, "abcdefgh\nabcdefghi\nabcdefghij\n"},
		{input{"./testlog/tailtestLF.log", true, 4, 1024}, "abcdefg\nabcdefgh\nabcdefghi\nabcdefghij\n"},
	}

	for no, test := range tests {
		output := &bytes.Buffer{}
		mytail(test.in.filename, test.in.ignoreBlankLineFlag, test.in.N, output, test.in.bufsize)
		if output.String() != test.want {
			t.Errorf("[%v] test error\n", no)
			t.Logf("output[%v]\n", output.String())
			t.Logf("want[%v]\n", test.want)
		}
	}

}
