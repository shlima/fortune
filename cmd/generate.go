package cmd

import (
	"fmt"
	"strings"

	"github.com/shlima/fortune/internal/mapper"
	"github.com/urfave/cli/v2"
)

func Generate(c *cli.Context) error {
	got, err := NewKeyGen(c).Generate()
	if err != nil {
		return fmt.Errorf("faield to generate: %w", err)
	}

	fmt.Printf(strings.TrimPrefix(`
Private: %s
Compressed: %s
Uncompressed: %s
`, "\n"),
		got.Private,
		mapper.BlockchainLink(got.Compressed),
		mapper.BlockchainLink(got.Uncompressed),
	)

	return nil
}
