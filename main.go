package main

import (
	"flag"
	"fmt"
	"goURL/http"
	"os"
	"regexp"
	"strings"
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

	//url.Parse()

	client := http.Client{
		Method: *method,
		URL:    URL,
	}

	client.Do()

}

func Validate(url string) bool {
	var validURL = regexp.MustCompile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)

	return validURL.MatchString(url)
}
