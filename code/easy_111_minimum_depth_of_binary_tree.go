package code

/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
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

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ld := minDepth(root.Left)
	rd := minDepth(root.Right)

	if ld == 0 && rd == 0 {
		return 1
	}
	if ld == 0 {
		return rd + 1
	}
	if rd == 0 {
		return ld + 1
	}
	return min(ld, rd) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
