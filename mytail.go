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
		ignoreBlankLineFlag = flag.Bool("v", false, "ignore blank lines select true or false(default false)")
	)
	flag.Parse()
	fileNames := flag.Args()

	// 引数で指定されたファイルを、1つずつループする
	for _, fileName := range fileNames {
		if len(fileNames) > 1 {
			// 複数のファイルが引数に来た場合に、ファイル名を出力
			fmt.Printf("==> %v <==\n", fileName)
		}
		// ベンチマークの結果、バッファサイズには4000を指定
		mytail(fileName, *ignoreBlankLineFlag, *linesNum, os.Stdout, 4000)
	}
}

func mytail(fileName string, ignoreBlankLineFlag bool, N int, output io.Writer, bufsize int64) {
	// 出力したい行数+1個目のLFより後のテキストを出力するため
	lfcount := N + 1

	// テキスト読み込み処理
	fp, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("file(%v) open error %v\n", fileName, err)
	}
	defer fp.Close()

	// ファイルサイズを取得
	info, err := fp.Stat()
	if err != nil {
		log.Fatalf("get fileinfo error %v\n", err)
	}
	size := info.Size()

	// mytailの結果を保存する変数
	lines := ""

	// 行を出力する処理
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

		// シークした分データを読み込む
		_, err := fp.Read(buf)
		if err != nil {
			log.Fatalf("read buffer error %v\n", err)
		}

		// 空行無視モードの処理
		if ignoreBlankLineFlag {
			buf = bytes.Replace(buf, []byte("\r"), []byte(""), -1)
			for bytes.Count(buf, []byte("\n\n")) != 0 {
				buf = bytes.Replace(buf, []byte("\n\n"), []byte("\n"), -1)
			}
		}

		// 読み込んだバイトの中のLFの数をカウント
		bufLfcount := bytes.Count(buf, []byte("\n"))
		if lfcount <= bufLfcount {
			for i := len(buf) - 1; i >= 0; i-- {
				if buf[i] == 10 {
					lfcount--
					if lfcount == 0 {
						// 結果を出力
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
				// テキストファイルの行数より、出力したい行数が大きい場合は、テキストファイルの内容を全て出力する
				fmt.Fprint(output, lines)
				break
			}
		}
		size -= bufsize
	}
}
