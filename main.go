package main

import (
	"errors"
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request Failed")

func main() {
	results := make(map[string]string)
	channel := make(chan result)
	urls := []string{
		"https://www.google.com/",
		"https://www.naver.com/",
		"https://www.airbnb.com/",
		"https://www.reddit.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
	}
	results["hello"] = "Hello"
	for _, url := range urls {
		go hitURL(url, channel)
	}

	for i := 0; i < len(urls); i++ {
		result := <-channel
		fmt.Println(result.status)
		fmt.Println(result.url)
	}
}

func hitURL(url string, channel chan<- result) {
	fmt.Println("Checking", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	channel <- result{url: url, status: status}

}
