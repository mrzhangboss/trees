package trees

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var tests = []struct {
		values []int
		result string
	}{
		{[]int{4}, "4"},
		{[]int{1, 1, 1, 1}, "1"},
		{[]int{1, 2, 3, 4}, "1,2,3,4"},
		{[]int{1, 4, 3, 2}, "1,2,3,4"},
		{[]int{4, 3, 2, 1}, "1,2,3,4"},
	}
	for _, test := range tests {
		tree := GetBinarryTree(test.values[0])
		for _, v := range test.values[1:] {
			tree.Add(v)
		}
		if tree.Value() != test.result {
			t.Errorf("tree %s != %s error", test.result, tree.Value())
			continue
		}
	}
}

func TestFind(t *testing.T) {
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
		tree := GetBinarryTree(test.values[0])
		for _, v := range test.values[1:] {
			tree.Add(v)
		}
		if tree.Find(test.find) != test.r {
			t.Errorf("find error %s : %d", tree.Value(), test.find)
		}
	}
}

func TestDelete(t *testing.T) {
	var tests = []struct {
		values []int
		d      int
		result string
	}{
		{[]int{4, 8, 6, 10}, 8, "4,6,10"},
		{[]int{10, 5, 3, 6}, 5, "3,6,10"},
		{[]int{10, 10, 10, 5, 15, 18, 16, 20}, 15, "5,10,16,18,20"},
		{[]int{10, 5, 2, 6, 15, 13, 18, 19, 17}, 13, "2,5,6,10,15,17,18,19"},
		{[]int{10, 5, 2, 6, 15, 13, 12, 14, 18, 19, 17}, 12, "2,5,6,10,13,14,15,17,18,19"},
	}
	for _, test := range tests {
		tree := GetBinarryTree(test.values[0])
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
