package main

import "fmt"

// TreeNode represents node in a binary tree
type TreeNode struct {
	val         int
	left, right *TreeNode
}

// InorderTraversal implements iterative inorder traversal of the binary tree
func InorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	// Initialize our stack
	stack := []*TreeNode{}
	curr := root
	for curr != nil || len(stack) > 0 {
		// While we can go left, go left
		for curr != nil {
			stack = append([]*TreeNode{curr}, stack...)
			curr = curr.left
		}
		// Pop top element from the stack and add to our result list
		curr, stack = stack[0], stack[1:]
		result = append(result, curr.val)

		//Go to the right of current element
		curr = curr.right
	}
	return result
}

// PreorderTraversal implements iterative preorder traversal of the binary tree
func PreorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	//Push root to stack
	stack := []*TreeNode{root}

	// while stack is not empty
	for len(stack) > 0 {
		curr := stack[0]
		stack = stack[1:]
		// Add popped element to result
		result = append(result, curr.val)

		//First add right element to stack
		if curr.right != nil {
			stack = append([]*TreeNode{curr.right}, stack...)
		}
		//then left
		if curr.left != nil {
			stack = append([]*TreeNode{curr.left}, stack...)
		}

	}
	return result
}

func main() {
	root := &TreeNode{
		val:   9,
		left:  &TreeNode{10, &TreeNode{1, nil, nil}, &TreeNode{7, &TreeNode{5, nil, nil}, &TreeNode{6, nil, nil}}},
		right: &TreeNode{2, &TreeNode{3, nil, nil}, nil},
	}
	fmt.Println("Inorder: ", InorderTraversal(root))
	fmt.Println("Preorder:", PreorderTraversal(root))

}
