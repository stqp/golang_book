package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for idx, val := range os.Args[1:] {
		fmt.Println("index: " + strconv.Itoa(idx) + " value: " + val)
	}

}
