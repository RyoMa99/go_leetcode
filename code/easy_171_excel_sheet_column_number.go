package code

import "math"

/*
 * @lc app=leetcode id=171 lang=golang
 *
 * [171] Excel Sheet Column Number
 */

// @lc code=start
func titleToNumber(columnTitle string) int {
	ans := 0
	chars := []rune(columnTitle)
	for i := len(chars) - 1; i >= 0; i-- {
		ans += (int(chars[i]-'A') + 1) * int(math.Pow(26, float64(len(chars)-1-i)))
	}
	return ans
}

// @lc code=end
