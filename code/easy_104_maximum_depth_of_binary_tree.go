package code

/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(measureDepth(root.Left, 1), measureDepth(root.Right, 1))
}

func measureDepth(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	return max(measureDepth(root.Left, depth+1), measureDepth(root.Right, depth+1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
