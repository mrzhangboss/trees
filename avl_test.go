package trees

import (
	"fmt"
	"testing"
)

func tesForPointer(l *[]int) {
	n := append(*l, 10)
	n = make([]int, 0, 100)
	for i := 1; i < 100; i = i + 1 {
		n = append(n, 10)
		fmt.Println(i, len(n), cap(n), &n)
	}

	fmt.Println(n[2])
	s := (*l)[1:2]
	s = append(s, 5)
	fmt.Println(s[1])
	fmt.Println(n[2])
	fmt.Println(n[len(n)-1], len(n), cap(n))
	fmt.Println((*l)[len(*l)-1], len(*l), cap(*l))

}

func testListPointer(t *testing.T) {
	l := make([]int, 3, 10)
	tesForPointer(&l)
	if len(l) != 4 {
		t.Error(l)
	}
}

func TestAVLAdd(t *testing.T) {
	var tests = []struct {
		values []int
		result string
	}{
		{[]int{6, 3, 7, 1, 4, 2}, "3,1,2,6,4,7"},
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
