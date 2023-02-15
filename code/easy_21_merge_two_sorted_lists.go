package code

/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 */
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// https://leetcode.com/problems/merge-two-sorted-lists/solutions/360472/golang-0ms-recursive-solution/?orderBy=most_votes&languageTags=golang
	// 模範解答
	// 再帰で表現する

	// l1がなくなれば後はl2を返すだけ
	if l1 == nil {
		return l2
	}
	// l2がなくなれば後はl1を返すだけ
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		// l1の方が小さければ、l1を頭にし、l1のNextとl2を再帰で比較する
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	// l2の方が小さければ、l2を頭にし、l2のNextとl1を再帰で比較する
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

// @lc code=end
