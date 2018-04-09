package blockchain

import (
	"time"
	"crypto/sha256"
	"encoding/hex"
	"sort"
	log "github.com/sirupsen/logrus"
)

type Blockchain struct {
	Chain               []Block
	PendingTransactions []Transaction

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
		previousHash = bc.LastBlock().Hash()
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

	log.WithFields(log.Fields{
		"index": block.Index,
		"timestamp": block.Timestamp,
	}).Info("Added block to chain")

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

	if transaction.isValid() {
		bc.PendingTransactions = append(bc.PendingTransactions, transaction)

		//TODO Refactoring Session Call
		GetSession().Consens.broadcastTransaction(transaction)

		log.WithFields(log.Fields{
			"ballot": transaction.Ballot,
			"voting": transaction.Voting,
		}).Info("Added transaction to pending transaction")

		return len(bc.PendingTransactions) -1
	}

	//TODO add Error
	return 0
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
	return guessHash[:2] == "00"
}

func (bc *Blockchain) Mine () (int, error){

	lastBlock := bc.LastBlock()
	lastProof := lastBlock.Proof

	proof := bc.ProofOfWork(lastProof)

	previousHash := lastBlock.Hash()
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

	 keys := make([]string, 0, len(votings))
	 for key := range votings {
		 keys = append(keys, key)
	 }
	 sort.Strings(keys) //sort by key

	 return keys
 }

func (bc *Blockchain) GetVotingCount() [] int{

	var votings =  bc.getVotings()

	keys := make([]string, 0, len(votings))
	for key := range votings {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var values []int
	for _,key := range keys{
		values = append(values,votings[key])
	}
	return values
}