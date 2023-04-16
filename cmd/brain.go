package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/logger"
	"github.com/olekukonko/tablewriter"
	"github.com/shlima/fortune/internal/pkg/blockchain"
	"github.com/shlima/fortune/internal/pkg/mapper"
	"github.com/urfave/cli/v2"
)

func Brain(c *cli.Context) error {
	password := c.Args().First()
	got, err := NewKeyGen(c).BrainSHA256([]byte(password))
	if err != nil {
		return fmt.Errorf("faield to generate: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	addresses, err := blockchain.New().Addresses(ctx, []string{got.Compressed, got.Uncompressed})
	if err != nil {
		logger.Error(fmt.Errorf("failed to get address balance: %w", err))
	}

	data := [][]string{
		{"Password", fmt.Sprintf("<%s>", password)},
		{"Private", got.Private},
		{"Compressed", mapper.BlockchainLink(got.Compressed)},
		{"Uncompressed", mapper.BlockchainLink(got.Uncompressed)},
	}

	for _, a := range addresses {
		data = append(data, []string{
			a.Address,
			fmt.Sprintf("%s (%d tx)", a.FinalBalanceAmount(), a.NTx),
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
	return nil
}
