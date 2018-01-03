package trees

import (
	"testing"
)

func TestRBTAdd(t *testing.T) {
	var tests = []struct {
		values []int
		result string
	}{
		//1 - 12
		{[]int{1, 2}, "1-0,2-1"},
		{[]int{1, 2, 3}, "2-0,1-1,3-1"},
		{[]int{1, 2, 3, 4, 5}, "2-0,1-0,4-0,3-1,5-1"},
		{[]int{1, 2, 3, 4, 5, 6}, "2-0,1-0,4-1,3-0,5-0,6-1"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, "4-0,2-0,1-0,3-0,8-0,6-1,5-0,7-0,10-1,9-0,11-0,12-1"},
	}
	for _, test := range tests {
		tree := GetRBTree(test.values[0])
		for _, v := range test.values[1:] {
			tree.Add(v)
		}
		if tree.Value() != test.result {
			t.Errorf("tree %s != %s error", test.result, tree.Value())
			continue
		}
	}
}
