package trees

import (
	"fmt"
	"strings"
)

type ATree struct {
	Val, Height int
	Left, Right *ATree
}

type AVLTree struct {
	root *ATree
}

func GetAVLTree(v int) *AVLTree {
	root := &ATree{Val: v, Height: 1}
	return &AVLTree{root}
}

func (t *AVLTree) Add(v int) {
	root := t.root
	fathers := []*ATree{}
	lefts := []bool{}
	for {
		if root != nil {
			if root.Val == v {
				break
			} else {
				if root.Val > v {
					if root.Left == nil {
						root.Left = &ATree{Val: v, Height: 1}
					} else {
						fathers = append(fathers, root)
						lefts = append(lefts, true)
						root = root.Left
					}
				} else {
					if root.Right == nil {
						root.Right = &ATree{Val: v, Height: 1}
					} else {
						fathers = append(fathers, root)
						lefts = append(lefts, false)
						root = root.Right
					}
				}
			}
		} else {
			fmt.Println("Tree is nil")
		}
	}
	for i := len(fathers) - 1; i >= 0; i-- {
		node := fathers[i]
		l, r, m, d := node.LRMD()
		if d > 1 {
			fmt.Printf("node %d -%d - %d", node.Val, v, i)
			r := balance(node, l, r)
			if i == 0 {
				t.root = r
			} else {
				isLeft := lefts[i-1]
				if isLeft {
					fathers[i-1].Left = r
				} else {
					fathers[i-1].Right = r
				}
			}
			node = r
		}
		l, r, m, d = node.LRMD()
		node.Height = m + 1

	}

}

func (node *ATree) LRMD() (l, r, m, d int) {
	// var l, r, m, d int
	if node.Left != nil {
		l = node.Left.Height
	} else {
		l = 0
	}
	if node.Right != nil {
		r = node.Right.Height
	} else {
		r = 0
	}
	if l > r {
		m = l
		d = l - r
	} else {
		m = r
		d = r - l
	}
	return
}

func (node *ATree) Update() {
	_, _, m, _ := node.LRMD()
	node.Height = m + 1
}

func leftBalance(father *ATree) *ATree {
	l, r, _, _ := father.Left.LRMD()
	root := father
	if l > r {
		left := root.Left
		root.Left, root.Left.Right = root.Left.Right, root
		root.Update()
		return left
	} else {
		right := father.Left.Right
		root.Left.Right, right.Left = right.Left, father.Left
		root.Left, right.Right = right.Right, father
		right.Left.Update()
		root.Update()
		return right

	}

}

func rightBalance(father *ATree) *ATree {
	l, r, _, _ := father.Right.LRMD()
	root := father
	if l > r {
		right := father.Right.Left
		root.Right.Left, right.Right = right.Right, father.Right
		root.Right, right.Left = right.Left, father
		right.Right.Update()
		root.Update()
		return right
	} else {
		left := root.Right
		root.Right, root.Right.Left = root.Right.Left, root
		root.Update()
		return left
	}
}

func balance(father *ATree, left, right int) *ATree {
	if left > right {
		return leftBalance(father)
	} else {
		return rightBalance(father)
	}

}

func (b *AVLTree) Value() string {
	list := []string{}
	FrontIndexTree(b.root, &list)
	return strings.Join(list, ",")

}

func FrontIndexTree(tree *ATree, list *[]string) {
	if tree == nil {
		return
	}
	v := fmt.Sprintf("%d-%d", tree.Val, tree.Height)
	*list = append(*list, v)
	FrontIndexTree(tree.Left, list)
	FrontIndexTree(tree.Right, list)

}
