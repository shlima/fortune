package cmd

import (
	"fmt"
	"github.com/shlima/fortune/internal/pkg/datum"
	"github.com/urfave/cli/v2"
)

var FlagFiles = &cli.StringSliceFlag{
	Name:    "file",
	EnvVars: []string{"FILE"},
	Value: cli.NewStringSlice(
		"addresses/Bitcoin/2023/04/p2pkh_Rich_Max_1.txt",
		"addresses/Bitcoin/2023/04/p2pkh_Rich_Max_10.txt",
		"addresses/Bitcoin/2023/04/p2pkh_Rich_Max_100.txt",
		"addresses/Bitcoin/2023/04/p2pkh_Rich_Max_1000.txt",
		"addresses/Bitcoin/2023/04/p2pkh_Rich_Max_10000.txt",
		"addresses/Bitcoin/2023/04/p2pkh_Rich_Max_100000.txt",
	),
}

func NewIndex(c *cli.Context) datum.Index {
	index, err := datum.ReadFiles(c.StringSlice(FlagFiles.Name)...)
	if err != nil {
		panic(fmt.Errorf("failed to read datum: %w", err))
	}

	return index
}
