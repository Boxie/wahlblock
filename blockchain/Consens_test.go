package blockchain

import (
"testing"

	"time"
	"fmt"
)

func TestConsens(t *testing.T) {
	con := GetSession().Consens

	cases := []struct {
		Host string
		Port int
		Registrant string
	}{
		{"135.133.60.201",3000, "localhost"},
		{"44.1.97.37",3000, "localhost"},
		{"140.89.159.178",3000, "localhost"},
		{"153.154.14.144",3000, "localhost"},
		{"35.84.251.157",3000, "localhost"},
		{"239.23.98.226",3000, "localhost"},
		{"159.87.176.160",3000, "localhost"},
		{"131.223.148.231",3000, "localhost"},
	}

	t.Run("Add nodes to consens", func(t *testing.T) {

		for _,c := range cases {
			node := Node{
				Host: c.Host,
				Port: c.Port,
				Registrant: c.Registrant,
				RegisteredAt: time.Now(),
			}
			con.Add(node)
		}

		if con.GetCount() != len(cases){
			t.Error("Consens Nodes does not equal case count")
		}
		t.Log("Finished adding nodes to consens")
	})

	t.Run("Get specific nodes by host string", func (t *testing.T){
		for _,c := range cases {
			node := con.NodeGetByHash(c.Host)
			if node.Host != c.Host || node.Port != c.Port || node.Registrant != c.Registrant{
				t.Error("Get specific nodes by host failed")
			}
		}
		t.Log("Finished get specific nodes")
	})

	t.Run("Get all transactions from node", func(t *testing.T){
		n := Node{
			Host: "5.196.97.83",
			Port: 3000,
		}

		transactions := n.getAllTransactionFromBlock(1)

		for transaction := range transactions {
			println(transaction)
		}
	})

	t.Run("Get nodes with latest chain", func(t *testing.T){

		nodes := con.GetNodesWithLatestChain()

		for _,node := range nodes {
			fmt.Println(node.Host)
		}
	})
}
