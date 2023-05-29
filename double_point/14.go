package double_point

// 快乐数字
// 如果按照个位十位数的平方，相加等于0，则为快乐数
func isHappy(n int) bool {
	m := make(map[int]int)

	for {
		res := calNSquare(n)
		if res == 1 {
			return true
		}

		_, ok := m[res]
		if ok {
			return false
		} else {
			m[res] += 1
			n = res
		}
	}
}

func calNSquare(n int) int {
	sum := 0
	for n > 0 {
		s := n % 10

		sum += s * s
		n = n / 10
	}
	return sum
}
