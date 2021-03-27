package cmd

import (
	"encoding/json"
	"fmt"
	"goURL/http"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goURL",
	Short: "Trying to act like cURL",
	Run: func(cmd *cobra.Command, args []string) {
		// The first argument is always the URL
		if len(os.Args) == 1 {
			fmt.Println("URL is not given")
			return
		}

		URL := os.Args[1]
		if !Validate(URL) {
			fmt.Println("URL is not in valid format")
		}

		var body string
		var format string

		if file != "" {
			dat, err := ioutil.ReadFile(file)
			if err != nil {
				panic(err)
			}

			format = "application/octet-stream"
			body = string(dat)
		}

		headerFlags := http.New(headers)
		queryFlags := http.New(queries)

		if data != "" && jsn != "" {
			log.Fatal("You can whether use --data or --json")
		}

		if data != "" {
			format = "application/json"
			body = data
		} else {
			format = "application/x-www-form-urlencoded"
			body = jsn
		}

		header, warning := headerFlags.ToHeaderMap(format)
		fmt.Println(warning)

		query, warning := queryFlags.ToQueryMap()
		fmt.Println(warning)
		//url.Parse()

		// *******************************************************************************
		if header["content-type"][0] == "application/x-www-form-urlencoded" {
			match, err := regexp.MatchString("([^&]+=[^&]*(&[^&]+=[^&]*)*)?", data)
			if err != nil {
				log.Fatal(err)
			}

			if !match {
				fmt.Println("Your body is not in the default format x-www-form-urlencoded")
			}
		} else {
			var js map[string]interface{}
			if json.Unmarshal([]byte(data), &js) != nil {
				fmt.Println("Your body is not in the json format")
			}
		}

		client := http.NewClient(method, URL, header, query, body, timeout)
		client.Do()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var method string
var data string
var jsn string
var file string
var timeout time.Duration
var headers []string
var queries []string

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goURL.yaml)")

	rootCmd.PersistentFlags().StringVarP(&method, "method", "M", "GET", "specify your method")
	rootCmd.PersistentFlags().StringVarP(&data, "data", "D", "", "specify your data with Content-Type header as application/x-www-form-urlencoded")
	rootCmd.PersistentFlags().StringVar(&jsn, "json", "", "specify your body with Content-Type header as application/json")
	rootCmd.PersistentFlags().StringVar(&file, "file", "", "specify a file path to put the file as the request data")
	rootCmd.PersistentFlags().DurationVar(&timeout, "timeout", 0, "specify timeout")
	rootCmd.PersistentFlags().StringSliceVarP(&headers, "headers", "H", nil, "specify header")
	rootCmd.PersistentFlags().StringSliceVarP(&queries, "queries", "Q", nil, "specify queries")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}

		// Search config in home directory with name ".goURL" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".goURL")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func Validate(url string) bool {
	var validURL = regexp.MustCompile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)

	return validURL.MatchString(url)
}
