package clierr

import (
	"errors"
	"fmt"

	cli "gopkg.in/urfave/cli.v1"
)

const (
	// ErrCouldNotInit is returned upon a client initialization failure
	ErrCouldNotInit = "error initializing IPrepd client"
)

// Handle determines what verbosity and exposure a user gets to see
// given context flag values and the error type
func Handle(ctx *cli.Context, pretty string, e error) error {
	if ctx.GlobalBool("verbose") {
		return fmt.Errorf("%s: %s", pretty, e)
	}
	return errors.New(pretty)
}
