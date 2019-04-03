package binarysearchtrees

import "testing"

func TestValidateBSTFalse(t *testing.T) {
	root := &TreeNode{
		val:   9,
		left:  &TreeNode{10, &TreeNode{1, nil, nil}, &TreeNode{7, &TreeNode{5, nil, nil}, &TreeNode{6, nil, nil}}},
		right: &TreeNode{2, &TreeNode{3, nil, nil}, nil},
	}
	got := ValidateBST(root)
	if got {
		t.Errorf("ValidateBST() = %v; want false", got)
	}
}

func TestValidateBSTTrue(t *testing.T) {
	root := &TreeNode{
		val:   9,
		left:  &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{6, nil, nil}}},
		right: &TreeNode{11, &TreeNode{10, nil, nil}, nil},
	}
	got := ValidateBST(root)
	if !got {
		t.Errorf("ValidateBST() = %v; want true", got)
	}
}
