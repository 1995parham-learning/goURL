package cmd

import (
	"log"

	"github.com/elahe-dastan/goURL/internal/cmd/url"
)

func Execute() {
	rootCmd := url.Build()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
