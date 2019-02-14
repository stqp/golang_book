package main

import "fmt"

func newInt() *int {
	return new(int)
}

func main() {

	var pc [256]byte

	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		fmt.Println(i, pc[i])
	}

}
