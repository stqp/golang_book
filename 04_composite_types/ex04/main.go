package main

import "fmt"

func rotate(s []int) {
	for i := len(s) - 1; i > 0; i-- {
		s[i], s[i-1] = s[i-1], s[i]
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(arr)
	rotate(arr)
	fmt.Println(arr)
}
