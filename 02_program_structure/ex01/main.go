package main

import "fmt"

func main() {
	var p = f()
	fmt.Println(*p)
	*p++
	fmt.Println(*p)
}

func f() *int {
	var x = 10
	return &x
}
