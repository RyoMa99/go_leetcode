package code

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	brackets := map[rune]rune{')': '(', '}': '{', ']': '['}

	queue := make([]rune, 0)
	for _, v := range s {
		if _, ok := brackets[v]; ok {
			if len(queue) > 0 && queue[len(queue)-1] == brackets[v] {
				queue = queue[:len(queue)-1]
				continue
			}
			return false
		}
		queue = append(queue, v)
	}

	return len(queue) == 0
}

// @lc code=end
