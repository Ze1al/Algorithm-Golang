package backTrace

var res4 [][]int

// 使任意数字相加等于Target
func combinationSum(candidates []int, target int) [][]int {
	res4 = [][]int{}

	backTracking4(candidates, target, []int{}, 0)
	return res4
}

func backTracking4(candidates []int, target int, path []int, startIndex int) {
	if listSum(path) == target {
		temp := make([]int, len(path))
		copy(temp, path)
		res4 = append(res4, temp)
		return
	}
	if listSum(path) > target {
		return
	}

	for i := startIndex; i < len(candidates); i++ {
		path = append(path, candidates[i])
		backTracking4(candidates, target, path, i)
		path = path[:len(path)-1]
	}
}
