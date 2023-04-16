package blockchain

import "github.com/btcsuite/btcd/btcutil"

type Addresses struct {
	Addresses []Address `json:"addresses"`
}

type Address struct {
	Address       string `json:"address"`
	FinalBalance  int64  `json:"final_balance"`
	NTx           int64  `json:"n_tx"`
	TotalReceived int64  `json:"total_received"`
	TotalSent     int64  `json:"total_sent"`
}

func (a *Address) FinalBalanceAmount() string {
	return btcutil.Amount(a.FinalBalance).String()
}

func (a *Address) TotalReceivedAmount() string {
	return btcutil.Amount(a.TotalReceived).String()
}

func (a *Address) TotalSentAmount() string {
	return btcutil.Amount(a.TotalSent).String()
}
