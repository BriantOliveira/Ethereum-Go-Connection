package main

import (
	"log"
	"context"
	"fmt"

	"Github.com/ethereum/go-ethereum/common"
	"Github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	conn, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatalf("Woops something went wrong!")
	}

	ctx := context.Background()
	tx, pending, _ := conn.TransactionByHash(ctx, common.HexToHash("0x6fb1730dab43028502058d41acb615a7c1b0ab2d108f729d310509c66f62cf26"))
	if !pending {
		fmt.Println(tx)
	}
}