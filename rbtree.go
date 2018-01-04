package trees

import (
	"fmt"
	"strings"
)

const RED = true
const BLACK = false

type RBTree struct {
	Color       bool // Red true Black false
	Val         int
	Left, Right *RBTree
}

type RBT struct {
	root *RBTree
}

var TNil = &RBTree{Color: BLACK}

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
						node = &RBTree{Val: v, Color: RED}
						root.Left = node
						break
					} else {

						root = root.Left
					}
				} else {
					fathers = append(fathers, root)
					lefts = append(lefts, false)
					if root.Right == nil {
						node = &RBTree{Val: v, Color: RED}
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
	n := len(fathers)
	if n == 0 {
		node.Color = BLACK
		return nil
	} else if fathers[n-1].Color == BLACK {
		return nil
	} else if n >= 2 && uncle(fathers, lefts) != nil && uncle(fathers, lefts).Color == RED {
		fathers[n-1].Color = BLACK
		uncle(fathers, lefts).Color = BLACK
		fathers[n-2].Color = RED
		return fixFatherBalance(fathers[n-2], fathers[:n-2], lefts[:n-2])
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

func (r *RBT) Delete(v int) {
	root := r.root
	fathers := []*RBTree{}
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

	root, ok := updateRoot(root, fathers, lefts)
	if ok {
		r.root = root
	}

}

func updateRoot(root *RBTree, fathers []*RBTree, lefts []bool) (*RBTree, bool) {
	var color bool
	if root.Left == nil && root.Right == nil {
		n := len(fathers)
		if n > 0 {
			if lefts[n-1] {
				fathers[n-1].Left = nil
			} else {
				fathers[n-1].Right = nil
			}
			return fixBalance(root.Color, TNil, fathers, lefts)

		} else {
			return nil, true
		}
	} else if root.Left == nil && root.Right != nil {
		color = root.Color
		root.Left, root.Right, root.Color, root.Val = root.Right.Left, root.Right.Right, root.Right.Color, root.Right.Val
		return fixBalance(color, root, fathers, lefts)
	} else if root.Left != nil && root.Right == nil {
		color = root.Color
		root.Left, root.Right, root.Color, root.Val = root.Left.Left, root.Left.Right, root.Left.Color, root.Left.Val
		return fixBalance(color, root, fathers, lefts)
	} else if root.Left.Right == nil {
		color = root.Left.Color
		root.Val = root.Left.Val
		root.Left = root.Left.Left
		fathers = append(fathers, root)
		lefts = append(lefts, true)
		if root.Left == nil {
			return fixBalance(color, TNil, fathers, lefts)
		}
		return fixBalance(color, root.Left, fathers, lefts)
	} else {
		father := root.Left
		node := root.Left.Right
		fathers = append(fathers, root)
		lefts = append(lefts, true)
		fathers = append(fathers, father)
		lefts = append(lefts, false)
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
		root.Val, father.Right = node.Val, node.Left
		color = node.Color
		if father.Right == nil {
			return fixBalance(color, TNil, fathers, lefts)
		}
		return fixBalance(color, father.Right, fathers, lefts)
	}

}

func fixBalance(color bool, child *RBTree, fathers []*RBTree, lefts []bool) (*RBTree, bool) {
	if color == BLACK {
		if child.Color == RED {
			child.Color = BLACK
		} else {
			return fixChildBalance(color, child, fathers, lefts)
		}
	}
	return nil, false
}

func sibling(fathers []*RBTree, lefts []bool) *RBTree {
	n := len(fathers)
	if lefts[n-1] {
		return fathers[n-1].Right
	} else {
		return fathers[n-1].Left
	}
}

func setSon(n int, son *RBTree, fathers []*RBTree, lefts []bool) {
	if lefts[n] {
		fathers[n].Left = son
	} else {
		fathers[n].Right = son
	}
}

func Color(r *RBTree) bool {
	if r == nil {
		return BLACK
	} else {
		return r.Color
	}
}

func fixChildBalance(color bool, child *RBTree, fathers []*RBTree, lefts []bool) (*RBTree, bool) {
	var s, father *RBTree
	var n int
	if len(fathers) != 0 {
		s = sibling(fathers, lefts)
		n = len(fathers)
		if s.Color == RED {
			fathers[n-1].Color = RED
			s.Color = BLACK
			father = fathers[n-1]
			if lefts[n-1] {
				father.Right, s.Left = s.Left, father
			} else {
				father.Left, s.Right = s.Right, father
			}
			fathers = append(fathers[:n-1], append([]*RBTree{s}, fathers[n-1:]...)...)
			lefts = append(lefts[:n-1], append([]bool{lefts[n-1]}, lefts[n-1:]...)...)
		}

		s = sibling(fathers, lefts)
		n = len(fathers)
		father = fathers[n-1]
		if father.Color == BLACK && s.Color == BLACK && Color(s.Left) == BLACK && Color(s.Right) == BLACK {
			s.Color = RED
			return fixChildBalance(BLACK, father, fathers[:n-1], lefts[:n-1])
		} else if father.Color == RED && s.Color == BLACK && Color(s.Left) == BLACK && Color(s.Right) == BLACK {
			s.Color = RED
			father.Color = BLACK
			return fathers[0], true
		} else {
			if s.Color == BLACK {

				if Color(s.Right) == BLACK && Color(s.Left) == RED && lefts[n-1] {
					s.Color = RED

					left := s.Left

					s.Left.Color = BLACK
					s.Left.Right, s.Left = s, s.Left.Right

					father.Right = left

					s = left
				} else if Color(s.Left) == BLACK && Color(s.Right) == RED && lefts[n-1] == false {
					s.Color = RED
					s.Right.Color = BLACK
					right := s.Right
					s.Right.Left, s.Right = s, s.Right.Left

					father.Left = right

					s = right
				}

			}
			s.Color = father.Color
			father.Color = BLACK
			if lefts[n-1] {
				s.Right.Color = BLACK
				father.Right, s.Left = s.Left, father
			} else {
				s.Left.Color = BLACK
				father.Left, s.Right = s.Right, father

			}
			fathers = append(fathers[:n-1], append([]*RBTree{s}, fathers[n-1:]...)...)
			lefts = append(lefts[:n-1], append([]bool{lefts[n-1]}, lefts[n-1:]...)...)
			return fathers[0], true

		}

	} else {
		return nil, false
	}

}

func (r *RBT) Find(v int) bool {
	root := r.root
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
