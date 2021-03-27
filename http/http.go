package http

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// Created so that multiple inputs can be accecpted
type ArrayFlag []string

func New(arr []string) *ArrayFlag {
	a := ArrayFlag{}
	a = arr
	return &a
}

func (a *ArrayFlag) ToHeaderMap(json bool) (map[string][]string, string) {
	header := make(map[string][]string)
	warning := ""

	for _, f := range *a {
			keyValue := strings.Split(f, ":")
			_, ok := header[strings.ToLower(keyValue[0])]
			if ok {
				warning += fmt.Sprintf("%s is a repetead header so only the last one is considered", keyValue[0])
			}

			header[strings.ToLower(keyValue[0])] = []string{keyValue[1]}
	}

	_, ok := header["content-type"]
	if !ok {
		if json {
			header["content-type"] = []string{"application/json"}
		} else {
			header["content-type"] = []string{"application/x-www-form-urlencoded"}
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
			_, ok := query[strings.ToLower(keyValue[0])]
			if ok {
				warning += fmt.Sprintf("%s is a repetead query parameter so only the last one is considered", keyValue[0])
			}

			query[strings.ToLower(keyValue[0])] = keyValue[1]
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
	Method  string
	URL     string
	Header  map[string][]string
	Query   map[string]string
	Body    string
	Timeout time.Duration
}

func NewClient(method string, url string, header map[string][]string, query map[string]string, body string, timeout time.Duration) *Client {
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
		Method:  method,
		URL:     url,
		Header:  header,
		Query:   query,
		Body:    body,
		Timeout: timeout,
	}
}

func (c *Client) Do() {
	client := &http.Client{}

	fmt.Println(c.URL)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * time.Duration(c.Timeout))
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, c.Method, c.URL, bytes.NewBuffer([]byte(c.Body)))
	if err != nil {
		panic(err)
	}

	req.Header = c.Header

	resp, err := client.Do(req)
	if err != nil {
		if err == context.DeadlineExceeded {
			fmt.Println(err)
			return
		}
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Method is:", c.Method)
	fmt.Println("Response status:", resp.Status)
	fmt.Println("headers are:")

	for k, v := range resp.Header{
		fmt.Println(fmt.Sprintf("%s = %s", k, v))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Println(sb)
}
