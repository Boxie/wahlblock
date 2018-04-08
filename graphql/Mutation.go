package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/boxie/wahlblock/blockchain"
	"time"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"blockchain": &graphql.Field {
			Type: blockchainMutationType,
			Description: "Blockchain mutation",
			Resolve: func(p graphql.ResolveParams) (interface{}, error){

				//TODO Error handling

				var chain = blockchain.GetSession().Blockchain
				return chain, nil
			},
		},
		"consens": &graphql.Field {
			Type: consensMutationType,
			Description: "Consens mutation",
			Resolve: func(p graphql.ResolveParams) (interface{}, error){

				//TODO Error handling

				var consens = blockchain.GetSession().Consens
				return consens, nil
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

var consensMutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "consensMutationType",
	Fields: graphql.Fields{
		"register": &graphql.Field{
			Type: graphql.Boolean,
			Args: nodeArguments,
			Resolve: func(p graphql.ResolveParams) (interface {}, error) {

				host := p.Args["host"].(string)
				port := p.Args["port"].(int)
				registrant := p.Args["registrant"].(string)

				if consens, ok := p.Source.( *blockchain.Consens); ok {

					node := blockchain.Node{
						Host: host,
						Port: port,
						Registrant: registrant,
						RegisteredAt: time.Now(),
					}

					success := consens.Add(node)

					return success , nil

				}
				return false, nil
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

var nodeArguments = graphql.FieldConfigArgument{
	"host": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"port": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"registrant": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}