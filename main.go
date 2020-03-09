package main

import (
	"fmt"

	"github.com/elastos/eos-sc-adapter/servers/httpjsonrpc"
)

func main()  {
	fmt.Println("eos-sc-adapter started")
	httpjsonrpc.StartRPCServer()
}
