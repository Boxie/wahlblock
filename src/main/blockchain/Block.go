package blockchain

import (
	"time"
)

type Block struct {
	Index int
	Timestamp time.Time
	Transactions []Transaction
	Proof int
	PreviousHash string
}


func (b Block) GetVotings() map[string]int{
	var votings map[string] int

	print("VOTINGS: ")
	print (votings)

	for _, transaction := range b.Transactions {
		if transaction.isValid(){
			votings[transaction.Voting] += 1
		}
	}
	return votings
}

func (b Block) isValid() bool{
	return true
}