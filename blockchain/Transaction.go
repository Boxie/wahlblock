package blockchain

import "time"

type Transaction struct {
	Ballot string
	Voting string
	Timestamp time.Time
}

/*
	Function

		isValid

	Description

		Checks if transaction is valid based on the ballot id. Performs API Request to validate ballot id

	Return

		bool	return true if ballot id is valid

	TODO API Request
 */

func (transaction Transaction) isValid() bool{
	return true
}