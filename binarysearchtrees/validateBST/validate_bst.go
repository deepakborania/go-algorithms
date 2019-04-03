package binarysearchtrees

// TreeNode represents node in a binary tree
type TreeNode struct {
	val         int
	left, right *TreeNode
}

// ValidateBST validates if a binary tree is a valid BST by checking if the inorder traversal is sorted or not.
func ValidateBST(root *TreeNode) bool {
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
