package main

import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		//fmt.Println(i, pc[i])
	}
}

func main() {
	fmt.Println(PopCount(999999))
}

func PopCount(x uint64) int {
	count := uint64(0)
	for i := uint(0); i < 64; i++ {
		count += uint64(x>>i) & uint64(1)
	}
	return int(count)
}
