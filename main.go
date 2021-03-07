package main

import (
	"flag"
	"fmt"
	"goURL/http"
	"io/ioutil"
	"os"
	"regexp"
)

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
	body := flag.String("D", "", "body")
	// --
	json := flag.Bool("json", false, "content type header")
	file := flag.String("file", "", "file path as body")
	timeout := flag.Int("timeout", 1000, "timeout")

	if *file != "" {
		dat, err := ioutil.ReadFile(*file)
		if err != nil {
			panic(err)
		}

		*body = string(dat)
	}

	var headerFlag http.ArrayFlag
	var queryFlag http.ArrayFlag

	flag.Var(&headerFlag, "H", "header")
	flag.Var(&queryFlag, "Q", "query parameter")
	err := flag.CommandLine.Parse(os.Args[2:])
	if err != nil {
		panic(err)
	}

	header, warning := headerFlag.ToHeaderMap(*json)
	fmt.Println(warning)

	query, warning := queryFlag.ToQueryMap()
	fmt.Println(warning)
	//url.Parse()

	client := http.NewClient(*method, URL, header, query, *body, *timeout)
	client.Do()

}

func Validate(url string) bool {
	var validURL = regexp.MustCompile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)

	return validURL.MatchString(url)
}
