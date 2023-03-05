package code

import "strconv"

/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */

// @lc code=start
func addBinary(a string, b string) string {
	if len(a) > len(b) {
		return exec(a, b)
	} else {
		return exec(b, a)
	}
}

func exec(a, b string) string {
	o := 0
	ans := ""

	for i := 0; i < len(a); i++ {
		r := 0
		if i < len(b) {
			r = int(a[len(a)-1-i]-'0') + int(b[len(b)-1-i]-'0') + o
		} else {
			r = int(a[len(a)-1-i]-'0') + o
		}
		if r > 1 {
			o = 1
		} else {
			o = 0
		}

		ans = strconv.Itoa(r%2) + ans
	}

	if o == 1 {
		ans = strconv.Itoa(1) + ans
	}

	return ans
}

// @lc code=end
