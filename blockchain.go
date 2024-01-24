package main

import (
	
)

const ZEROBITS int = 17 // Constraint on the hash (used in mining)

// Blockchain implements interactions with a DB
type Blockchain struct {
	GHash []byte // Hash of the Genesis Block
	Chain []Block// Slice of blocks
}

// AddBlock adds a new block with the provided transactions, the block
// is mined before addition 
func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	if len(bc.GHash) == 0 {
		// Chain is empty
		// Adding a the Genisis block
		GBlock := NewGBlock(transactions, ZEROBITS)
		bc.GHash = GBlock.Hash
		bc.Chain = append(bc.Chain, *GBlock)
	} else {
		// Chain is not empty
		block := NewBlock(transactions, bc.Chain[len(bc.Chain)-1].Hash, true, ZEROBITS)
		bc.Chain = append(bc.Chain, *block)
	}
}


// NewSendTransaction creates a set of transactions to transfert the
// amount from sender to receiver. If the sender has not the required
// amount, an error is returned 
func  (bc *Blockchain) NewTransfertTX(from, to string, amount int) (*Transaction,error) {
	if bc.GetBalance(from) < amount {
		return nil, ErrInsufficientFunds
	}
	txin  := TXInput{[]byte{}, -1, from}
	txout := TXOutput{amount, to}
	transaction := NewTransaction([]byte{}, []TXInput{txin}, []TXOutput{txout})
	transaction.Hash = transaction.ComputeHash()
	return transaction, nil
}

// CreateBlockchain creates a new blockchain, evey adress in adresses
// is given the initial 
func NewBlockchain(addresses []string) *Blockchain {
	tx1 := NewCoinbaseTX(addresses[0], "Reward for " + addresses[0])
	tx2 := NewCoinbaseTX(addresses[1], "Reward for " + addresses[1])
	GBlock := NewGBlock([]*Transaction{tx1, tx2}, ZEROBITS)
	bc := NewBlockchainFromGB(GBlock)
	return bc
}

// creates a new blockchain given a valid genesis block
func NewBlockchainFromGB(genesis *Block) *Blockchain {
	return &Blockchain{nil, []Block{*genesis}}
}

func (bc *Blockchain) GetBalance(address string) int {
	balance := 0
	// 1. Calculating the amount credited (and adding it)
		for _, block := range bc.Chain {
			for _, transaction := range block.Transactions {
				for _, txOut := range transaction.TxOuts {
					if txOut.ScriptPubKey == address {
						balance += txOut.Value
					}
				}
			}
		}

	// 2. Caluating the amount debited (and substracting it)
	// - Get all transacations ids from which it (the current address) has credited money from
	outputTransactions := bc.getAllTransactionsByInputScriptSig(address)
	// - For each transaction id, loop for all transactions where this transaction id is the input's "scriptSig", and substract the value of the txIutput from the balance
	for _, transaction := range outputTransactions {
		for _, txOutput := range transaction.TxOuts {
			balance -= txOutput.Value
		}
	}
	return balance
}

/**
Return all transactions that have the given scriptSig in their input
*/
func (bc *Blockchain) getAllTransactionsByInputScriptSig(scriptSig string) []*Transaction {
	var transactions []*Transaction
	for _, block := range bc.Chain {
		for _, transaction := range block.Transactions {
			for _, txInput := range transaction.TxIns {
				if txInput.ScriptSig == scriptSig {
					transactions = append(transactions, transaction)
				}
			}
		}
	}
	return transactions
}
