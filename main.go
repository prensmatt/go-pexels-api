package main

import(
	"fmt"
	"net/http"
	"os"
	"log"
	"encoding/json"
	"strconv"
	"time"
	"math/rand"
	"io"
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

type CuratedResult struct{
	Page				int32				`json:"page"`
	PerPage			int32				`json:"per_page"`
	NextPage		string			`json:"next_page"`
	Photos      []Photo			`json:"photos"`
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

	data, err := io.ReadAll(resp.Body)
	if err != nil{
		return nil,err
	}
	var result SearchResult

	err = json.Unmarshal(data, &result)

	return &result, err
}

func(c *Client) CuratedPhotos(perPage, page int)(*CuratedResult, error){
	url := fmt.Sprintf(PhotoApi+"/curated?per_page=%d&page=%d",perPage,page)
	resp, err := c.requestDoWithAuth("GET",url)
	if err != nil{
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil{
		return nil, err
	}
	var result CuratedResult
	err = json.Unmarshal(data, &result)
	if err != nil{
		return nil, err
	}
	return &result, err
}

func(c *Client) requestDoWithAuth(method, url string)(*http.Response, error){
	req, err := http.NewRequest(method, url, nil)
	if err != nil{
		return nil, err
	}
	req.Header.Add("Authorization", c.Token)
	resp, err := c.hc.Do(req)
	if err != nil{
		return nil, err
	}
	times, err := strconv.Atoi(resp.Header.Get("X-Ratelimit-Remaining"))
	if err != nil{
		return resp, nil
	}
	c.RemainingTimes = int32(times)
	return resp, nil
}


func (c *Client) GetPhoto(id int32)(*Photo, error){
	url := fmt.Sprintf(PhotoApi+"/photos/%d",id)
	resp, err := c.requestDoWithAuth("GET", url)
	if err != nil{
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil{
		return nil, err
	}
	var result Photo
	err = json.Unmarshal(data, &result)
	return &result, err
}

func(c *Client) GetRandomPhoto()(*Photo, error){
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(1001)
	result, err := c.CuratedPhotos(1, randNum)
	if err == nil && len(result.Photos)==1{
		return &result.Photos[0], nil
	}
	return nil, err
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