package code

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
//
//lint:ignore U1000 //
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	// https://leetcode.com/problems/palindrome-number/solutions/981758/go-8ms-solution/?orderBy=most_votes&languageTags=golang
	origin := x
	ans := 0
	for x != 0 {
		// 剰余算を上手く使って数列を逆にする
		ans = ans*10 + x%10
		x /= 10
	}

	return origin == ans
}

// @lc code=end
