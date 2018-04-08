package blockchain

import (
	"time"
	"github.com/shurcooL/graphql"
	"strconv"
	"context"
	"crypto/sha256"
	"encoding/hex"
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
	keys := []string{}
	for k := range c.ActiveNodes{
		keys = append(keys, k)
	}
	keys = keys[start:end]

	nodes := []Node{}
	for _,k := range keys{
		nodes = append(nodes, c.ActiveNodes[k])
	}
	return nodes
}

func (c Consens) Add(n Node) bool{
	if n.isActive(){
		c.ActiveNodes[n.GetHash()] = n
		return true
	}
	return false
}

type Node struct {
	Host string
	Port int
	Registrant string
	RegisteredAt time.Time
	LastMessageAt time.Time
}

func (n Node) getAddress() string{
	return n.Host + ":" + strconv.Itoa(n.Port)
}

func (n Node) GetHash() string{
	hasher := sha256.New()
	hasher.Write([]byte (n.Host + ":" + strconv.Itoa(n.Port)))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (n Node) isActive() bool{
	var query struct {
		status graphql.Boolean
	}

	client := graphql.NewClient(n.getAddress(), nil)
	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		// Handle error.
	}

	if query.status == true {
		return true
	}
	return false
}