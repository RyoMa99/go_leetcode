package code

import "strconv"

/*
 * @lc app=leetcode id=191 lang=golang
 *
 * [191] Number of 1 Bits
 */

// @lc code=start
func hammingWeight(num uint32) int {
	cnt := 0
	for _, r := range strconv.FormatUint(uint64(num), 2) {
		if r == '1' {
			cnt++
		}
	}

	return cnt
}

// @lc code=end
