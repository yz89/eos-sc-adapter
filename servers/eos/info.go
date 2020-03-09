package eos

import "github.com/eoscanada/eos-go"

func GetInfo() (out *eos.InfoResp, err error) {
	return api.GetInfo()
}
