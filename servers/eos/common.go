package eos

import (
	"fmt"
	"github.com/eoscanada/eos-go"
)

const (
	ActionAccount = "eosio.token"
	RechargeName  = "recharge"
	WithdrawName  = "withdraw"
)

var (
	api *eos.API
)

func init() {
	eos.RegisterAction(ActionAccount, WithdrawName, Withdraw{})
	api = eos.New("http://127.0.0.1:8888")
	keyBag := eos.NewKeyBag()
	err := keyBag.ImportPrivateKey("5KQwrPbwdL6PhXujxW37FSSQZ1JiwsST4cqQzDeyXtP79zkvFD3")
	if err != nil {
		panic(fmt.Errorf("import private key: %s", err))
	}
	api.SetSigner(keyBag)
}

// Recharge represents the `recharge` struct on the `eosio.token` contract.
type Recharge struct {
	TxID     string          `json:"txid"`
	To       eos.AccountName `json:"to"`
	Quantity eos.Asset       `json:"quantity"`
	Fee      eos.Asset       `json:"fee"`
}

// Withdraw represents the `withdraw` struct on the `eosio.token` contract.
type Withdraw struct {
	From     eos.AccountName `json:"from"`
	To       string          `json:"to"`
	Quantity eos.Asset       `json:"quantity"`
	Fee      eos.Asset       `json:"fee"`
}

type WithdrawOutputInfo struct {
	CrossChainAddress string `json:"crosschainaddress"`
	CrossChainAmount  string `json:"crosschainamount"`
	OutputAmount      string `json:"outputamount"`
}

type WithdrawTxInfo struct {
	TxID             string                `json:"txid"`
	CrossChainAssets []*WithdrawOutputInfo `json:"crosschainassets"`
}
