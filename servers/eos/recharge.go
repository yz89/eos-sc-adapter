package eos

import (
	"github.com/eoscanada/eos-go"
)

// Inspired by https://github.com/eoscanada/eos-go/blob/master/example_api_transfer_eos_test.go

func SendRechargeTx(txID string) (out *eos.PushTransactionFullResp, err error) {
	quantity, err := eos.NewAssetFromString("10.0000 SYS")
	if err != nil {
		return nil, err
	}
	fee, err := eos.NewAssetFromString("0.1000 SYS")
	if err != nil {
		return nil, err
	}
	action := &eos.Action{
		Account: "eosio.token",
		Name:    "recharge",
		Authorization: []eos.PermissionLevel{
			{
				Actor:      ActionAccount,
				Permission: "active",
			},
		},
		ActionData: eos.NewActionData(Recharge{
			TxID:     txID,
			To:       "bob",
			Quantity: quantity,
			Fee:      fee,
		}),
	}

	return api.SignPushActions(action)
}

//func cli()  {
	//contractParam := fmt.Sprintf(`[ "%s", "bob", "100.0000 SYS", "1.0000 SYS" ]`,
	//	txID,
	//)
	//cmd := exec.Command(
	//	"/Users/yzhou/eosio/2.0/bin/cleos",
	//	"push",
	//	"action",
	//	"eosio.token",
	//	"recharge",
	//	contractParam,
	//	"-p",
	//	"eosio.token@active",
	//)
	//
	//return cmd.Run()
//}