package Blockchain

import (
	"testing"
	"github.com/boxie/wahlblock/src/main/blockchain"
)

func TestBlockchain(t *testing.T) {
	var bc = blockchain.GetInstance()
	var blockIndex int

	t.Run("Add transaction", func(t *testing.T) {
		cases := []struct {
			Ballot string
			Voting string
			Expected bool
		}{
			{"2735abc040de351e9e8cfb802e6575016bc66a8e2e0ae047dbb193b1274b285d","Yes", true},
			{"9fa5b12c172c9ddf58763351cdb8b44da0f9dc3d25e90597186418a2e643861c","Yes", true},
			{"65a0e96b9b3b55d12d80da65c24d49659f2816259bbb06ff642710d17a42a1df","No", true},
			{"68fbd524d5679e62da4726558033a911a41f3756dd3c0e9105b75620e22d891e","Yes", true},
			{"807d59f6af5281e87b88cf2a30e334e92b5a3d6b71bfc7aa6870df2115cd4bd4","Yes", true},
			{"e5cf37d24ec5bda22d23c03340cdc3422bf4177c98af11deba7c1d3d8de0f7d0","Yes", true},
			{"2ac36eafedc747db6984ecff2747e706a343c2ea45ded929e8719e0c4de1a25b","No", true},
			{"e06cb22f3f12b308db3e3a1820a2a025ca162b0b4b15ec9004f2928f1ce61564","No", true},
		}

		var index = 0

		for _, c := range cases {
			bc.NewTransaction(
				c.Ballot,
				c.Voting,
			)
			index += 1
			if bc.GetPendingTransactionCount() != index {
				t.Fail()
			}
		}
		t.Log("Successfully added four votings (3 Yes, 1 No)")
	})

	t.Run("Mining block", func(t *testing.T) {
		var index = bc.GetPendingTransactionCount()

		blockIndex,_ = bc.Mine()


		if bc.GetPendingTransactionCount() != 0 {
			t.Fail()
		}
		newBlock := bc.Chain[blockIndex]
		if newBlock.GetTransactionCount() != index {
			t.Fail()
		}
		t.Log("Successfully mined block")
	})

	t.Run("Get block votings", func(t *testing.T) {
		cases := []struct {
			Vote string
			Expected int
		}{
			{"Yes", 5},
			{"No", 3},
		}

		var votings = bc.Chain[blockIndex].GetVotings()

		for _, c := range cases {
			if votings[c.Vote] != c.Expected {
				t.Fail()
			}
		}
	})

	//TODO CHECK BLOCKCHAIN STATS

	t.Run("Check possibilities", func(t *testing.T) {
		cases := []struct {
			Vote string
			Expected bool
		}{
			{"Yes", true},
			{"YeS", false},
			{"No", true},
			{"SA", false},
			{"Qass", false},
			{"asdf", false},
			{"YES", false},
			{"NO", false},
		}

		var possibilities []string = bc.GetPossibilities()

		for _,c := range cases {
			found := false
			for _, value := range possibilities {
				if value == c.Vote {
					found = true
				}
			}

			if found != c.Expected {
				t.Fail()
			}
		}
	})







}