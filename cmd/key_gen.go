package cmd

import (
	"github.com/shlima/fortune/internal/pkg/key"
	"github.com/urfave/cli/v2"
)

func NewKeyGen(c *cli.Context) *key.Generator {
	return key.New()
}
