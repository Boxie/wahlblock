package blockchain

import (
	"sync"
)

var instance *Session
var once sync.Once

type Session struct {
	Blockchain *Blockchain
	Consens *Consens
}

func GetSession() *Session{
	once.Do(func() {
		instance = &Session{
			Blockchain: &Blockchain{
				Chain: make([] Block, 0),
				PendingTransactions: make([] Transaction, 0),
			},
			Consens: &Consens{
				ActiveNodes: make(map[string]Node, 0),
			},
		}
		instance.Blockchain.NewBlock(1,"None")
		instance.addCommonServer()
	})
	return instance
}

func (s Session) addCommonServer(){

	nodes := []Node{
		{Host: "5.196.97.83", Port: 3000, Registrant: "Common server"},
	}

	s.Consens.AddAll(nodes)

}