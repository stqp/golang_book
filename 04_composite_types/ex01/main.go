package main

import (
	"crypto/sha256"
	"fmt"
)

func PopCount(b [32]uint8) int {
	count := uint8(0)
	for i := 0; i < 32; i++ {
		count +=
			((b[i] >> 0) & uint8(1)) +
				((b[i] >> 1) & uint8(1)) +
				((b[i] >> 2) & uint8(1)) +
				((b[i] >> 3) & uint8(1)) +
				((b[i] >> 4) & uint8(1)) +
				((b[i] >> 5) & uint8(1)) +
				((b[i] >> 6) & uint8(1)) +
				((b[i] >> 7) & uint8(1))

	}
	return int(count)
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	c3 := [32]uint8{255}

	fmt.Printf("%d\n", PopCount(c1))
	fmt.Printf("%d\n", PopCount(c2))
	fmt.Printf("%d\n", PopCount(c3))
}
