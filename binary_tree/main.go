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

// validateBST validates if a binary tree is a valid BST by checking if the inorder traversal is sorted or not.
func validateBST(root *TreeNode) bool {
	stack := []*TreeNode{}
	var prev *TreeNode
	curr := root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append([]*TreeNode{curr}, stack...)
			curr = curr.left
		}
		curr, stack = stack[0], stack[1:]

		// If current element is not larger than previous element in inorder traversal, it is not BST.
		if prev != nil && curr.val <= prev.val {
			return false
		}

		prev = curr
		curr = curr.right
	}
	return true
}

func main() {
	root := &TreeNode{
		val:   9,
		left:  &TreeNode{10, &TreeNode{1, nil, nil}, &TreeNode{7, &TreeNode{5, nil, nil}, &TreeNode{6, nil, nil}}},
		right: &TreeNode{2, &TreeNode{3, nil, nil}, nil},
	}
	fmt.Println(InorderTraversal(root))
	fmt.Println("Is valid bst? ", validateBST(root))
}
