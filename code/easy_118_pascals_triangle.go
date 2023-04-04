package code

/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */

// @lc code=start
func generate(numRows int) [][]int {
	var ans [][]int

	ans = append(ans, []int{1})
	for i := 1; i < numRows; i++ {
		ans = append(ans, gen(ans[i-1]))
	}

	return ans
}

func gen(nums []int) []int {
	var ans []int

	for i := 0; i < len(nums)+1; i++ {
		if i == 0 || i == len(nums) {
			ans = append(ans, 1)
			continue
		}
		ans = append(ans, nums[i-1]+nums[i])
	}

	return ans
}

// @lc code=end
