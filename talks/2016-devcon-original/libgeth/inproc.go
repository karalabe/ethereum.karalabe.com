package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/node"
	"golang.org/x/net/context"
)

var ctx = context.Background()

// START OMIT
func main() {
	stackConf := &node.Config{DataDir: "/home/karalabe/.ethereum", ListenAddr: ":0"}
	stack, _ := node.New(stackConf) // HL

	ethConf := &eth.Config{ChainConfig: utils.MainnetChainConfig}
	_ = stack.Register(func(ctx *node.ServiceContext) (node.Service, error) { // HL
		return eth.New(ctx, ethConf)
	})

	_ = stack.Start() // HL
	fmt.Println("My datadir:", stack.DataDir())
	fmt.Println("My address:", stack.Server().NodeInfo().ListenAddr, "\n")

	api, _ := stack.Attach() // HL
	block, _ := ethclient.NewClient(api).BlockByNumber(ctx, nil)
	fmt.Println("Latest block:", block.Number())
}

// END OMIT
