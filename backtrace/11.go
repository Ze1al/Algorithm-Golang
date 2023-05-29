package backTrace

var res11 [][]int

func permute(nums []int) [][]int {
	res11 = [][]int{}
	used := make([]int, len(nums))

	backTracking11(nums, []int{}, used)
	return res11
}

func backTracking11(nums []int, path []int, used []int) {
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		res11 = append(res11, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] == 1 {
			continue
		}
		path = append(path, nums[i])
		used[i] = 1
		backTracking11(nums, path, used)
		used[i] = 0
		path = path[:len(path)-1]
	}
}
