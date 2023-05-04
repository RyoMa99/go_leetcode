package code

/*
 * @lc app=leetcode id=171 lang=golang
 *
 * [171] Excel Sheet Column Number
 */

// @lc code=start
// https://leetcode.com/problems/excel-sheet-column-number/solutions/654014/very-simple-5-line-solution/?orderBy=most_votes&languageTags=golang
func titleToNumber(columnTitle string) int {
	col := 0

	for _, char := range columnTitle {
		col *= 26
		col += int(char-'A') + 1
	}

	return col
}

// @lc code=end
