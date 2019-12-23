package main

import (
	"os"

	"github.com/cbrwflo/puppetbeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
