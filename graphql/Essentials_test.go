package graphql

import (
	"testing"
	"strconv"
)

func TestGraphQLEssentials(t *testing.T) {
	t.Run("Check essentials calculate paging", func (t *testing.T){
		cases := []struct {
			Offset int
			First int
			Length int
			ExpectedOffset int
			ExpectedFirst int
		}{
			{0,10, 5, 0, 4},
			{7, 10, 3, 0,0},
			{1, 3, 10, 1,3},
			{10, 20, 100, 10,29},
			{0, 10, 1, 0,0},
		}

		for index,c := range cases {
			offset, first := calculatePaging(c.Offset,c.First,c.Length)

			if offset != c.ExpectedOffset {
				t.Log("Paging calulation offset failed Case Index: " + strconv.Itoa(index))
				t.Fail()
			}

			if first != c.ExpectedFirst {
				t.Log("Paging calulation first failed Case Index: "  + strconv.Itoa(index))
				t.Fail()
			}
		}
	})
}

