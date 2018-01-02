package trees

import (
	_ "fmt"
	"testing"
)

func TestAVLAdd(t *testing.T) {
	var tests = []struct {
		values []int
		result string
	}{
		//LL
		{[]int{4, 3, 2}, "3-2,2-1,4-1"},
		{[]int{6, 3, 7, 1, 4, 2}, "3-3,1-2,2-1,6-2,4-1,7-1"},
		{[]int{10, 5, 15, 4, 3}, "10-3,4-2,3-1,5-1,15-1"},

		//LR
		{[]int{4, 2, 3}, "3-2,2-1,4-1"},
		{[]int{6, 2, 7, 1, 4, 5}, "4-3,2-2,1-1,6-2,5-1,7-1"},
		{[]int{6, 2, 7, 1, 4, 3}, "4-3,2-2,1-1,3-1,6-2,7-1"},

		// RL
		{[]int{1, 3, 2}, "2-2,1-1,3-1"},
		{[]int{2, 1, 6, 4, 8, 3}, "4-3,2-2,1-1,3-1,6-2,8-1"},
		{[]int{2, 1, 6, 4, 8, 5}, "4-3,2-2,1-1,6-2,5-1,8-1"},

		//RR
		{[]int{1, 2, 3}, "2-2,1-1,3-1"},
		{[]int{2, 1, 6, 4, 8, 10}, "6-3,2-2,1-1,4-1,8-2,10-1"},
		{[]int{2, 1, 6, 4, 8, 7}, "6-3,2-2,1-1,4-1,8-2,7-1"},
	}
	for _, test := range tests {
		tree := GetAVLTree(test.values[0])
		for _, v := range test.values[1:] {
			tree.Add(v)
		}
		if tree.Value() != test.result {
			t.Errorf("tree %s != %s error", test.result, tree.Value())
			continue
		}
	}
}
