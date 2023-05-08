package code

/*
 * @lc app=leetcode id=190 lang=golang
 *
 * [190] Reverse Bits
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	// https://leetcode.com/problems/reverse-bits/solutions/426492/golang-100-super-simple/?orderBy=most_votes&languageTags=golang
	var ans uint32
	for i := 0; i < 32; i++ {
		ans = ans<<1 + num&1
		num >>= 1
	}
	return ans
}

// @lc code=end
