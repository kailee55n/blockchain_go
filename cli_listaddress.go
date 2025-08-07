package main

import (
	"fmt"
	"log"
)

func (cli *CLI) listAddresses(nodeID string) {
	wallets, err := NewWallets(nodeID) // Initialize wallets
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses() // Retrieve all addresses from the wallets

	for _, address := range addresses {
		fmt.Println(address)
	}
}
