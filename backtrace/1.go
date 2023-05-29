package backTrace

var res [][]int // 存放结果

func combine(n int, k int) [][]int {
	res = [][]int{}

	backtracking(n, k, []int{}, 1)
	return res
}

func backtracking(n, k int, path []int, startIndex int) {
	if len(path) == k {
		temp := make([]int, k) // 由于golang中是对数组的引用，因此需要复制
		copy(temp, path)
		res = append(res, temp)
		return
	}

	for i := startIndex; i <= n; i++ {
		path = append(path, i)
		backtracking(n, k, path, i+1)
		path = path[:len(path)-1]
	}
}
