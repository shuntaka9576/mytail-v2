package main

import (
	//"fmt"
	"os"
	"strings"
)

const bufsize = 1024

func main() {
	tail(bufsize)
}

func tail(bufsize int64) {
	fp, _ := os.Open("testlog.log")
	info, _ := fp.Stat()
	size := info.Size()

	n := 5
	lines := ""

	for n != 0 {
		buf := make([]byte, bufsize)
		fp.Seek(size-bufsize, 0)
		_, _ = fp.Read(buf)

		num := strings.Count(string(buf), "\n")

		if n <= num {
			for i := len(buf) - 1; i >= 0; i-- {
				if buf[i] == 10 {
					n--
					if n == 0 {
						//fmt.Print(string(buf[i+1:]) + lines)
						break
					}
				}
			}
		} else {
			n -= num
			lines = string(buf) + lines
		}
		size -= bufsize
	}
}
