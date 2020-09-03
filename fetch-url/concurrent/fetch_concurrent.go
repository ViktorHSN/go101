package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	start := time.Now()
	ch := make(chan string) // cria canal

	file, err := os.OpenFile("./fetch-all-output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	check(err)
	defer file.Close()

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // inicia uma GO ROUTINE
	}
	for range os.Args[1:] {
		out := <-ch // recebe do canal
		b, err := file.WriteString(out + "\n")
		check(err)
		fmt.Printf("wrote %d bytes\n", b)

	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}
