package main

import (
	"fmt"
	"unicode/utf8"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	words := utf8.RuneCount(p)
	*w += WordCounter(words)
	return words, nil
}

type LineCounter int

func (l *LineCounter) Write(p string) (int, error) {
	lines := 1
	for i := 0; i < len(p); i++ {
		if p[i] == '\n' {
			lines++
		}
	}
	*l += LineCounter(lines)
	return lines, nil
}

func main() {

	str := `first lineあ
		second line
		third line`

	var c ByteCounter
	var w WordCounter
	var l LineCounter

	c.Write([]byte(str))
	w.Write([]byte(str))
	l.Write(str)

	fmt.Println(c)
	fmt.Println(w)
	fmt.Println(l)
}
