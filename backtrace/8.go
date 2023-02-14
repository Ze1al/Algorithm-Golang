package backTrace

var res8 [][]int

func subsets(nums []int) [][]int {
	res8 = [][]int{}

	backTracking8(nums, []int{}, 0)
	return res8
}

func backTracking8(nums []int, path []int, startIndex int) {
	temp := make([]int, len(path))
	copy(temp, path)
	res8 = append(res8, temp)

	for i := startIndex; i < len(nums); i++ {
		path = append(path, nums[i])
		backTracking8(nums, path, i+1)
		path = path[:len(path)-1]
	}
}
