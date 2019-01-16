package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	start := time.Now()

	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello"
	}

	secs := time.Since(start).Seconds()

	fmt.Printf("%.3fs\n", secs)

	start = time.Now()

	arr := []string{}
	for i := 0; i < 100000; i++ {
		arr = append(arr, "hello")
	}
	str = strings.Join(arr, "")

	secs = time.Since(start).Seconds()

	fmt.Printf("%.3fs\n", secs)

}
