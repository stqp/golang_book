package main

import (
	"fmt"
	"io"
	"os"
)

type WriteCounter struct {
	count  int64
	writer io.Writer
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	wc.count += int64(len(p))
	fmt.Println(wc.count)
	return wc.writer.Write(p)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var wc WriteCounter
	wc.writer = w
	wc.count = 0
	return &wc, &wc.count
}

func main() {
	wc, _ := CountingWriter(os.Stdout)
	_, _ = wc.Write([]byte("hello"))
}
