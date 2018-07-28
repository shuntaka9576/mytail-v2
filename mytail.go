package main

import (
	//"fmt"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var (
		linesNum            = flag.Int("n", 10, "output the last N lines")
		ignoreBlankLineFlag = flag.Bool("v", false, "output the last N lines")
	)
	flag.Parse()
	fileName := flag.Args()

	mytail(fileName, *ignoreBlankLineFlag, *linesNum, os.Stdout, 1024)
}

func mytail(fileNames []string, ignoreBlankLineFlag bool, N int, output io.Writer, bufsize int64) {
	for _, fileName := range fileNames {
		if len(fileNames) > 1 {
			fmt.Printf("==> %v <==\n", fileName)
		}
		fp, err := os.Open(fileName)
		if err != nil {
			log.Fatalf("file(%v) open error %v\n", fileName, err)
		}

		info, err := fp.Stat()
		if err != nil {
			log.Fatalf("get fileinfo error %v\n", err)
		}

		size := info.Size()

		lines := ""

		for N != 0 {
			buf := make([]byte, bufsize)
			fp.Seek(size-bufsize, 0)
			_, _ = fp.Read(buf)

			if ignoreBlankLineFlag {
				fmt.Println(bytes.Count(buf, []byte("\n\n")))
				for bytes.Count(buf, []byte("\n\n")) != 0 {
					buf = bytes.Replace(buf, []byte("\n\n"), []byte("\n"), -1)
				}
			}
			num := bytes.Count(buf, []byte("\n"))

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
}
