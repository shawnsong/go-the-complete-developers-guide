package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.stackoverflow.com",
		"https://www.amazon.com",
		"https://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkWebsite(link, c)
	}

	for {
		checkWebsite(<-c, c)
	}

}

func checkWebsite(url string, c chan string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		c <- url
		return
	}

	fmt.Println(url + " is up!")
	c <- url
}
