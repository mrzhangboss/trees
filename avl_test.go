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

		//1 - 12
		{[]int{1, 2, 3, 4, 5}, "2-3,1-1,4-2,3-1,5-1"},
		{[]int{1, 2, 3, 4, 5, 6}, "4-3,2-2,1-1,3-1,5-2,6-1"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, "8-4,4-3,2-2,1-1,3-1,6-2,5-1,7-1,10-3,9-1,11-2,12-1"},
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

func TestAVLFind(t *testing.T) {
	var tests = []struct {
		values []int
		find   int
		r      bool
	}{
		{[]int{1}, 1, true},
		{[]int{1}, 2, false},
		{[]int{1, 2, 5, 9}, 1, true},
		{[]int{1, 2, 5, 9}, 5, true},
	}
	for _, test := range tests {
		tree := GetAVLTree(test.values[0])
		for _, v := range test.values[1:] {
			tree.Add(v)
		}
		if tree.Find(test.find) != test.r {
			t.Errorf("find error %s : %d", tree.Value(), test.find)
		}
	}
}

func TestAVLDelete(t *testing.T) {
	var tests = []struct {
		values []int
		d      int
		result string
	}{
		{[]int{10, 5, 15, 8}, 10, "8-2,5-1,15-1"},
		{[]int{10, 5, 15, 1}, 10, "5-2,1-1,15-1"},
		{[]int{10, 5}, 10, "5-1"},
		{[]int{10, 15}, 10, "15-1"},
	}
	for _, test := range tests {
		tree := GetAVLTree(test.values[0])
		for _, v := range test.values[1:] {
			tree.Add(v)
		}
		tree.Delete(test.d)
		if tree.Value() != test.result {
			t.Errorf("tree %s != %s error", test.result, tree.Value())
			continue
		}
	}
}
