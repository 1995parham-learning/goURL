package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Created so that multiple inputs can be accecpted
type ArrayFlag []string

func (a *ArrayFlag) String() string {
	// change this, this is just can example to satisfy the interface
	return "my string representation"
}

func (a *ArrayFlag) Set(value string) error {
	*a = append(*a, strings.TrimSpace(value))
	return nil
}

func (a *ArrayFlag) ToHeaderMap(json bool) (map[string][]string, string) {
	header := make(map[string][]string)
	warning := ""

	for _, f := range *a {
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

	_, ok := header["content-type"]
	if !ok {
		if json {
			header["content-type"] = []string{"application/json"}
		}else {
			header["content-type"] = []string{"x-www-form-urlencoded"}
		}
	}

	return header, warning
}

func (a *ArrayFlag) ToQueryMap() (map[string]string, string) {
	query := make(map[string]string)
	warning := ""

	for _, q := range *a {
		a := strings.Split(q, "&")
		for _, b := range a {
			keyValue := strings.Split(b, "=")
			_, ok := query[keyValue[0]]
			if ok {
				warning += fmt.Sprintf("%s is a repetead header so only the last one is considered", keyValue[0])
			}

			query[keyValue[0]] = keyValue[1]
		}
	}

	return query, warning
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
	Query  map[string]string
	Body   string
}

func NewClient(method string, url string, header map[string][]string, query map[string]string, body string) *Client {
	if method != GET && method != POST && method != PATCH && method != PUT && method != DELETE {
		panic("Method is not recognized. Use GET, POST, PUT, PATCH and DELETE instead")
	}

	// check if there are query parameters
	url += "?"
	for k, v := range query {
		url += k + "=" + v + "&"
	}

	url = strings.TrimSuffix(url, "&")

	return &Client{
		Method: method,
		URL:    url,
		Header: header,
		Query:  query,
		Body:   body,
	}
}

func (c *Client) Do() {
	client := &http.Client{}

	fmt.Println(c.URL)

	req, err := http.NewRequest(c.Method, c.URL, bytes.NewBuffer([]byte(c.Body)))
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
