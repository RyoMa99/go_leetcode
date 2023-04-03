package code

/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
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

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	num := targetSum - root.Val
	if root.Left == nil && root.Right == nil && num == 0 {
		return true
	}
	return hasPathSum(root.Left, num) || hasPathSum(root.Right, num)
}

// @lc code=end
