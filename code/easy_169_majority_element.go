package code

/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
func majorityElement(nums []int) int {
	ans := make(map[int]int, len(nums))
	for _, n := range nums {
		ans[n]++
	}

	max_i, max_c := 0, 0
	for i, n := range ans {
		if n > max_c {
			max_i = i
			max_c = n
		}
	}
	return max_i
}

// @lc code=end
