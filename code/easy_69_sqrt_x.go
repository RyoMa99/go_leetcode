package code

/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) int {
	for ans := 0; ; ans++ {
		if ans*ans > x {
			return ans - 1
		}
	}
}

// @lc code=end
