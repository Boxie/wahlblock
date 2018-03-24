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
	Chain               []Block
	PendingTransactions []Transaction

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
	Function

		New Block

	Description

		Create a new block and add it to the chain
		Takes all pendingTransactions and add it to the new block

	Parameter

		int		proof			proof
		string	previousHash	hash of the previous block of blockchain

 */

func (bc *Blockchain) NewBlock(proof int, previousHash string) int{

	if len(previousHash) == 0 {
		previousHash = bc.Hash(bc.LastBlock())
	}

	block := Block {
		len(bc.Chain),
		time.Now(),
		bc.PendingTransactions,
		proof,
		previousHash,
	}

	bc.PendingTransactions = bc.PendingTransactions[:0]
	bc.Chain = append(bc.Chain, block)

	return len(bc.Chain) -1
}

/**
	Function

		New Transaction

	Description

		Create a new transaction and add it to the pendingTransactions list
		Yet pending transactions are not permanently added to the blockchain. To add a transaction permanently to
		Returns index of newTransaction (this index needs to be mined)

	Parameter

		string 	sender		ID of the sender
		string 	recipient	ID of the recipient
		int 	amount		amount of coins to transfer

	Return

		int		index of added transaction in pending transaction list

	TODO Add transaction validation
**/

func (bc *Blockchain) NewTransaction(ballot string, voting string) int{
	transaction := Transaction{
		Ballot: ballot,
		Voting: voting,
		Timestamp: time.Now(),
	}
	bc.PendingTransactions = append(bc.PendingTransactions, transaction)

	//return index of
	return len(bc.PendingTransactions) -1
}

/*
	Function

		Hash

	Description

		Creates a hash of a json formatted block by first creating a json file out of the given block and second
		hashing the json file

	Parameter

		Block	block	block to be hashed

	Return

		string	hash of the given block
 */

func (bc *Blockchain) Hash(block Block) string{
	//TODO Add Hasher to blockchain struct
	hasher := sha256.New()
	blockString, _:= json.Marshal(block)
	hasher.Write([]byte (blockString))
	return hex.EncodeToString(hasher.Sum(nil))
}

/*
	Function

		LastBlock

	Description

		get and return the last block in blockchain

	Return

		Block	last Block in blockchain

 */

func (bc *Blockchain) LastBlock() Block{
	return bc.Chain[len(bc.Chain) - 1]
}

/*
	Function

		ProofOfWork

	Description

		Mining algorithm to find the next valid proof
		begins at 0 and last until valid proof is found

	Parameter

		int	lastProof	proof of last valid block of blockchain

	Return
 */

func (bc *Blockchain) ProofOfWork(lastProof int) int {
	proof := 0
	for !bc.validProof(lastProof,proof,bc.LastBlock().PreviousHash) {
		proof += 1
	}
	return proof
}

/*
	Function

		validProof

	Description

		Interact with function ProofOfWork. Checks if the given proof is valid
		Try to find a hash based on concatenation of the multiplication of two integers and the previous hash
		which has at least four zeros at the end

	Parameter

		int		lastProof		proof of the last valid block in blockchain
		int		proof			proof to check
		string	previousHash

	Return

		bool	returns true if generated hash has at least four zeros at the end
 */

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

func (bc *Blockchain) Mine () (int, error){

	lastBlock := bc.LastBlock()
	lastProof := lastBlock.Proof

	proof := bc.ProofOfWork(lastProof)

	//TODO Maybe reward miner?

	previousHash := bc.Hash(lastBlock)
	index := bc.NewBlock(proof, previousHash)

	return bc.Chain[index].Index, nil
}

/*
	Function

		getVotings

	Description

		Get a map of all votings in blockchain. Does not count pending transactions

	Return

		map[string]int	returns a map with the voting term as key and the number of votes as value
 */

 func (bc *Blockchain) getVotings() map[string] int {
	 var votings = make (map[string] int)
	 for _, block := range bc.Chain {
		 if block.isValid(){
			 for key, value := range block.GetVotings(){
			 	votings[key] += value
			 }
		 }
	 }
	 return votings
 }

 func (bc *Blockchain) GetBlockCount() int{
 	return len(bc.Chain)
 }

 func (bc *Blockchain) GetPendingTransactionCount() int {
 	return len(bc.PendingTransactions)
 }

 func (bc *Blockchain) GetPossibilities() [] string{

 	var votings =  bc.getVotings()
 	var keys []string;
 	for key,_ := range(votings){
 		keys = append(keys,key)
	}
	return keys
 }

func (bc *Blockchain) GetVotingCount() [] int{
	var votings =  bc.getVotings()
	var values []int;
	for _,value := range(votings){
		values = append(values,value)
	}
	return values
}

 func (bc *Blockchain) averageTransactionPerBlock() float32 {
 	return 0
 }