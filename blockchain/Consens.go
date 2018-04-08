package blockchain

import (
	"time"
	"github.com/shurcooL/graphql"
	"strconv"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"regexp"
)

type Consens struct {
	ActiveNodes map[string]Node
}

func (c Consens) GetCount() int{
	return len(c.ActiveNodes)
}

func (c Consens) NodeGetByHash(hash string) Node {
	return c.ActiveNodes[hash]
}

func (c Consens) GetNodes(start int, end int) []Node{
	var keys []string
	for k := range c.ActiveNodes{
		keys = append(keys, k)
	}
	keys = keys[start:end]

	var nodes []Node
	for _,k := range keys{
		nodes = append(nodes, c.ActiveNodes[k])
	}
	return nodes
}

func (c Consens) Add(n Node) bool{
	if validIP4(n.Host){
		if n.isActive(){
			n.RegisteredAt = time.Now()
			n.LastMessageAt = time.Now()
			c.ActiveNodes[n.GetHash()] = n
			return true
		}
	}
	return false
}

func (c Consens) AddAll (nodes []Node) {
	for _,n := range nodes {
		c.Add(n)
	}
}

type Node struct {
	//Schema string 	//TODO Add Schema https/http --> graphql enum
	Host string
	Port int
	Registrant string
	RegisteredAt time.Time
	LastMessageAt time.Time
}

//TODO ADD /graphql path to configfile

func (n Node) getAddress() string{
	return "http://" + n.Host + ":" + strconv.Itoa(n.Port) + "/graphql" //CHANGE HTTP:// to Schema
}

func (n Node) GetHash() string{
	hasher := sha256.New()
	hasher.Write([]byte (n.Host + ":" + strconv.Itoa(n.Port)))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (n Node) isActive() bool{
	var query struct {
		Status graphql.Boolean
	}

	client := graphql.NewClient(n.getAddress(), nil)
	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		// Handle error.
	}

	if query.Status == true {
		return true
	}
	return false
}

func validIP4(ipAddress string) bool {
	ipAddress = strings.Trim(ipAddress, " ")

	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	if re.MatchString(ipAddress) {
		return true
	}
	return false
}

func (c Consens) broadcastTransaction(t Transaction){
	var m struct {
		Blockchain struct {
			AddTransaction struct {
				Ballot graphql.String
				Voting graphql.String
			} `graphql:"addTransaction(ballot: $ballot, voting: $voting)"`
		}
	}

	v := map[string]interface{}{
		"ballot": t.Ballot,
		"voting": t.Voting,
	}

	for _,node := range c.ActiveNodes{
		client := graphql.NewClient(node.getAddress(), nil)

		err := client.Mutate(context.Background(), &m, v)
		if err != nil {
			println("Error")
		}
	}

}

func (n Node) getAllTransactionFromBlock(index int) []Transaction{
	var q struct {
		Blockchain struct {
			Chain struct {
				Block struct {
					Transactions []struct {
						Ballot graphql.String
						Voting graphql.String
					} `graphql:"transactions(offset: $offset, first: $first)"`
				} `graphql:"block(index: $index)"`
			}
		}
	}

	notEnd := true
	currentNumber := 0
	var transactions []Transaction
	client := graphql.NewClient(n.getAddress(), nil)

	for notEnd{
		v := map[string]interface{}{
			"offset": graphql.Int(currentNumber),
			"first": graphql.Int(100),
			"index": graphql.Int(index),
		}

		err := client.Query(context.Background(), &q, v)
		if err != nil {
			println("Error")
		}

		for _, qT := range q.Blockchain.Chain.Block.Transactions {
			t := Transaction {
				Ballot: string(qT.Ballot),
				Voting: string(qT.Voting),
			}
			transactions = append(transactions, t)
		}
		currentNumber += 100

		if len(transactions) < currentNumber{
			notEnd = false
		}
	}
	return transactions
}

func (c Consens) GetNodesWithLatestChain () []Node{
	var q struct {
		Blockchain struct {
			Chain struct {
				Length graphql.Int
				LastHash graphql.String
			}
		}
	}

	var copyNodes = make(map[string][]Node,0)
	var longestCount = 0

	for _, n := range c.ActiveNodes {
		client := graphql.NewClient(n.getAddress(), nil)
		err := client.Query(context.Background(), &q, nil)
		if err != nil {
			// Handle error.
		}

		if int(q.Blockchain.Chain.Length) > longestCount {
			longestCount = int(q.Blockchain.Chain.Length)
			for k := range copyNodes {
				delete(copyNodes, k)
			}
		}
		if int(q.Blockchain.Chain.Length) >= longestCount {
			hash := string(q.Blockchain.Chain.LastHash)
			copyNodes[hash] = append(copyNodes[hash],n)
		}
	}

	var total = 0
	var mostSupportedHash string
	var highestCount int

	for k,v := range copyNodes {
		total += len(v)
		if len(v) > highestCount {
			highestCount = len(v)
			mostSupportedHash = k
		}
	}

	if float32(highestCount / total) >= 0.51 {
		return copyNodes[mostSupportedHash]
	}
	return nil
}