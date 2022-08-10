package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func scrape() string {
	//default connector
	c := colly.NewCollector()
	var filetype string
	var pagelinks [100]string
	var imagelinks [80]string
	count := 0
	//On every a element which has href attribute print the link to the console
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		link := e.Request.AbsoluteURL(e.Attr("href"))
		fmt.Printf("Link found: %v\n", link)
		pagelinks[count] = link
		count++

	})

	//Before request, print out intent to console
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	//Make the request to visit a site
	c.Visit(randomUrl())

	//find & select a random link to visit
	for i := 0; i < 100; i++ {
		if strings.Contains(pagelinks[i], "/ESA_Multimedia/Images/") {
			imagelinks[i] = pagelinks[i]
		}
	}

	fmt.Println("Visitng second url...")
	c.OnHTML("a.dropdown__item", func(b *colly.HTMLElement) {
		link := b.Request.AbsoluteURL(b.Attr("href"))
		fmt.Printf("Detected Image source: %v\n", link)
		fmt.Println("Attemtping Download...")
		filetype = string(link[len(link)-4:])
		fmt.Printf("Detected Filetype:%v\n", filetype)
		downloadFile(link, "upload"+filetype)
		fmt.Println("Success! Download Complete!")
	})

	c.OnHTML("h1.heading", func(c *colly.HTMLElement) {
		writeToFile("title.txt", c.Text)
	})
	c.Visit(selectRandom(imagelinks))

	return filetype

}

// functions
func selectRandom(a [80]string) string {
	min := 0
	max := len(a)
	randomguess := rand.Intn(max-min) + min
	if a[randomguess] != "" {
		return a[randomguess]
	} else {
		return selectRandom(a)
	}
}

// generate a random url to visit
func randomUrl() string {
	randint := rand.Intn(580)
	randint = randint * 50
	new := strconv.Itoa(randint)
	url := "https://www.esa.int/ESA_Multimedia/Search/(offset)/" + new + "/(sortBy)/published?result_type=images&SearchText=%2A"
	return url
}

func downloadFile(URL, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("received non 200 response code")
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func writeToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}
