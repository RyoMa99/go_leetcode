package code

/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 */

// @lc code=start
func singleNumber(nums []int) int {
	numSet := NewSet()

	for _, num := range nums {
		if numSet.Contains(num) {
			numSet.Remove(num)
		} else {
			numSet.Add(num)
		}
	}

	return numSet.Values()[0]
}

type Set map[int]struct{}

func NewSet() Set {
	return make(Set)
}

func (s Set) Add(t int) {
	s[t] = struct{}{}
}

func (s Set) Remove(t int) {
	delete(s, t)
}

func (s Set) Contains(t int) bool {
	_, exists := s[t]
	return exists
}

func (s Set) Values() []int {
	values := make([]int, 0, len(s))
	for k := range s {
		values = append(values, k)
	}
	return values
}

// @lc code=end
