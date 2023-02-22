package code

/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	if target > nums[r] {
		return r + 1
	}

	if target <= nums[l] {
		return l
	}

	for l != r {
		mid := (l + r) / 2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			l = mid
		} else {
			r = mid
		}

		if l+1 == r {
			return r
		}
	}

	return l

}

// @lc code=end
