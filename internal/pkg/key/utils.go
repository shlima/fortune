package key

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

func KeyChainFromPriv(priv []byte) (out KeyChain, err error) {
	_, pub := btcec.PrivKeyFromBytes(priv)
	compressed, err := btcutil.NewAddressPubKey(pub.SerializeCompressed(), &chaincfg.MainNetParams)
	if err != nil {
		return out, fmt.Errorf("failed to compress key: %w", err)
	}

	uncompressed, err := btcutil.NewAddressPubKey(pub.SerializeUncompressed(), &chaincfg.MainNetParams)
	if err != nil {
		return out, fmt.Errorf("failed to uncompress key")
	}

	return KeyChain{
		Private:      hex.EncodeToString(priv),
		Compressed:   compressed.EncodeAddress(),
		Uncompressed: uncompressed.EncodeAddress(),
	}, nil
}

// PrivToWIF converts raw private key bytes to WIF format
func PrivToWIF(key []byte) (out string, err error) {
	priv, _ := btcec.PrivKeyFromBytes(key)
	wif, err := btcutil.NewWIF(priv, &chaincfg.MainNetParams, false)
	if err != nil {
		return out, fmt.Errorf("failed to new wif: %w", err)
	}

	return wif.String(), nil
}
