package blockchain

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type Cli struct {
	client *resty.Client
}

func New() *Cli {
	return &Cli{
		client: resty.New(),
	}
}

func (c *Cli) Addresses(ctx context.Context, list []string) (out []Address, err error) {
	json := new(Addresses)

	resp, err := c.client.R().
		SetQueryParams(map[string]string{
			"active": strings.Join(list, "|"),
			"n":      "1",
		}).
		SetHeader("Accept", "application/json").
		SetContext(ctx).
		SetResult(json).
		Get("https://blockchain.info/multiaddr")

	switch {
	case err != nil:
		return nil, fmt.Errorf("failed to query: %w", err)
	case !resp.IsSuccess():
		return nil, fmt.Errorf("response is not success: <%s>", resp.Status())
	}

	return json.Addresses, nil
}
