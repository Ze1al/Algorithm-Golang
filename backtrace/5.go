package backTrace

import "sort"

var res5 [][]int

func combinationSum2(candidates []int, target int) [][]int {
	res5 = [][]int{}
	path := []int{}
	used := make([]bool, len(candidates))
	sort.Ints(candidates)
	backTracking5(candidates, target, path, used, 0)
	return res5
}

func backTracking5(candidates []int, target int, path []int, used []bool, startIndex int) {
	if listSum(path) == target {
		temp := make([]int, len(path))
		copy(temp, path)
		res5 = append(res5, temp)
		return
	}
	if listSum(path) > target {
		return
	}

	for i := startIndex; i < len(candidates); i++ {
		if i > 0 && candidates[i-1] == candidates[i] && !used[i-1] {
			continue
		}
		path = append(path, candidates[i])
		used[i] = true
		backTracking5(candidates, target, path, used, i+1)
		used[i] = false
		path = path[:len(path)-1]
	}
}
