package blockchain

import (
	"time"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

type Block struct {
	Index int
	Timestamp time.Time
	Transactions []Transaction
	Proof int
	PreviousHash string
}


func (b Block) GetVotings() map[string]int{
	var votings = make(map[string] int)

	for _, transaction := range b.Transactions {
		if transaction.isValid(){
			votings[transaction.Voting] += 1
		}
	}
	return votings
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

func (b Block) Hash() string{
	//TODO Add Hasher to blockchain struct
	hasher := sha256.New()
	blockString, _:= json.Marshal(b)
	hasher.Write([]byte (blockString))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (b Block) isValid() bool{
	return true
}

func (b Block) GetTransactionCount() int{
	return len(b.Transactions)
}