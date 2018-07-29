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
		no int
		in   input
		want string
	}{
		{1, input{"./testlog/tailtestLF.log", false, 3, 10}, "abcdefghi\nabcdefghij\n\n"},
		{2, input{"./testlog/tailtestLF.log", false, 15, 10}, "a\nab\nabcd\nabcde\nabcdef\nabcdefg\n\nabcdefgh\n\nabcdefghi\nabcdefghij\n\n"},
	}

	for _, test := range tests {
		output := &bytes.Buffer{}
		mytail(test.in.filename, test.in.ignoreBlankLineFlag, test.in.N, output, test.in.bufsize)
		if output.String() != test.want {
			t.Errorf("[%v] test error\n", test.no)
			t.Logf("output[%v]\n", output.String())
			t.Logf("want[%v]\n", test.want)
		}
	}

}
