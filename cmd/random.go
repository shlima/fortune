package cmd

import (
	"fmt"
	"strings"

	"github.com/shlima/fortune/internal/pkg/mapper"
	"github.com/urfave/cli/v2"
)

func Random(c *cli.Context) error {
	address, index := NewIndex(c).Random()
	fmt.Printf(strings.TrimPrefix(`
Index: %d
Random address: %s
URL: %s
`, "\n"),
		index,
		address,
		mapper.BlockchainLink(address),
	)

	return nil
}
