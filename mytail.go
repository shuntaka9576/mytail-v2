package main

import (
	//"fmt"
	"os"
	"strings"
	"io"
	"fmt"
	"flag"
)

var (
	linesNum = flag.Int("n", 10, "output the last N lines")
)

func main() {
	flag.Parse()
	mytail(*linesNum, os.Stdout, 1024)
}

func mytail(N int, output io.Writer, bufsize int64) {
	fp, _ := os.Open("testlog.log")
	info, _ := fp.Stat()
	size := info.Size()

	lines := ""

	for N != 0 {
		buf := make([]byte, bufsize)
		fp.Seek(size-bufsize, 0)
		_, _ = fp.Read(buf)

		num := strings.Count(string(buf), "\n")

		if N <= num {
			for i := len(buf) - 1; i >= 0; i-- {
				if buf[i] == 10 {
					N--
					if N == 0 {
						fmt.Fprint(output, string(buf[i+1:])+lines)
						break
					}
				}
			}
		} else {
			N -= num
			lines = string(buf) + lines
		}
		size -= bufsize
	}
}
