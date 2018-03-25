package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/boxie/wahlblock/blockchain"
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

					index, err := blockchain.Mine()

					if err != nil{
						panic(err)
					}

					return blockchain.Chain[index], err

				}
				return nil, nil
			},
		},
	},
})


var transactionArguments = graphql.FieldConfigArgument{
	"ballot": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"voting": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}