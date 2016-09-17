package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"
)

// START OMIT

var tipjar = common.HexToAddress("0xfb6916095ca1df60bb79ce92ce3ea74c37c5d359")

func main() {
	ec, _ := ethclient.Dial("/home/karalabe/.ethereum/geth.ipc") // HL

	head, _ := ec.BlockByNumber(ctx, nil) // HL
	fmt.Println("Chain height:", head.Number())

	balance, _ := ec.BalanceAt(ctx, tipjar, nil) // HL
	fmt.Println("TipJar funds:", balance)

	heads := make(chan *types.Header, 16)
	_, _ = ec.SubscribeNewHead(ctx, heads) // HL
	for header := range heads {
		fmt.Printf("#%d: %s…\n", header.Number, header.Hash().Hex()[:10])
	}
}

// END OMIT

var ctx = context.Background()
