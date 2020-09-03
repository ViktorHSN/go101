package main

import (
	"fmt"
	"os"
	string "strings"
	"time"
)

func main() {

	start := time.Now()
	fmt.Println("Starting echo join")
	echoJoin()
	seconds := time.Since(start).Seconds()
	fmt.Printf("Ending echo join, done in: %f", seconds)
	fmt.Println()

	start = time.Now()
	fmt.Println("Starting echo for")
	echoFor()
	seconds = time.Since(start).Seconds()
	fmt.Printf("Ending echo join, done in: %f", seconds)

}

func echoJoin() {
	fmt.Println(string.Join(os.Args[0:], " "))
}

func echoFor() {
	for _, arg := range os.Args[0:] {
		fmt.Println(arg)
	}

}
