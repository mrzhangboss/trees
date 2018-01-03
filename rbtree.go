package trees

import (
	"fmt"
	"strings"
)

type RBTree struct {
	Color       bool // Red true Black false
	Val         int
	Left, Right *RBTree
}

type RBT struct {
	root *RBTree
}

func GetRBTree(v int) *RBT {
	return &RBT{&RBTree{Color: false, Val: v}}
}

func (r *RBT) Add(v int) {
	var node *RBTree
	root := r.root
	fathers := []*RBTree{}
	lefts := []bool{}
	for {
		if root != nil {
			if root.Val == v {
				return
			} else {
				if root.Val > v {
					fathers = append(fathers, root)
					lefts = append(lefts, true)
					if root.Left == nil {
						node = &RBTree{Val: v, Color: true}
						root.Left = node
						break
					} else {

						root = root.Left
					}
				} else {
					fathers = append(fathers, root)
					lefts = append(lefts, false)
					if root.Right == nil {
						node = &RBTree{Val: v, Color: true}
						root.Right = node
						break
					} else {

						root = root.Right
					}
				}
			}
		} else {
			fmt.Println("Tree is nil")
			return
		}
	}
	root = fixFatherBalance(node, fathers, lefts)
	if root != nil {
		r.root = root
	}

}

func uncle(fathers []*RBTree, lefts []bool) *RBTree {
	grandfather := fathers[len(fathers)-2]
	isFatherLeft := lefts[len(lefts)-2]
	if isFatherLeft {
		return grandfather.Right
	} else {
		return grandfather.Left
	}
}

func ll(root *RBTree) *RBTree {
	left := root.Left
	root.Left, root.Left.Right = root.Left.Right, root
	return left
}

func lr(father *RBTree) *RBTree {
	right := father.Left.Right
	father.Left.Right, right.Left = right.Left, father.Left
	father.Left, right.Right = right.Right, father
	return right
}

func rl(father *RBTree) *RBTree {
	right := father.Right.Left
	father.Right.Left, right.Right = right.Right, father.Right
	father.Right, right.Left = right.Left, father
	return right
}

func rr(father *RBTree) *RBTree {
	left := father.Right
	father.Right, father.Right.Left = father.Right.Left, father
	return left
}

func rbBalance(node *RBTree, fathers []*RBTree, lefts []bool) *RBTree {
	var root *RBTree
	n := len(lefts)
	fathers[n-2].Color = true
	if lefts[n-1] && lefts[n-1] {
		root = ll(fathers[n-2])
	} else if lefts[n-1] == false && lefts[n-2] {
		root = lr(fathers[n-2])
	} else if lefts[n-1] && lefts[n-2] == false {
		root = rl(fathers[n-2])
	} else {
		root = rr(fathers[n-2])
	}
	root.Color = false
	if n > 2 {
		if lefts[n-3] {
			fathers[n-3].Left = root
		} else {
			fathers[n-3].Right = root
		}
		return nil
	} else {
		return root
	}

}

func fixFatherBalance(node *RBTree, fathers []*RBTree, lefts []bool) *RBTree {
	if len(fathers) == 0 {
		node.Color = false
		return nil
	} else if fathers[len(fathers)-1].Color == false {
		return nil
	} else if len(fathers) >= 2 && uncle(fathers, lefts) != nil && uncle(fathers, lefts).Color == true {
		fathers[len(fathers)-1].Color = false
		uncle(fathers, lefts).Color = false
		fathers[len(fathers)-2].Color = true
		return fixFatherBalance(fathers[len(fathers)-2], fathers[:len(fathers)-2], lefts[:len(lefts)-2])
	} else {
		return rbBalance(node, fathers, lefts)
	}

}

func (r *RBT) Value() string {
	list := []string{}
	RBFrontIndexTree(r.root, &list)
	return strings.Join(list, ",")
}

func RBFrontIndexTree(tree *RBTree, list *[]string) {
	if tree == nil {
		return
	}
	var v string
	if tree.Color {
		v = fmt.Sprintf("%d-1", tree.Val)

	} else {
		v = fmt.Sprintf("%d-0", tree.Val)
	}
	*list = append(*list, v)
	RBFrontIndexTree(tree.Left, list)
	RBFrontIndexTree(tree.Right, list)

}
