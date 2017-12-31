package trees

import (
	"fmt"
	"strings"
)

type Tree struct {
	Val         int
	Left, Right *Tree
}

type BinarryTree struct {
	root *Tree
}

func Hello() {
	fmt.Println("hello world!")
	v := new(Tree)
	fmt.Println(v.Val)
}

func GetBinarryTree(v int) *BinarryTree {
	t := Tree{Val: v}
	tree := BinarryTree{&t}
	return &tree
}

func (b *BinarryTree) Add(v int) {
	root := b.root
	for {
		if root != nil {
			if root.Val == v {
				break
			} else {
				if root.Val > v {
					if root.Left == nil {
						root.Left = &Tree{Val: v}
					} else {
						root = root.Left
					}
				} else {
					if root.Right == nil {
						root.Right = &Tree{Val: v}
					} else {
						root = root.Right
					}
				}
			}
		} else {
			fmt.Println("Tree is nil", b.Value())
		}
	}
}

func (b *BinarryTree) Find(v int) bool {
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

func (b *BinarryTree) Delete(v int) {
	var root, father *Tree
	root = b.root
	isLeft := true
	if root.Val == v {
		b.root = delete(root)
		return
	}
	for {
		if root != nil {
			if root.Val == v {
				if isLeft {
					father.Left = delete(root)
				} else {
					father.Right = delete(root)
				}
				break
			} else if root.Val > v {
				father = root
				root = root.Left
				isLeft = true
			} else {
				father = root
				root = root.Right
				isLeft = false
			}

		} else {
			fmt.Printf("can find %d", v)
			return
		}
	}
}

func delete(root *Tree) *Tree {
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
	if root.Right.Left == nil {
		root.Right.Left = root.Left
		return root.Right
	}
	father := root.Right
	node := root.Right.Left
	for {
		if node.Left == nil {
			break
		} else {
			father = node
			node = node.Left
		}
	}
	node.Left = root.Left
	father.Left = node.Right
	node.Right = root.Right
	return node
}

func (b *BinarryTree) Value() string {
	list := []string{}
	BackIndexTree(b.root, &list)
	return strings.Join(list, ",")
}

func BackIndexTree(tree *Tree, list *[]string) {
	if tree == nil {
		return
	}
	BackIndexTree(tree.Left, list)
	v := fmt.Sprintf("%d", tree.Val)
	*list = append(*list, v)
	BackIndexTree(tree.Right, list)

}
