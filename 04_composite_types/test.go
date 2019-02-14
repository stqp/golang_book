package main

import "fmt"

func main() {
	type Currency int
	const (
		USD Currency = iota
		EUR
		RMB
		GBP
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}

	fmt.Println(RMB, symbol[RMB])

	a := [2]int{1, 2}
	b := []int{1, 2}

	fmt.Println(a)
	fmt.Println(b)
	//fmt.Println(a == b)

}
