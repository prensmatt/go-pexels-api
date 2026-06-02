package main

import(
	"fmt"
	"net/http"
	"os"
	"log"
	"encoding/json"
	"strconv"
	"time"
	"io/ioutil"
)

const(
	PhotoApi = "https://api.pexels.com/v1"
	VideoApi = "https://api.pexels.com/videos"
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
	Src								PhotoSource	`json:"src"`
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

type SearchResult struct{
	Page 					int32	`json:"page"`
	PerPage 			int32  `json:"per_page"`
	TotalResults	int32		`json:"total_results"`
	NextPage 			string	`json:"next_page"`
	Photos 				[]Photo	`json:"photos"`
}

func NewClient(token string) *Client{
	c := http.Client{}
	return &Client{Token: token, hc: c}
}

func (c *Client) SearchPhotos(query string, perPage, page int)(*SearchResult, error){
	url := fmt.Sprintf(PhotoApi + "/search?query=%s&per_page=%d&page=%d", query, perPage, page)
	resp, err := c.requestDoWithAuth("GET", url)
	if err != nil{
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return nil,err
	}
	var result SearchResult

	err = json.Unmarshal(data, &result)

	return &result, err
}




func main(){
	os.Setenv("PexelsToken","F3hzZ2J0a00nhBk8qucFyRDO1qhL7ixftLqMUacH9dfWNOP0Tq1OTdQM")

	var TOKEN = os.Getenv("PexelsToken")
	var c = NewClient(TOKEN)

	result, err := c.SearchPhotos("waves",15,1)

	if err != nil{
		log.Fatal("Search error: ", err)
	}
	fmt.Println(result)
}