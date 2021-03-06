package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	return helper(root, nil)
}

func helper(root *TreeNode, arr []int) []int {
	if root == nil {
		return arr
	}
	if root.Left != nil {
		arr = helper(root.Left, arr)
	}
	arr = append(arr, root.Val)
	if root.Right != nil {
		arr = helper(root.Right, arr)
	}
	return arr
}
