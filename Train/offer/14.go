package offer

// 求幂等 快速幂
func myPow(x float64, n int) float64 {
	if n < 0 {
		return -1 / quickPow(x, -1*n)
	}
	return quickPow(x, n)
}

// 举例，x = 2, n = 10
// 后续遍历，所以是左右中
func quickPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	y := quickPow(x, n/2)
	if n%2 == 1 {
		return y * y * x
	}
	return y * y
}
