package code

/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
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

// https://leetcode.com/problems/symmetric-tree/solutions/2796828/0-ms100-iteration-recursion/?orderBy=most_votes&languageTags=golang
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSame(root.Left, root.Right)
}

func isSame(l *TreeNode, r *TreeNode) bool {
	if l == nil || r == nil {
		return l == r
	}

	if l.Val != r.Val {
		return false
	}

	return isSame(l.Left, r.Right) && isSame(l.Right, r.Left)
}

// @lc code=end
