package code

/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	v := make([]int, 0)

	if root == nil {
		return v
	}

	v = append(v, inorderTraversal(root.Left)...)
	v = append(v, root.Val)
	v = append(v, inorderTraversal(root.Right)...)

	return v
}

// @lc code=end
