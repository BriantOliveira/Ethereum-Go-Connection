package main

import (
	"Github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	conn, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatalf("Woops something went wrong!")
	}
}