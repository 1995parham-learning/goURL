package http

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

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
	Header  map[string]string
	Query   map[string]string
	Body    string
	Timeout time.Duration
}

// NewClient creates a new http client based on given configuration.
func NewClient(
	method string,
	url string,
	header map[string]string,
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

func (c *Client) Do() (*http.Response, error) {
	client := &http.Client{
		Transport: &http.Transport{
			ResponseHeaderTimeout: c.Timeout,
		},
	}

	logrus.Infof("sending request into %s", c.URL)

	req, err := http.NewRequestWithContext(context.Background(), c.Method, c.URL, bytes.NewBuffer([]byte(c.Body)))
	if err != nil {
		return nil, fmt.Errorf("request creation failed: %w", err)
	}

	for k, v := range c.Header {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, fmt.Errorf("request timeout: %w", err)
		}

		return nil, fmt.Errorf("request failed: %w", err)
	}

	logrus.Infof("Method is: %s", c.Method)
	logrus.Infof("Response status: %s", resp.Status)

	logrus.Info("headers are:")

	for k, v := range resp.Header {
		logrus.Infof("%s = %s", k, v)
	}

	return resp, nil
}
