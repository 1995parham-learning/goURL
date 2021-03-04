package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const (
	GET = "GET"
	POST = "POST"
	PATCH = "PATCH"
	DELETE = "DELETE"
	PUT = "PUT"
)

// Created so that multiple inputs can be accecpted
type arrayFlags []string

func (i *arrayFlags) String() string {
	// change this, this is just can example to satisfy the interface
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, strings.TrimSpace(value))
	return nil
}

func main() {

	// The first argument is always the URL
	if len(os.Args) == 1 {
		fmt.Println("URL is not given")
		return
	}

	URL := os.Args[1]
	if !Validate(URL) {
		fmt.Println("URL is not in valid format")
	}

	method := flag.String("M", "GET", "method")
	var myFlags arrayFlags

	flag.Var(&myFlags, "H", "header")
	flag.CommandLine.Parse(os.Args[2:])

	header := make(map[string]string)
	warning := ""

	for _, f := range myFlags {
		a := strings.Split(f, ",")
		for _, b := range a {
			keyValue := strings.Split(b, ":")
			_, ok := header[keyValue[0]]
			if ok {
				warning += fmt.Sprintf("%s is a repetead header so only the last one is considered", keyValue[0])
			}

			header[keyValue[0]] = keyValue[1]
		}
	}

	fmt.Println(warning)

	switch *method {
	case GET:
		do(URL, GET)
	case POST:
		do(URL, POST)
	case PATCH:
		do(URL, PATCH)
	case DELETE:
		do(URL, DELETE)
	case PUT:
		do(URL, PUT)
	default:
		fmt.Println("Method is not recognized. Use GET, POST, PUT, PATCH and DELETE instead")
	}

}

func do(url string, method string) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
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

func Validate(url string) bool {
	var validURL = regexp.MustCompile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)

	return validURL.MatchString(url)
}
