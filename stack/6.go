package stack

func maxSlidingWindow(nums []int, k int) []int {
	var (
		windows []int
		res     []int
	)
	for _, v := range nums {
		if len(windows) >= k {
			res = append(res, maxNum(windows))
			windows = windows[1:]
		}
		windows = append(windows, v)
	}
	res = append(res, maxNum(windows))

	return res
}

func maxNum(num []int) int {
	max := num[0]
	for i := 1; i < len(num); i++ {
		if max < num[i] {
			max = num[i]
		}
	}
	return max
}

func sum(num []int) int {
	sum := 0
	for _, v := range num {
		sum += v
	}
	return sum
}
