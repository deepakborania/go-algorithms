package main

import "fmt"

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func bottomView(root *TreeNode) []int {
	m := map[int]*TreeNode{}
	helper(root, 0, m)
	fmt.Println(m)
	return nil
}

func helper(x *TreeNode, dist int, m map[int]*TreeNode) {
	if x == nil {
		return
	}
	m[dist] = x
	helper(x.left, dist-1, m)
	helper(x.right, dist+1, m)
}

func main() {
	root := &TreeNode{
		val:   9,
		left:  &TreeNode{10, &TreeNode{1, nil, nil}, &TreeNode{7, &TreeNode{5, nil, nil}, &TreeNode{6, nil, nil}}},
		right: &TreeNode{2, &TreeNode{3, nil, nil}, nil},
	}
	fmt.Println(bottomView(root))
}
