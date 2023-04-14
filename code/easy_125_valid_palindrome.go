package code

import "unicode"

/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

// @lc code=start
func isPalindrome(s string) bool {
	sr := []rune(s)
	ans := true
	l := 0
	r := len(s) - 1
	for {
		for {
			if l >= r {
				break
			}
			if !(unicode.IsDigit(sr[l]) || unicode.IsLetter(sr[l])) {
				l++
				continue
			}
			if !(unicode.IsDigit(sr[r]) || unicode.IsLetter(sr[r])) {
				r--
				continue
			}
			break
		}
		if unicode.ToLower(sr[l]) != unicode.ToLower(sr[r]) {
			ans = false
			break
		}
		if l >= r {
			break
		}
		l++
		r--
	}
	return ans
}

// @lc code=end
