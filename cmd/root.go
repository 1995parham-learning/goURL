package cmd

import (
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"goURL/http"
	"io/ioutil"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goURL",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
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

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goURL.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

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
