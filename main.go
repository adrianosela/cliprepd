package main

import (
	"log"
	"os"

	"github.com/adrianosela/cliprepd/tool"
)

// injected at build time
var version string

func main() {
	if err := tool.GetApp(version).Run(os.Args); err != nil {
		log.Fatalf("error: %s", err)
	}
}
