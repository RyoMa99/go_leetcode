package code

import "sort"

/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
// https://leetcode.com/problems/majority-element/solutions/457064/majority-element/?orderBy=most_votes&languageTags=golang
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// @lc code=end
