package code

/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) int {
	if x == 1 {
		return 1
	}

	s, l := 0, x
	for {
		m := (s + l) / 2
		if m*m > x {
			l = m
		} else {
			if (m+1)*(m+1) > x {
				return m
			}
			s = m
		}
	}
}

// @lc code=end
