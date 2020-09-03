package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	dupToFile := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, dupToFile)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, dupToFile)
			f.Close()
		}
	}

	printCounts(counts, dupToFile)
}

func countLines(f *os.File, counts map[string]int, dupFile map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		appendToMappingIfNotContains(dupFile, input.Text(), f.Name())
		counts[input.Text()]++
	}

}

func appendToMappingIfNotContains(dupFile map[string]string, key string, fileName string) {
	if !strings.Contains(dupFile[key], fileName) {
		dupFile[key] += fileName + " "
	}
}

func printCounts(counts map[string]int, dupFile map[string]string) {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, dupFile[line])
		}
	}
}
