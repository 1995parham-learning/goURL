package http

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	GET    = "GET"
	POST   = "POST"
	PATCH  = "PATCH"
	DELETE = "DELETE"
	PUT    = "PUT"
)

type Client struct {
	Method string
	URL    string
	Header map[string]string
}

func NewClient(method string, url string, header map[string]string) *Client {
	if method != GET && method != POST && method != PATCH && method != PUT && method != DELETE {
		panic("Method is not recognized. Use GET, POST, PUT, PATCH and DELETE instead")
	}

	return &Client{
		Method: method,
		URL:    url,
		Header: header,
	}
}

func (c *Client) Do() {
	client := &http.Client{}

	req, err := http.NewRequest(c.Method, c.URL, nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Println(sb)
}
