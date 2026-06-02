package main

import(
	"fmt"
	"net/http"
	"os"
	"encoding/json"
	"strconv"
	"time"
)

type Client struct{
	Token string
	hc http.Client
	RemainingTimes int32
}

func NewClient(token string) *Client{
	c := http.Client()
	return &Client{Token: token, hc: c}
}

const(
	PhotoApi = "https://api.pexels.com/v1"
	VideoApi = "https://api.pexels.com/videos"
)

func main(){
	os.Setenv("PexelsToken","F3hzZ2J0a00nhBk8qucFyRDO1qhL7ixftLqMUacH9dfWNOP0Tq1OTdQM")

	var TOKEN := os.Getenv("PexelsToken")
	var c = NewClient(TOKEN)

	result, err := c.SearchPhotos("waves")

	if err != nil{
		fmt.Errorf("Search error: %v", err)
	}
	fmt.Println(result)
}