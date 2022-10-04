package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	urls := []string{
		"http://google.com",
		"http://golang.org",
		"http://fb.com",
	}

	c := make(chan string)

	for _, url := range urls {
		time.Sleep(2 * time.Second)
		go checkStatus(url, c)
	}

	for link := range c {

		go checkStatus(link, c)

	}

}

func checkStatus(url string, c chan string) {

	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "  i think link is down")
		c <- url
	}

	fmt.Println(url, " Yep its up")
	c <- url

}
