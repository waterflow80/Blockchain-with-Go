package main

import (

)

// Blockchain implements interactions with a DB
type Blockchain struct {
	GHash []byte // Hash of the Genesis Block
	Chain []Block// Slice of blocks
}

// AddBlock adds a new block with the provided transactions, the block
// is mined before addition 
func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	// What happens when the chain is empty?

}

// NewSendTransaction creates a set of transactions to transfert the
// amount from sender to receiver. If the sender has not the required
// amount, an error is returned 
func  (bc *Blockchain) NewTransfertTX(from, to string, amount int) (*Transaction,error) {
	return nil,nil
}

// CreateBlockchain creates a new blockchain, evey adress in adresses
// is given the initial 
func NewBlockchain(addresses []string) *Blockchain {
	return nil
}

// creates a new blockchain given a valid genesis block
func NewBlockchainFromGB(genesis *Block) *Blockchain {
	return nil
}

func (bc *Blockchain) GetBalance(address string) int {
	balance := 0
	return balance
}
