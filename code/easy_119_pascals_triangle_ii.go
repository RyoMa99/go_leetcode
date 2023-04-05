package code

/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	nums := []int{1}
	for i := 1; i < rowIndex+1; i++ {
		next := []int{}
		for j := 0; j < len(nums)+1; j++ {
			if j == 0 || j == len(nums) {
				next = append(next, 1)
				continue
			}
			next = append(next, nums[j-1]+nums[j])
		}
		nums = next
	}
	return nums
}

// @lc code=end
