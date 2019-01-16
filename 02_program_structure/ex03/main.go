package main

import (
	"fmt"
)

func main() {

	var s = "a世界"
	n := 0
	for range s {
		n++
	}
	fmt.Println(n)
	fmt.Println(len(s))

}
