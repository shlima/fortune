package cmd

import (
	"fmt"
	"strings"

	"github.com/shlima/fortune/internal/pkg/mapper"
	"github.com/urfave/cli/v2"
)

func Brain(c *cli.Context) error {
	password := c.Args().First()
	got, err := NewKeyGen(c).BrainSHA256([]byte(password))
	if err != nil {
		return fmt.Errorf("faield to generate: %w", err)
	}

	fmt.Printf(strings.TrimPrefix(`
Password: %s
Private: %s
Compressed: %s
Uncompressed: %s
`, "\n"),
		password,
		got.Private,
		mapper.BlockchainLink(got.Compressed),
		mapper.BlockchainLink(got.Uncompressed),
	)

	return nil
}
