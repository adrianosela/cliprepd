package main

import (
	"fmt"
	"os"

	"github.com/adrianosela/cliprepd/app"
)

// injected at build time
var version string

func main() {
	if err := app.New(version).Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
