package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
)

func tweetImage(filename, title string) {

	//login to the anaconda twitter api
	api := anaconda.NewTwitterApiWithCredentials(
		os.Getenv("ACCESS_TOKEN"),
		os.Getenv("ACCESS_SECRET"),
		os.Getenv("API_KEY"),
		os.Getenv("API_SECRET"),
	)

	//read the file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	mediaResponse, err := api.UploadMedia(base64.StdEncoding.EncodeToString(data))
	if err != nil {
		panic(err)
	}

	v := url.Values{}
	v.Set("media_ids", strconv.FormatInt(mediaResponse.MediaID, 10))

	result, err := api.PostTweet(title, v)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Tweeted: %s\n", result.Text)
		fmt.Println("Success! Tweet Complete!")
	}
}
