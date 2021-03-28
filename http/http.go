package http

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/cheggaaa/pb/v3"
	"github.com/sirupsen/logrus"
)

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

// NewClient creates a new http client based on given configuration.
func NewClient(
	method string,
	url string,
	header map[string][]string,
	query map[string]string,
	body string,
	timeout time.Duration,
) *Client {
	if method != GET && method != POST && method != PATCH && method != PUT && method != DELETE {
		panic("method is not recognized. Use GET, POST, PUT, PATCH and DELETE instead")
	}

	// check if there are query parameters and then append them into the url
	if len(query) > 0 {
		url += "?"
		for k, v := range query {
			url += k + "=" + v + "&"
		}

		url = strings.TrimSuffix(url, "&")
	}

	return &Client{
		Method:  method,
		URL:     url,
		Header:  header,
		Query:   query,
		Body:    body,
		Timeout: timeout,
	}
}

func (c *Client) Do() error {
	client := &http.Client{
		Transport: &http.Transport{
			ResponseHeaderTimeout: time.Second * c.Timeout,
		},
	}

	logrus.Infof("sending request into %s", c.URL)

	req, err := http.NewRequest(c.Method, c.URL, bytes.NewBuffer([]byte(c.Body)))
	if err != nil {
		return fmt.Errorf("request creation failed: %w", err)
	}

	req.Header = c.Header

	resp, err := client.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return fmt.Errorf("request timeout: %w", err)
		}

		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	logrus.Infof("Method is: %s", c.Method)
	logrus.Infof("Response status: %s", resp.Status)

	logrus.Info("headers are:")

	for k, v := range resp.Header {
		logrus.Infof("%s = %s", k, v)
	}

	rd := resp.Body

	size, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err == nil && size > 0 {
		// start a new progress bar
		bar := pb.New(size)
		bar.Set(pb.Bytes, true)
		bar.Start()

		// creates a proxy reader for showing the progress bar
		rd = bar.NewProxyReader(resp.Body)
	}

	body, err := ioutil.ReadAll(rd)
	if err != nil {
		return fmt.Errorf("reading body failed: %w", err)
	}

	logrus.Info(string(body))

	return nil
}
