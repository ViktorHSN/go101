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
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		fmt.Printf("Requesting from %s\n", url)
		resp, err := http.Get(url)
		handleErr(err, "fecth: %v\n")

		fmt.Printf("response statuscode: %v\n", resp.StatusCode)
		fmt.Printf("response status message: %s\n", resp.Status)

		w, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		handleErr(err, "fetch: reading %s: %v\n")
		fmt.Printf("\n%d", w)

	}
}

func handleErr(err error, fmtMsg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, fmtMsg, err)
		os.Exit(1)
	}
}
