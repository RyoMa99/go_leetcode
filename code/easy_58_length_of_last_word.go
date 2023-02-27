package code

/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	ans := make([]int, 0)

	c := 0
	for _, v := range s {
		if v == ' ' {
			if c != 0 {
				ans = append(ans, c)
			}
			c = 0
		} else {
			c++
		}
	}

	if c != 0 {
		ans = append(ans, c)
	}

	return ans[len(ans)-1]
}

// @lc code=end
