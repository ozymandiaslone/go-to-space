package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Scraping...")
	filetype := scrape()

	title, err := os.ReadFile("title.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Tweeting...")
	// upload media
	tweetImage("upload"+filetype, string(title))
}
