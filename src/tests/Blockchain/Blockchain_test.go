package Blockchain

import (
	"testing"
	bc "github.com/boxie/wahlblock/src/main/blockchain"
	"reflect"
)

func TestBlockchain(t *testing.T) {
	t.Run("Getting blockchain instance", func(t *testing.T) {
		var bc = bc.GetInstance()

		if reflect.TypeOf(bc).Kind() == {

		}
	})

}