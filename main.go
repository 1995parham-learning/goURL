package main

import (
	"flag"
	"fmt"
	"goURL/http"
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

	var headerFlag http.ArrayFlag
	var queryFlag http.ArrayFlag

	flag.Var(&headerFlag, "H", "header")
	flag.Var(&queryFlag, "Q", "query parameter")
	flag.CommandLine.Parse(os.Args[2:])

	header, warning := headerFlag.ToHeaderMap()
	fmt.Println(warning)

	query, warning := queryFlag.ToHeaderMap()
	fmt.Println(warning)
	//url.Parse()

	client := http.NewClient(*method, URL, header, query)
	client.Do()

}

func Validate(url string) bool {
	var validURL = regexp.MustCompile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)

	return validURL.MatchString(url)
}
