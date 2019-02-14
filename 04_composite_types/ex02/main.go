package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {

	var hash_type = flag.String("h", "", "hash type.")
	flag.Parse()

	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	in := s.Text()

	switch *hash_type {
	case "sha384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(in)))
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(in)))
	default:
		fmt.Printf("%x\n", sha256.Sum256([]byte(in)))
	}
}
