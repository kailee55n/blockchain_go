package main

import (
	"fmt"
	"log"
)

func (cli *CLI) startNode(nodeID, minerAddress string) {
	fmt.Printf("Starting node %s\n", nodeID)
	if len(minerAddress) > 0 {
		if ValidateAddress(minerAddress) {
			fmt.Println("Mining is on. Address to receive rewards: ", minerAddress)
			// Start the mining process
			bc := NewBlockchain(nodeID)
			defer bc.db.Close()
			UTXOSet := UTXOSet{bc}
			miner := NewMiner(minerAddress, &UTXOSet)
			go miner.Mine()
			fmt.Printf("Node %s started with mining enabled.\n", nodeID)
		} else {
			log.Panic("Wrong miner address!")
		}
	}
	StartServer(nodeID, minerAddress)
}
