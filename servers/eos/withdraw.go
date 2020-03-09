package eos

import (
	"fmt"
	"strings"

	"github.com/eoscanada/eos-go"
)

// Inspired by https://github.com/eoscanada/eos-go/issues/100

func GetWithdrawTransactionsByHeight(num uint32) (interface{}, error) {
	block, err := api.GetBlockByNum(num)
	if err != nil {
		return nil, fmt.Errorf("get block: %s", err)
	}

	var trans []*WithdrawTxInfo
	for _, rawTx := range block.Transactions {
		signTx, err := rawTx.Transaction.Packed.Unpack()
		if err != nil {
			return nil, err
		}
		crossChainInfo, err := getCrossChainInfo(signTx.Transaction)
		if err != nil {
			return nil, err
		}
		if len(crossChainInfo) == 0 {
			continue
		}

		txWithdraw := &WithdrawTxInfo{
			TxID:             rawTx.Transaction.ID.String(),
			CrossChainAssets: crossChainInfo,
		}
		trans = append(trans, txWithdraw)
	}

	return trans, nil
}

func GetWithdrawTransactionByHash(txID string) (interface{}, error) {
	txResp, err := api.GetTransaction(txID)
	if err != nil {
		return nil, fmt.Errorf("get transaction: %s", err)
	}
	var trans []*WithdrawTxInfo
	crossChainInfo, err := getCrossChainInfo(txResp.Transaction.Transaction.Transaction)
	if err != nil {
		return nil, err
	}
	txWithdraw := &WithdrawTxInfo{
		TxID:             txResp.ID.String(),
		CrossChainAssets: crossChainInfo,
	}
	trans = append(trans, txWithdraw)

	return trans, nil
}

func getCrossChainInfo(tx *eos.Transaction) ([]*WithdrawOutputInfo, error) {
	var txOuputsInfo []*WithdrawOutputInfo
	for _, action := range tx.Actions {
		if action.Account == ActionAccount && action.Name == WithdrawName {
			err := action.MapToRegisteredAction()
			if err != nil {
				return nil, err
			}
			withdraw, ok := action.Data.(*Withdraw)
			if !ok {
				return nil, fmt.Errorf("unpack withdraw action failed")
			}
			txOuputsInfo = append(txOuputsInfo, &WithdrawOutputInfo{
				CrossChainAddress: withdraw.To,
				CrossChainAmount:  strings.Split(withdraw.Quantity.String(), " ")[0],
				OutputAmount:      strings.Split(withdraw.Quantity.Add(withdraw.Fee).String(), " ")[0],
			})
		}
	}
	return txOuputsInfo, nil
}
