package code

/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
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

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := make([]int, 0)
	ans = append(ans, postorderTraversal(root.Left)...)
	ans = append(ans, postorderTraversal(root.Right)...)

	return append(ans, root.Val)
}

// @lc code=end
