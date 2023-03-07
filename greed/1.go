package greed

import "sort"

// 分发饼干
// https://leetcode.cn/problems/assign-cookies/

func findContentChildren(g []int, s []int) int {
	sort.Ints(g) // 胃口
	sort.Ints(s) // 饼干

	index := len(s) - 1
	res := 0
	for i := len(g) - 1; i >= 0; i-- {
		if index >= 0 && g[i] <= s[index] {
			res += 1
			index -= 1
		}
	}
	return res
}
