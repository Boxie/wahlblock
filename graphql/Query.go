package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/boxie/wahlblock/blockchain"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"blockchain": &graphql.Field{
			Type: blockchainType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				//TODO add error handling
				pChain := blockchain.GetInstance()
				return pChain, nil
			},
		},
		//TODO ADD Nodes
	},
})

/* Blockchain type


 */

var blockchainType = graphql.NewObject(graphql.ObjectConfig{
	Name: "blockchainType",
	Fields: graphql.Fields{
		"chain": &graphql.Field{
			Type: chainType,
			Description: "Returns array of blocks in chain",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				//TODO add transaction
				//TODO add error handling
				if pBlockchain, ok := p.Source.( *blockchain.Blockchain); ok {
					return pBlockchain.Chain, nil
				}
				return nil, nil
			},
		},
		"currentTransactions": &graphql.Field{
			Type: graphql.NewList(transactionType),
			Args: pagingArguments,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {

				offset := p.Args["offset"].(int)
				first := p.Args["first"].(int)

				//validate offset and first
				if pBlockchain, ok := p.Source.( *blockchain.Blockchain); ok {

					start, end := calculatePaging(offset,first, len(pBlockchain.CurrentTransactions))

					return pBlockchain.CurrentTransactions[start:end], nil
				}
				return nil, nil
			},
		},
	},
})

var chainType = graphql.NewObject(graphql.ObjectConfig{
	Name: "chainType",
	Fields: graphql.Fields{
		"length": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				//TODO add error handling
				if chain, ok := p.Source.([]blockchain.Block); ok {
					return len(chain), nil
				}
				return nil, nil
			},
		},
		"block": &graphql.Field{
			Type: blockType,
			Args: indexArguments,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {

				index, err := p.Args["index"].(int)

				if err {
					//TODO Add error handling
				}

				if chain, ok := p.Source.([]blockchain.Block); ok {

					// Is index valid?
					if index - 1 > len(chain){
						//TODO add error message
						return nil, nil
					}
					return chain[index], nil
				}
				return nil, nil
			},
		},
		"blocks": &graphql.Field{
			Type: graphql.NewList(blockType),
			Args: pagingArguments,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {

				offset := p.Args["offset"].(int)
				first := p.Args["first"].(int)



				//TODO valid offset and first arguments


				if chain, ok := p.Source.([]blockchain.Block); ok {
					return chain[offset:first], nil
				}
				return nil, nil
			},
		},
	},
})

var blockType = graphql.NewObject(graphql.ObjectConfig{
	Name: "blockType",
	Fields: graphql.Fields{
		"index": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if block, ok := p.Source.(blockchain.Block); ok {
					return block.Index, nil
				}
				return nil, nil
			},
		},
		"timestamp": &graphql.Field{
			Type: graphql.DateTime,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if block, ok := p.Source.(blockchain.Block); ok {
					return block.Timestamp, nil
				}
				return nil, nil
			},
		},
		"transactions": &graphql.Field{
			Type: graphql.NewList(transactionType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if block, ok := p.Source.(blockchain.Block); ok {
					return block.Transactions, nil
				}
				return nil, nil
			},
		},
		"proof": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if block, ok := p.Source.(blockchain.Block); ok {
					return block.Proof, nil
				}
				return nil, nil
			},
		},
		"previousHash": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if block, ok := p.Source.(blockchain.Block); ok {
					return block.PreviousHash, nil
				}
				return nil, nil
			},
		},
	},
})


var transactionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "transactionType",
	Fields: graphql.Fields{
		"sender": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if pTransaction, ok := p.Source.( blockchain.Transaction); ok {
					return pTransaction.Sender, nil
				}
				return nil, nil
			},
		},
		"recipient": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if pTransaction, ok := p.Source.( blockchain.Transaction); ok {
					return pTransaction.Recipient, nil
				}
				return nil, nil
			},
		},
		"amount": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if pTransaction, ok := p.Source.( blockchain.Transaction); ok {
					return pTransaction.Amount, nil
				}
				return nil, nil
			},
		},
	},
})

/* Arguments

 */

var pagingArguments = graphql.FieldConfigArgument{
	"offset": &graphql.ArgumentConfig{
		Type:         graphql.Int,
		DefaultValue: 0,
	},
	"first": &graphql.ArgumentConfig{
			Type:         graphql.Int,
			DefaultValue: 10,
		},
}

var indexArguments = graphql.FieldConfigArgument{
	"index": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
}