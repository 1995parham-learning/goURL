package array

import (
	"fmt"
	"strings"
)

// Created so that multiple inputs can be accecpted
type ArrayFlag []string

func New(arr []string) *ArrayFlag {
	a := ArrayFlag{}
	a = arr

	return &a
}

// ToHeaderMap creates a http header map based on given array of headers.
// format replaces the content-type if it doesn't exist.
func (a *ArrayFlag) ToHeaderMap(format string) (map[string][]string, string) {
	headers := make(map[string][]string)
	warning := ""

	for _, header := range *a {
		splitedHeader := strings.Split(header, ":")

		name := strings.ToLower(splitedHeader[0])
		value := splitedHeader[1]

		if _, ok := headers[name]; ok {
			warning += fmt.Sprintf("%s is a repetead header so only the last one is considered", name)
		}

		headers[name] = []string{value}
	}

	_, ok := headers["content-type"]
	if !ok && format != "" {
		headers["content-type"] = []string{format}
	}

	return headers, warning
}

func (a *ArrayFlag) ToQueryMap() (map[string]string, string) {
	query := make(map[string]string)
	warning := ""

	for _, q := range *a {
		a := strings.Split(q, "&")
		for _, b := range a {
			keyValue := strings.Split(b, "=")

			if _, ok := query[strings.ToLower(keyValue[0])]; ok {
				warning += fmt.Sprintf("%s is a repetead query parameter so only the last one is considered", keyValue[0])
			}

			query[strings.ToLower(keyValue[0])] = keyValue[1]
		}
	}

	return query, warning
}
