package backTrace

var res1 [][]int

func combinationSum3(k int, n int) [][]int {
	res1 = [][]int{}

	backtracking2(k, n, []int{}, 1)
	return res1
}

func backtracking2(k, n int, path []int, startIndex int) {
	if len(path) == k && listSum(path) == n {
		temp := make([]int, k)
		copy(temp, path)
		res1 = append(res1, temp)
		return
	}

	for i := startIndex; i <= 9; i++ {
		path = append(path, i)
		backtracking2(k, n, path, i+1)
		path = path[:len(path)-1]
	}
}

func listSum(nums []int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	return sum
}
