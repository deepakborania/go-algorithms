package main

import "fmt"

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func leftPerimeter(root *TreeNode) {
	if root != nil {
		if root.left != nil {
			fmt.Printf("%d ", root.val)
			leftPerimeter(root.left)
		} else if root.right != nil {
			fmt.Printf("%d ", root.val)
			leftPerimeter(root.right)
		}
	}
}

func rightPerimeter(root *TreeNode) {
	if root != nil {
		if root.right != nil {
			rightPerimeter(root.right)
			fmt.Printf("%d ", root.val)
		} else if root.left != nil {
			rightPerimeter(root.left)
			fmt.Printf("%d ", root.val)
		}
	}
}

func leaves(x *TreeNode) {
	if x != nil {
		leaves(x.left)
		if x.left == nil && x.right == nil {
			fmt.Printf("%d ", x.val)
		}
		leaves(x.right)
	}
}

func main() {
	root := &TreeNode{
		val:   1,
		left:  &TreeNode{2, &TreeNode{4, nil, &TreeNode{6, &TreeNode{9, nil, nil}, nil}}, &TreeNode{5, nil, &TreeNode{3, nil, &TreeNode{8, nil, nil}}}},
		right: &TreeNode{3, &TreeNode{7, nil, &TreeNode{10, nil, nil}}, nil},
	}
	leftPerimeter(root)
	leaves(root)
	rightPerimeter(root)
}
