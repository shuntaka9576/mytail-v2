package main

import (
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
	fileNames := flag.Args()

	for _, fileName := range fileNames {
		if len(fileNames) > 1 {
			fmt.Printf("==> %v <==\n", fileName)
		}
		// ベンチマークの結果、バッファサイズには4000を指定
		mytail(fileName, *ignoreBlankLineFlag, *linesNum, os.Stdout, 4000)
	}
}

func mytail(fileName string, ignoreBlankLineFlag bool, N int, output io.Writer, bufsize int64) {
	lfcount := N + 1
	fp, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("file(%v) open error %v\n", fileName, err)
	}
	defer fp.Close()

	info, err := fp.Stat()
	if err != nil {
		log.Fatalf("get fileinfo error %v\n", err)
	}

	size := info.Size()

	lines := ""

	for lfcount >= 0 {
		var buf []byte
		if size-bufsize > 0 {
			buf = make([]byte, bufsize)
			fp.Seek(size-bufsize, 0)
		} else {
			buf = make([]byte, size)
			fp.Seek(0, 0)
			size = 0
		}

		_, err := fp.Read(buf)
		if err != nil {
			log.Fatalf("read buffer error %v\n", err)
		}

		// 空行無視モード
		if ignoreBlankLineFlag {
			buf = bytes.Replace(buf, []byte("\r"), []byte(""), -1)
			for bytes.Count(buf, []byte("\n\n")) != 0 {
				buf = bytes.Replace(buf, []byte("\n\n"), []byte("\n"), -1)
			}
		}

		bufLfcount := bytes.Count(buf, []byte("\n"))
		if lfcount <= bufLfcount {
			for i := len(buf) - 1; i >= 0; i-- {
				if buf[i] == 10 {
					lfcount--
					if lfcount == 0 {
						fmt.Fprint(output, string(buf[i+1:])+lines)
						break
					}
				}
			}
			break
		} else {
			lfcount -= bufLfcount
			lines = string(buf) + lines
			if size == 0 {
				fmt.Fprint(output, lines)
				break
			}
		}
		size -= bufsize
	}
}
