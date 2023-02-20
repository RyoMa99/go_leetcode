package code

/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	ans := 0
	offset := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			offset++
			continue
		}
		ans++
		nums[i-offset] = nums[i]
	}
	return ans
}

// @lc code=end
