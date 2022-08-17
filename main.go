package main

import (
	"fmt"
	"os"
    "time"
)

func main() {
    for {
        // begin tweet sequence
     fmt.Println("beginning scrape...")
     filetype := scrape()
     title, err := os.Readfile("title.txt")
      if err!= nil {
         panic(err)
     }
        // send tweet
      fmt.Println("sending tweet...")
      tweetImage("upload"+filetype, string(title))
      time.sleep(3600 * time.Second)
     }
}
