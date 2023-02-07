package code

import (
	"strconv"
	"strings"
)

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
	var r string
	s := strings.Split(strconv.Itoa(x), "")
	for i := len(s) - 1; i >= 0; i-- {
		r += s[i]
	}
	num, _ := strconv.Atoi(r)

	return x == num
}

// @lc code=end
