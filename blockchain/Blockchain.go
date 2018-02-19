package blockchain

import (
	"time"
	"encoding/json"
	"crypto/sha256"
	"encoding/hex"
	"sync"
	"fmt"
)

type Blockchain struct {
	Chain []Block
	CurrentTransactions []Transaction

}

/* Blockchain Singleton


 */

var instance *Blockchain
var once sync.Once

func GetInstance() *Blockchain{
	once.Do(func() {
		instance = &Blockchain{}
		instance.Chain = make([] Block, 0)
		instance.NewBlock(1,"None")
		fmt.Println("Instance: ",len(instance.Chain))
	})
	return instance
}



/*
Create a new block and add it to the chain
 */

func (bc *Blockchain) NewBlock(proof int, previousHash string) int{

	if len(previousHash) == 0 {
		previousHash = bc.Hash(bc.LastBlock())
	}

	block := Block {
		len(bc.Chain),
		time.Now(),
		bc.CurrentTransactions,
		proof,
		previousHash,
	}

	bc.CurrentTransactions = bc.CurrentTransactions[:0]
	bc.Chain = append(bc.Chain, block)

	return len(bc.Chain) -1
}

/*
Create a new transaction and add it to the chain
Returns index of newTransaction (this index needs to be mined)
 */

func (bc *Blockchain) NewTransaction(sender string,recipient string, amount int) int{
	transaction := Transaction{
		sender,
		recipient,
		amount,
	}
	bc.CurrentTransactions = append(bc.CurrentTransactions, transaction)

	//return index of
	return len(bc.CurrentTransactions) -1
}

/*
Create json of block and hash it
 */

func (bc *Blockchain) Hash(block Block) string{
	//TODO Add Hasher to blockchain struct
	hasher := sha256.New()
	blockString, _:= json.Marshal(block)
	hasher.Write([]byte (blockString))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (bc *Blockchain) LastBlock() Block{
	return bc.Chain[len(bc.Chain) - 1]
}

func (bc *Blockchain) ProofOfWork(lastProof int) int {
	proof := 0
	for !bc.validProof(lastProof,proof,bc.LastBlock().PreviousHash) {
		proof += 1
	}
	return proof
}

func (bc *Blockchain) validProof(lastProof int, proof int, previousHash string) bool{
	//TODO Add Hasher to blockchain struct
	hasher := sha256.New()

	//TODO CHECK previousHash
	guessString := string(lastProof*proof) + previousHash
	guess := []byte(guessString)
	hasher.Write(guess)
	guessHash := hex.EncodeToString(hasher.Sum(nil))
	//TODO add validation to blockchain struct
	return guessHash[:4] == "0000"
}

type Block struct {
	Index int
	Timestamp time.Time
	Transactions []Transaction
	Proof int
	PreviousHash string
}

type Transaction struct {
	Sender string
	Recipient string
	Amount int
}
