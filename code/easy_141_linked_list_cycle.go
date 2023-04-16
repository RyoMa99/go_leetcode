package code

/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// https://leetcode.com/problems/linked-list-cycle/solutions/1829693/python-go-floyd-s-cycle-detection-solution-and-explanation/?orderBy=most_votes&languageTags=golang
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for fast.Next != nil && fast.Next.Next != nil && fast != slow {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow == fast
}

// @lc code=end
