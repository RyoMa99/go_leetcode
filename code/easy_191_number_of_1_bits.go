package code

/*
 * @lc app=leetcode id=191 lang=golang
 *
 * [191] Number of 1 Bits
 */

// @lc code=start
func hammingWeight(num uint32) int {
	// https://leetcode.com/problems/number-of-1-bits/solutions/468836/python-js-java-go-c-o-lg-n-by-bit-manipulation/?orderBy=most_votes&languageTags=golang
	cnt := 0

	for i := 0; i < 32; i++ {
		if num&1 == 1 {
			cnt++
		}
		num >>= 1
	}

	return cnt
}

// @lc code=end
