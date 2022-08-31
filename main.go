package main

import (
	"gitlab.com/tokend/blobs/internal/cli"
	"os"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
