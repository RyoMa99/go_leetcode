package code

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	// https://leetcode.com/problems/longest-common-prefix/solutions/1517792/simple-go-solution/?orderBy=most_votes&languageTags=golang
	ans := strs[0]
	for _, s := range strs {
		i := 0
		for ; i < len(s) && i < len(ans) && ans[i] == s[i]; i++ {
		}
		ans = ans[:i]
	}
	return ans
}

// @lc code=end
