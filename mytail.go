package main

import (
	//"fmt"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"log"
)

var (
	linesNum = flag.Int("n", 10, "output the last N lines")
)

func main() {
	flag.Parse()

	fileName := ""
	switch len(os.Args) {
	case 2:
		fileName = os.Args[1]
	case 4:
		fileName = os.Args[3]
	default:
		log.Fatal("args error")
	}

	mytail(fileName, *linesNum, os.Stdout, 1024)
}

func mytail(fileName string, N int, output io.Writer, bufsize int64) {
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
