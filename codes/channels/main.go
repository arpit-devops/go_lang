package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	urls := []string{

		"http://google.com",
		"http://facebook.com",
		"http://wikipedia.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range urls {
		go checkStatus(link, c)
	}

	for link := range c {
		time.Sleep(2 * time.Second)
		go checkStatus(link, c)
	}

}

func checkStatus(link string, c chan string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")
		c <- link
		return
	}
	fmt.Println(link, " is up")
	c <- link
}
