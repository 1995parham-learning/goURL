package cmd

import (
	"log"

	"github.com/1995parham-learning/gourl/internal/cmd/url"
)

func Execute() {
	rootCmd := url.Build()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
