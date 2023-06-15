package code

/*
 * @lc app=leetcode id=203 lang=golang
 *
 * [203] Remove Linked List Elements
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeElements(head *ListNode, val int) *ListNode {
	// https://leetcode.com/problems/remove-linked-list-elements/solutions/284872/100-golang-solution/
	if head == nil {
		return nil
	}

	tmp := head
	for tmp.Next != nil {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}

	if head.Val == val {
		return head.Next
	}
	return head
}

// @lc code=end
