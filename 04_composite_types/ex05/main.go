package main

import "fmt"

func nondup(s []string) []string {
	c := 0
	for i := 0; i < len(s); i++ {
		dup := false
		for j := 0; j < i; j++ {
			if s[j] == s[i] {
				dup = true
			}
		}
		if dup == false {
			s[c] = s[i]
			c++
		}
	}
	return s[:c]
}

func main() {
	arr := []string{"a", "b", "a", "c", "b", "a", "d", "e"}
	fmt.Println(nondup(arr))
}
