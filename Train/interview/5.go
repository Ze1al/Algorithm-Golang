package interview

import "strconv"

// 问题：给定数字num，删除其中K个数字，让剩余的数字保持相对顺序不变动，使其数字最大
// 举例：
// num = 32406  k = 1, 删除2，剩余3406
// num = 32406  k = 2, 删除32，剩余406

// 思路：
// 贪心算法，局部最优保证全局最优
// 怎么样让一个数最大，本质肯定是删除数字中最小的number，但是需要考虑高低位问题，比如例子一中的，删除高位的2比删除低位的0更加有效果
// 让剩余的数字，优先让大的数字处于高位，满足数字前半部分是一个递减的规律，能满足其最大

// 边界情况，如果数字本来就是递减的，类似987654321，则需要删除最小位置

func removeKDigits(num int, k int) int {
	nums2 := strconv.Itoa(num)
	if k >= len(nums2) {
		return 0
	}

	for k > 0 {
		i := 1
		flag := false
		for i < len(nums2) {
			// 递增规律，将i-1的数字删除，让后面大的数字顶上
			if nums2[i] > nums2[i-1] {
				nums2 = nums2[:i-1] + nums2[i:]
				flag = true
				break
			} else {
				i += 1
			}
		}

		// 如果没有删除过，则需要删除最后一位
		if !flag {
			nums2 = nums2[:len(nums2)-1]
		}

		k -= 1
	}

	val, _ := strconv.Atoi(nums2)
	return val
}
