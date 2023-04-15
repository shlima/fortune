package cmd

import (
	"fmt"
	"github.com/shlima/fortune/internal/pkg/datum"
	"github.com/urfave/cli/v2"
)

func NewIndex(c *cli.Context) datum.Index {
	index, err := datum.ReadFiles(c.StringSlice(FlagFiles.Name)...)
	if err != nil {
		panic(fmt.Errorf("failed to read datum: %w", err))
	}

	return index
}
