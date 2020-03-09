// Copyright (c) 2017-2019 The Elastos Foundation
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.
//

package servers

import (
	"github.com/elastos/eos-sc-adapter/servers/eos"
	. "github.com/elastos/eos-sc-adapter/servers/errors"
)

type SidechainIllegalDataInfo struct {
	IllegalType     uint8  `json:"illegaltype"`
	Height          uint32 `json:"height"`
	IllegalSigner   string `json:"illegalsigner"`
	Evidence        string `json:"evidence"`
	CompareEvidence string `json:"compareevidence"`
}

func GetBlockCount(param Params) map[string]interface{} {
	info, err := eos.GetInfo()
	if err != nil {
		return ResponsePack(Error, err)
	}
	return ResponsePack(Success, info.LastIrreversibleBlockNum+1)
}

func GetBlockByHeight(param Params) map[string]interface{} {
	info, err := eos.GetInfo()
	if err != nil {
		return ResponsePack(Error, err)
	}
	return ResponsePack(Success, info.LastIrreversibleBlockNum)
}

func GetWithdrawTransactionsByHeight(param Params) map[string]interface{} {
	height, ok := param.Uint("height")
	if !ok {
		return ResponsePack(Error, "need param height")
	}
	withdrawInfo, err := eos.GetWithdrawTransactionsByHeight(height)
	if err != nil {
		return ResponsePack(Error, err.Error())
	}

	return ResponsePack(Success, withdrawInfo)
}

func GetWithdrawTransactionByHash(param Params) map[string]interface{} {
	txID, ok := param.String("txid")
	if !ok {
		return ResponsePack(Error, "need param txid")
	}
	withdrawInfo, err := eos.GetWithdrawTransactionByHash(txID)
	if err != nil {
		return ResponsePack(Error, err.Error())
	}

	return ResponsePack(Success, withdrawInfo)
}

func GetIllegalEvidenceByHeight(param Params) map[string]interface{} {
	result := make([]*SidechainIllegalDataInfo, 0)
	return ResponsePack(Success, result)
}

func CheckIllegalEvidence(param Params) map[string]interface{} {
	return ResponsePack(Success, false)
}

func GetExistDepositTransactions(param Params) map[string]interface{} {
	txs, ok := GetStringArray(param, "txs")
	if !ok {
		return ResponsePack(Error, "need param txs")
	}
	return ResponsePack(Success, txs)
}

func SendRechargeTransaction(param Params) map[string]interface{} {
	txID, ok := param.String("txid")
	if !ok {
		return ResponsePack(Error, "need param txid")
	}
	result, err := eos.SendRechargeTx(txID)
	if err != nil {
		return ResponsePack(Error, err.Error())
	}

	return ResponsePack(Success, result)
}

func GetStringArray(param Params, key string) ([]string, bool) {
	value, ok := param[key]
	if !ok {
		return nil, false
	}
	switch v := value.(type) {
	case []interface{}:
		var arrayString []string
		for _, param := range v {
			paramString, ok := param.(string)
			if !ok {
				return nil, false
			}
			arrayString = append(arrayString, paramString)
		}
		return arrayString, true
	default:
		return nil, false
	}
}

func ResponsePack(errCode ServerErrCode, result interface{}) map[string]interface{} {
	if errCode != 0 && (result == "" || result == nil) {
		result = ErrMap[errCode]
	}
	return map[string]interface{}{"Result": result, "Error": errCode}
}
