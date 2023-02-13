package code

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {

	ans := make([]rune, 0)

	for i := 0; i <= 200; i++ {
		if i == len(strs[0]) {
			return string(ans)
		}

		tmp := []rune(strs[0])[i]
		f := true
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) {
				return string(ans)
			}

			if tmp != []rune(strs[j])[i] {
				f = false
			}
		}

		if f {
			ans = append(ans, tmp)
		} else {
			break
		}

	}

	return string(ans)
}

// @lc code=end
