package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func serialize(root *TreeNode) string {
	q := []*TreeNode{root}
	res := ""
	for len(q) > 0 {
		elem := q[0]
		q = q[1:]
		if elem != nil {
			q = append(q, elem.left, elem.right)
			res += strconv.Itoa(elem.val) + ","
		} else {
			res += "nil,"
		}
	}
	return res
}

func deserialize(treeStr string) *TreeNode {
	if treeStr == "" {
		return nil
	}
	ts := strings.Split(treeStr, ",")
	v, _ := strconv.Atoi(ts[0])
	root := &TreeNode{val: v}
	q := []*TreeNode{root}
	for i := 1; i < len(ts) && len(q) > 0; i++ {
		parent := q[0]
		q = q[1:]
		if ts[i] != "nil" {
			v, _ := strconv.Atoi(ts[i])
			parent.left = &TreeNode{val: v}
			q = append(q, parent.left)
		}
		if i+1 < len(ts) {
			i++
			if ts[i] != "nil" {
				v, _ := strconv.Atoi(ts[i])
				parent.right = &TreeNode{val: v}
				q = append(q, parent.right)
			}

		}
	}
	return root
}

func PreOrderTravesal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d,", root.val)
	PreOrderTravesal(root.left)
	PreOrderTravesal(root.right)
}

func main() {
	root := &TreeNode{
		val:   9,
		left:  &TreeNode{10, &TreeNode{1, nil, nil}, &TreeNode{7, &TreeNode{5, nil, nil}, &TreeNode{6, nil, nil}}},
		right: &TreeNode{2, &TreeNode{3, nil, nil}, nil},
	}
	PreOrderTravesal(root)
	fmt.Println("")
	fmt.Println(serialize(root))
	str := serialize(root)
	PreOrderTravesal(deserialize(str))
}
