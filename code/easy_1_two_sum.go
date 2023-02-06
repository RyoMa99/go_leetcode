package code

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
//
//lint:ignore U1000 //
func twoSum(nums []int, target int) []int {
	// hashmapを使う方法
	// https://www.code-recipe.com/post/two-sum#viewer-9p12l
	// 最初にnumsを全てhashmapにしてしまうと、同値のものの扱いに困る
	// numsを走査しつつ、(target - 走査対象の一要素)の値がhashmapにあれば、その要素のインデックスとhashmapの値が答えになる
	// 無ければhashmapにその要素をキー、要素位置をバリューとし追加する
	hashmap := make(map[int]int, 0)

	for i, v := range nums {
		if j, ok := hashmap[target-v]; ok {
			return []int{j, i}
		} else {
			hashmap[v] = i
		}
	}

	return nil
}

// @lc code=end
