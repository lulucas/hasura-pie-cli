package errors

import "github.com/urfave/cli/v2"

var (
	ErrMissingPath   = cli.Exit("missing path", 1)
	ErrAlreadyExists = cli.Exit("already exists", 2)
)
