package graphql

import (
	"github.com/graphql-go/graphql"
	"fmt"
	"github.com/boxie/wahlblock/src/main/blockchain"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"blockchain": &graphql.Field {
			Type: blockchainMutationType,
			Description: "Blockchain mutation",
			Resolve: func(p graphql.ResolveParams) (interface{}, error){

				//TODO Error handling

				var chain = blockchain.GetInstance()
				return chain, nil
			},
		},
	},
})

var blockchainMutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "blockchainMutationType",
	Fields: graphql.Fields{
		"transactionAdd": &graphql.Field{
			Type: transactionType,
			Args: transactionArguments,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {

				ballot := p.Args["ballot"].(string)
				voting := p.Args["voting"].(string)

				if blockchain, ok := p.Source.( *blockchain.Blockchain); ok {

					index := blockchain.NewTransaction(ballot, voting)
					return blockchain.PendingTransactions[index], nil

				}
				return nil, nil


			},
		},
		"mine": &graphql.Field{
			Type: blockType,
			Resolve: func(p graphql.ResolveParams) (interface {}, error) {

				if blockchain, ok := p.Source.( *blockchain.Blockchain); ok {

					// start mining

					// TODO errorHandling
					fmt.Println(len(blockchain.Chain))
					lastBlock := blockchain.LastBlock()
					lastProof := lastBlock.Proof

					proof := blockchain.ProofOfWork(lastProof)

					//TODO Maybe reward miner?

					// Forge new Block bz adding it to the chain

					previousHash := blockchain.Hash(lastBlock)
					index := blockchain.NewBlock(proof, previousHash)

					return blockchain.Chain[index], nil

				}
				return nil, nil
			},
		},
	},
})



/* inputs

var transactionInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "transactionInputType",
		Fields: graphql.InputObjectConfigFieldMap{
			"sender": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"recipient": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"amount": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
		},
	},
)

*/

/*
arguments
 */

var transactionArguments = graphql.FieldConfigArgument{
	"ballot": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"voting": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}