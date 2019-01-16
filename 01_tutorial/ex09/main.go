package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
			fmt.Println("You are missing 'https://' or 'http://' ")
			os.Exit(1)
		}

		resp, err := http.Get(url)
		defer resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %vn", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s", url, err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, resp.Status)

	}
}
