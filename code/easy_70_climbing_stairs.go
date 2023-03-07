package code

/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start
func climbStairs(n int) int {
	dp := []int{}
	dp = append(dp, 1)
	dp = append(dp, 1)

	for i := 2; i <= n; i++ {
		dp = append(dp, dp[i-2]+dp[i-1])
	}

	return dp[n]
}

// @lc code=end
