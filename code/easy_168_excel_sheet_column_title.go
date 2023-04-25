package code

/*
 * @lc app=leetcode id=168 lang=golang
 *
 * [168] Excel Sheet Column Title
 */

// @lc code=start
func convertToTitle(columnNumber int) string {
	ans := make([]byte, 0)

	for columnNumber != 0 {
		columnNumber--
		ans = append([]byte{byte('A' + columnNumber%26)}, ans...)
		columnNumber /= 26
	}

	return string(ans)
}

// @lc code=end
