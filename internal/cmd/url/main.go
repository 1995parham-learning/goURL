package url

import (
	"encoding/json"
	"io"
	"net/url"
	"os"
	"regexp"

	"github.com/1995parham-learning/gourl/internal/css"
	"github.com/1995parham-learning/gourl/internal/http"
	"github.com/cheggaaa/pb/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	MethodFlag  = "method"
	DataFlag    = "data"
	JSONFlag    = "json"
	TimeoutFlag = "timeout"
	HeaderFlag  = "header"
	QueryFlag   = "query"
	FileFlag    = "file"
)

// nolint: funlen, cyclop
func main(cmd *cobra.Command, args []string) {
	method, _ := cmd.Flags().GetString(MethodFlag)
	data, _ := cmd.Flags().GetString(DataFlag)
	jsonData, _ := cmd.Flags().GetString(JSONFlag)
	file, _ := cmd.Flags().GetString(FileFlag)

	timeout, err := cmd.Flags().GetDuration(TimeoutFlag)
	if err != nil {
		logrus.Errorf("timeout flag: %s", err)

		return
	}

	headers, err := cmd.Flags().GetStringSlice(HeaderFlag)
	if err != nil {
		logrus.Errorf("header flag: %s", err)

		return
	}

	queries, err := cmd.Flags().GetStringSlice(QueryFlag)
	if err != nil {
		logrus.Errorf("query flag: %s", err)

		return
	}

	// The first argument is always the URL
	if len(args) != 1 {
		logrus.Error("URL is not given")

		return
	}

	URL := args[0]
	{
		u, err := url.Parse(URL)
		if err != nil {
			logrus.Errorf("URL isn't valid %s", err)

			return
		}
		if u.Scheme != "https" && u.Scheme != "http" {
			logrus.Errorf("URL isn't valid because of incorrect scheme")

			return
		}
	}

	var body, format string

	if file != "" {
		dat, err := os.ReadFile(file)
		if err != nil {
			logrus.Errorf("cannot read %s: %s", file, err)

			return
		}

		body = string(dat)
	}

	if data != "" && jsonData != "" {
		logrus.Error("You can whether use --data or --json")

		return
	}

	switch {
	case data != "":
		format = "application/x-www-form-urlencoded"
		body = data
	case jsonData != "":
		format = "application/json"
		body = jsonData
	case file != "":
		format = "application/octet-stream"
	}

	h, err := css.NewWithOptions(headers, ":", ",").ToMap()
	if err != nil {
		logrus.Warn(err)
	}

	q, err := css.NewWithOptions(queries, "=", "&").ToMap()
	if err != nil {
		logrus.Warn(err)
	}

	if h["Content-Type"] == "" {
		h["Content-Type"] = format
	}

	switch h["Content-Type"] {
	case "application/x-www-form-urlencoded":
		// validate form data by a regular expression
		match, err := regexp.MatchString("^([^&]+=[^&]*(&[^&]+=[^&]*)*)?$", data)
		if err != nil {
			logrus.Fatal(err)
		}

		if !match {
			logrus.Error("your body is not in the default format x-www-form-urlencoded")
		}
	case "application/json":
		// only validate the json by parsing it
		var js map[string]interface{}
		if err := json.Unmarshal([]byte(jsonData), &js); err != nil {
			logrus.Errorf("your body is not in the json format: %s", err)
		}
	}

	client := http.NewClient(method, URL, h, q, body, timeout)

	resp, err := client.Do()
	if err != nil {
		logrus.Error(err)

		return
	}

	defer resp.Body.Close()

	rd := resp.Body

	var bar *pb.ProgressBar

	if resp.ContentLength > 0 {
		// start a new progress bar
		bar = pb.Full.Start64(resp.ContentLength)

		// creates a proxy reader for showing the progress bar
		rd = bar.NewProxyReader(resp.Body)
	}

	respBody, err := io.ReadAll(rd)
	if err != nil {
		logrus.Errorf("reading body failed: %s", err)

		return
	}

	if bar != nil {
		bar.Finish()
	}

	cmd.Println(string(respBody))
}

// Create and return url command.
func Build() *cobra.Command {
	// nolint: exhaustivestruct
	cmd := &cobra.Command{
		Use:   "goURL",
		Short: "cURL clone in Go",
		Run: func(cmd *cobra.Command, args []string) {
			main(cmd, args)
		},
	}

	cmd.Flags().StringP(MethodFlag, "X", "GET", "specify your method")
	cmd.Flags().StringP(DataFlag, "d", "",
		"specify your data with Content-Type header as application/x-www-form-urlencoded")
	cmd.Flags().StringP(JSONFlag, "j", "", "specify your body with Content-Type header as application/json")
	cmd.Flags().StringP("file", "D", "", "specify a file path to put the file as the request data")
	cmd.Flags().DurationP(TimeoutFlag, "t", 0, "specify timeout")
	cmd.Flags().StringSliceP(HeaderFlag, "H", nil, "specify header")
	cmd.Flags().StringSliceP(QueryFlag, "Q", nil, "specify queries")

	return cmd
}
