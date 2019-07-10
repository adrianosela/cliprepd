package main

import (
	"log"
	"os"

	"github.com/adrianosela/cliprepd/app"
)

// injected at build time
var version string

func main() {
	if err := app.New(version).Run(os.Args); err != nil {
		log.Fatalf("error: %s", err)
	}
}
