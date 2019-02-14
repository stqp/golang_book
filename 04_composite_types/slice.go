package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	var sli []int

	fmt.Printf("%T\n", sli)
	fmt.Println(arr)
}

func change(arr [5]int) {
	arr[0] = 999
	fmt.Println(arr)
}
