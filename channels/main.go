package main

import (
	"fmt"
	"net/http"
	"time"
)


func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second) // Simulate some processing delay
			checkLink(link, c)
		}(l)
	}
}

// Concurrency vs Parallelism - Multiple threads executing concurrently is not-
// the same as multiple threads executing in parallel.

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")

	c <- link
}