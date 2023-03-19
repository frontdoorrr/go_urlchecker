package main

import "fmt"

func main() {
	urls := []string{
		"https://www.google.com/",
		"https://www.naver.com/",
		"https://www.airbnb.com/",
		"https://www.reddit.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
	}
	for _, url := range urls {
		fmt.Println(url)
	}
}
