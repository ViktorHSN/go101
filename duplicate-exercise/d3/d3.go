package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//This version store all the file in memory

func main() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		//Read File stores all the file in memory and read the data all at once as bytes[]
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
