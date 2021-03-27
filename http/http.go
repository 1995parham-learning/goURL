package http

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/cheggaaa/pb/v3"
)

// Created so that multiple inputs can be accecpted
type ArrayFlag []string

func New(arr []string) *ArrayFlag {
	a := ArrayFlag{}
	a = arr
	return &a
}

func (a *ArrayFlag) ToHeaderMap(format string) (map[string][]string, string) {
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
	if !ok && format != "" {
		header["content-type"] = []string{format}
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
	client := &http.Client{Transport: &http.Transport{
		ResponseHeaderTimeout: time.Second * c.Timeout,
	}}

	fmt.Println(c.URL)

	req, err := http.NewRequest(c.Method, c.URL, bytes.NewBuffer([]byte(c.Body)))
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

	for k, v := range resp.Header {
		fmt.Println(fmt.Sprintf("%s = %s", k, v))
	}

	rd := resp.Body

	size, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	fmt.Println(size)
	fmt.Println("parrrrrrrrrrrrrrrhaaaaaaaaaaaaaammmmmmmmmmmm")
	if err != nil && size > 0 {
		// start new bar
		bar := pb.New(size)
		bar.Set(pb.Bytes, true)
		bar.Start()
		// create proxy reader
		rd = bar.NewProxyReader(resp.Body)
	}

	body, err := ioutil.ReadAll(rd)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Println(sb)
}
