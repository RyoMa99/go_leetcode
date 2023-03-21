package code

/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 type ListNode struct {
	 Val  int
	 Next *ListNode
 }
*/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	next := head.Next

	for next != nil {
		if head.Val == next.Val {
			next = next.Next
			continue
		}
		break
	}
	head.Next = next
	deleteDuplicates(head.Next)

	return head
}

// @lc code=end
