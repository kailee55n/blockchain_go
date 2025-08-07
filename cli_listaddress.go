// cli_listaddress.go
// This file is part of the Go Blockchain project, which implements a simple CLI for listing addresses.
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"log"
)

func (cli *CLI) listAddresses(nodeID string) {
	wallets, err := NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}
