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
					fathers = append(fathers, root)
					lefts = append(lefts, true)
					if root.Left == nil {
						root.Left = &ATree{Val: v, Height: 1}
					} else {
						root = root.Left
					}
				} else {
					fathers = append(fathers, root)
					lefts = append(lefts, false)
					if root.Right == nil {
						root.Right = &ATree{Val: v, Height: 1}
					} else {

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
	if l > r {
		left := father.Left
		father.Left, father.Left.Right = father.Left.Right, father
		father.Update()
		return left
	} else {
		right := father.Left.Right
		father.Left.Right, right.Left = right.Left, father.Left
		father.Left, right.Right = right.Right, father
		right.Left.Update()
		father.Update()
		return right

	}

}

func rightBalance(father *ATree) *ATree {
	l, r, _, _ := father.Right.LRMD()
	if l > r {
		right := father.Right.Left
		father.Right.Left, right.Right = right.Right, father.Right
		father.Right, right.Left = right.Left, father
		right.Right.Update()
		father.Update()
		return right
	} else {
		left := father.Right
		father.Right, father.Right.Left = father.Right.Left, father
		father.Update()
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

func (b *AVLTree) Find(v int) bool {
	root := b.root
	for {
		if root != nil {
			if root.Val == v {
				return true
			} else if root.Val > v {
				root = root.Left
			} else {
				root = root.Right
			}

		} else {
			return false
		}
	}

}

func (b *AVLTree) Delete(v int) {
	root := b.root
	fathers := []*ATree{}
	lefts := []bool{}
	for {
		if root != nil {
			if root.Val == v {
				break
			} else {
				if root.Val > v {
					if root.Left == nil {
						return
					} else {
						fathers = append(fathers, root)
						lefts = append(lefts, true)
						root = root.Left
					}
				} else {
					if root.Right == nil {
						return
					} else {
						fathers = append(fathers, root)
						lefts = append(lefts, false)
						root = root.Right
					}
				}
			}
		} else {
			fmt.Println("Tree is nil")
			return
		}
	}

	if len(fathers) == 0 {
		node := deleteAVL(root)
		node.Update()
		l, r, _, d := node.LRMD()
		if d > 1 {
			node = balance(node, l, r)
			node.Update()
		}
		b.root = node
	} else {
		updateFather(fathers[len(fathers)-1], deleteAVL(root), lefts[len(lefts)-1])

		for i := len(fathers) - 1; i >= 0; i-- {
			node := fathers[i]
			l, r, _, d := node.LRMD()
			if d > 1 {
				r := balance(node, l, r)
				if i == 0 {
					b.root = r
				} else {
					isLeft := lefts[i-1]
					updateFather(fathers[i-1], r, isLeft)
				}
				node = r
			}
			node.Update()
		}
	}

}

func updateFather(father, son *ATree, left bool) {
	if left {
		father.Left = son
	} else {
		father.Right = son
	}
}

func sonBalance(son, father *ATree, isLeft bool) {
	son.Update()
	l, r, _, d := son.LRMD()
	if d > 1 {
		node := balance(son, l, r)
		updateFather(father, node, isLeft)
	}
	father.Update()

}

func deleteAVL(root *ATree) *ATree {

	if root.Left == nil && root.Right == nil {
		root = nil
		return root
	}
	if root.Left == nil && root.Right != nil {
		return root.Right
	}
	if root.Left != nil && root.Right == nil {
		return root.Left
	}
	if root.Left.Right == nil {
		root.Left.Right = root.Right
		return root.Left
	}
	father := root.Left
	node := root.Left.Right
	fathers := []*ATree{root, father}
	lefts := []bool{true}
	for {
		if node.Right == nil {
			break
		} else {
			father = node
			fathers = append(fathers, father)
			lefts = append(lefts, false)
			node = node.Right
		}
	}
	node.Right = root.Right
	father.Right = node.Left
	node.Left = root.Left
	for i := len(fathers) - 1; i > 0; i-- {
		sonBalance(fathers[i], fathers[i-1], lefts[i-1])
	}
	return node

}
