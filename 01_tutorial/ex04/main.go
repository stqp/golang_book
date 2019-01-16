package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	file_names := make(map[string]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, file_names)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %vn", err)
				continue
			}
			countLines(f, counts, file_names)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, file_names[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, file_names map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		file_names[input.Text()] += f.Name() + ","
	}
}
