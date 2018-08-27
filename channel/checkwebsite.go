package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkWebsites() {
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.stackoverflow.com",
		"https://www.amazon.com",
		"https://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go check(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			check(link, c)
		}(l)
	}
}

func check(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		c <- url
		return
	}

	fmt.Println(url + " is up!")
	c <- url
}
