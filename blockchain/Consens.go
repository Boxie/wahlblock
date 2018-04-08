package blockchain

import "time"

type Consens struct {
	ActiveNodes map[string]Node
}

func (c Consens) GetCount() int{
	return len(c.ActiveNodes)
}

func (c Consens) NodeGetByHost(host string) Node {
	return c.ActiveNodes[host]
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

func (c Consens) Add(n Node){
	c.ActiveNodes[n.Host] = n
}

type Node struct {
	Host string
	Port int
	Registrant string
	RegisteredAt time.Time
	LastMessageAt time.Time
}