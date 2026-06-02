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

type Photo struct{
	ID								int32				`json:"id"`
	Width							int32				`json:"width"`
	Height						int32				`json:"height"`
	URL								string			`json:"url"`
	Photographer			string			`json:"photographer"`
	PhotographerURL		string			`json:"photographer_url"`
	Src								PhotoSource	`json:"json:"src"`
}

type PhotoSource struct{
	Original		string		`json:"original"`
	Large				string		`json:"large"`
	Large2x			string		`json:"large2x"`
	Medium			string		`json:"medium"`
	Small				string		`json:"small"`
	Portrait		string		`json:"portrait"`
	Square			string		`json:"square"`
	Landscape		string		`json:"landscape"`
	Tiny				string		`json:"tiny"`
}

func NewClient(token string) *Client{
	c := http.Client()
	return &Client{Token: token, hc: c}
}

type SearchResult struct{
	Page 					int32	`json:"page"`
	PerPage 			int32  `json:"per_page"`
	TotalResults	int32		`json:"total_results"`
	NextPage 			string	`json:"next_page"`
	Photos 				[]Photo	`json:"photos"`
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