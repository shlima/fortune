package mapper

import (
	"fmt"
	"strings"

	"github.com/shlima/fortune/internal/pkg/key"
)

func KeyChainHTML(chain key.KeyChain) string {
	b := new(strings.Builder)
	b.WriteString("💰")
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("<b>Private</b>: %s", chain.Private))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf(`<b>Compressed</b>: <a href="%s">%s</a>`, BlockchainLink(chain.Compressed), chain.Compressed))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf(`<b>Uncompressed</b>: <a href="%s">%s</a>`, BlockchainLink(chain.Uncompressed), chain.Uncompressed))

	return b.String()
}

func BlockchainLink(address string) string {
	return fmt.Sprintf("https://www.blockchain.com/explorer/addresses/btc/%s", address)
}
