package greed

import "sort"

// 根据身高重建队列
// 第一个元素为hi 代表身高，第二个元素为Ki 代表前头的几个元素
func reconstructQueue(people [][]int) [][]int {
	// 按照身高排序，如果身高相同，则K更小的排在前头
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})

	for i, p := range people {
		copy(people[p[1]+1:i+1], people[p[1]:i+1]) // 空出一个位置
		people[p[1]] = p
	}
	return people
}
