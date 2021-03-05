package http

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Created so that multiple inputs can be accecpted
type HeaderFlag []string

func (h *HeaderFlag) String() string {
	// change this, this is just can example to satisfy the interface
	return "my string representation"
}

func (h *HeaderFlag) Set(value string) error {
	*h = append(*h, strings.TrimSpace(value))
	return nil
}

func (h *HeaderFlag) ToMap() (map[string][]string, string) {
	header := make(map[string][]string)
	warning := ""

	for _, f := range *h {
		a := strings.Split(f, ",")
		for _, b := range a {
			keyValue := strings.Split(b, ":")
			_, ok := header[keyValue[0]]
			if ok {
				warning += fmt.Sprintf("%s is a repetead header so only the last one is considered", keyValue[0])
			}

			header[keyValue[0]] = []string{keyValue[1]}
		}
	}

	return header, warning
}

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
	Header map[string][]string
}

func NewClient(method string, url string, header map[string][]string) *Client {
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

	req.Header = c.Header

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
