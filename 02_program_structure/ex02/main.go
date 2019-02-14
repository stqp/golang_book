package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var in string
	s := bufio.NewScanner(os.Stdin)
	if len(os.Args) > 1 {
		in = os.Args[1]
	} else {
		s.Scan()
		in = s.Text()
	}
	ini, _ := strconv.Atoi(in)
	fmt.Printf("Input: %d km, equals %d m .\n", ini, ini*1000)
}
