package code

/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	nx := 0
	for i := 0; i < m+n; i++ {
		if nx < n && (nums1[i] > nums2[nx] || i >= m+nx && nums1[i] == 0) {
			nums1 = nums1[:i+copy(nums1[i:], append([]int{nums2[nx]}, nums1[i:]...))]
			nx++
		}
	}
	nums1 = nums1[:n+m]
}

// @lc code=end
