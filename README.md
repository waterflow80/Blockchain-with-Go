# Bloockchain Implementation with Go
This is a basic blockchain implementation using Go. Some of the functions and methods were already provided, and some others were developed by us.

## Design and Data Structures
![Screenshot from 2024-01-24 15-48-14](https://github.com/waterflow80/Blockchain-with-Go/assets/82417779/c37d38f4-c02b-40ae-b0f3-8abeae566b7b)
### Blockchain
- The blockchain system is composed of a set of blocks
- These blocks are linked together using their hash
- The first block is called **Genesis Block**, it doesn't have previous hash
- This is how the blockchain is represented in Go:
  ```golang
  type Blockchain struct {
	  GHash []byte 
	  Chain []Block
  }
  ```
### Block
A block is composed of the following elements:
- **Hash**: the hash of the entire block
- **PrevBlockHash**: the hash of the previous block in the blockchain
- **Transactions**: a set of transactions made by some actors of the blockchain
- **Nonce**: an integer used by miners to satisfy some block conditions. Eg: the hash should start with N zeros
- **Timestamp**: Unix time, indicates the time the block was added to the blockchain
- Here's how the Block is represented in Go:
  ```golang
  type Block struct {
	  Hash          []byte
	  PrevBlockHash []byte
	  Transactions  []*Transaction
	  Timestamp     int64
	  Nonce         int
  }
  ```
### Transaction
- A transaction represent a transfer from one actor to another inside the blockchain.
- A transaction can have multiple **inputs** and **outputs**.
- The **TXInput** contains information about the rights that allow the sender to send the desired output (**ScritpSig**), plus other transaction info.
- The **TXOutput** contains information about how the output value can be redeemed; **ScriptPubKey** (the receiver's pubKey).
- To learn more about transactions, you can watch this lecture: [MIT MAS.S62 Cryptocurrency Engineering and Design, Spring 2018](https://www.youtube.com/watch?v=VT2o4KCEbes&list=PLUl4u3cNGP61KHzhg3JIJdK08JLSlcLId&index=4)
- This is how the Transaction, TXInput, and TXOutput are represented in Go:
  ```golang
  type Transaction struct {
	  Hash   []byte
	  TxIns  []TXInput
	  TxOuts []TXOutput
  }
  ```
  ```golang
  type TXInput struct {
	  Txid      []byte
	  Vout      int
 	  ScriptSig string
  }
  ```
  ```golang
  type TXOutput struct {
	  Value        int
	  ScriptPubKey string
  }
  ```
## Implementation
In this section, we'll walk through some of the functions and methods that we've implemented:
- `func NewCoinbaseTX(...) *Transaction {..}`: create a new coinbaseTX that doesn't have an input, and the receiver will get a reward defined by the system that will be stored in the TXOutput's value.
- `func NewBlockchain(...) *Blockchain {..}`: create a new blockchain that contains initially the provided adresses (actors).
- `func (bc *Blockchain) AddBlock(...) {..}`: add a block with the given transactions to the blockchain. If the blockchain is empty, then the block will be a Genesis block (first block), otherwise, add the new block to the end of the chain.
- `func  (bc *Blockchain) NewTransfertTX(...) (*Transaction,error) {..}`: create and add a new transaction that represents the transfer of the given amount, from the given sender (scriptSig) to the given receiver (ScriptPubKey).
- `func (bc *Blockchain) GetBalance(...) int {..}`: return the current balance of the given address (actor). First look for transactions where the acotr has been the output/receiver of certain amounts, and add those amounts to the balance. Then, look for the transactions where the actor has been the sender (scriptSig), and substract that amount from the balance.
- `func BlockchainToMap(...) map[string]interface{} {..}`: convert the given blockchain object into a map using the **github.com/mitchellh/mapstructure** library to make it esaier to parse and convert nested objects.
- `func SaveBlockchain(...) error {..}`: transform the given blockchain into a map, then into a json object, and finally save it to the disk.
## Dependencies
- github.com/mitchellh/mapstructure
