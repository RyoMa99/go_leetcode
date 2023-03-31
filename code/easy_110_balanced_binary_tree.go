package code

/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
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

func isBalanced(root *TreeNode) bool {
	ok, _ := checkDepth(root)
	return ok
}

func checkDepth(r *TreeNode) (bool, int) {
	if r == nil {
		return true, 0
	}

	ok_l, num_l := checkDepth(r.Left)
	if !ok_l {
		return false, 0
	}

	ok_r, num_r := checkDepth(r.Right)
	if !ok_r {
		return false, 0
	}

	if max(num_l, num_r)-min(num_l, num_r) > 1 {
		return false, 0
	}

	return true, max(num_l, num_r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
