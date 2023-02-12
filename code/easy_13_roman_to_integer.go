package code

/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
//
//lint:ignore U1000 //
func romanToInt(s string) int {
	// 参考
	// https://leetcode.com/problems/roman-to-integer/solutions/580404/golang-simplest-and-efficient-solution-100-faster/?orderBy=most_votes&languageTags=golang

	var romanMap = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	// "MCMXCIV"
	// 一番後ろの文字をintに変更
	var result = romanMap[s[len(s)-1]] // 5

	// 後ろから見ていく
	for i := len(s) - 2; i >= 0; i-- {
		if romanMap[s[i]] < romanMap[s[i+1]] {
			// 右隣の値がチェックしている値より大きい場合はひく
			result -= romanMap[s[i]]
		} else {
			// 右隣の値がチェックしている値と同じか小さい場合は足す
			result += romanMap[s[i]]
		}
	}
	return result
}

// @lc code=end
