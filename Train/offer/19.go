package offer

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	c, x0 := float64(x), float64(x)
	for {
		xi := (x0 + c/x0) / 2
		if abs2(x0-xi) < 0.01 {
			break
		}
		x0 = xi
	}
	return int(x0)
}

func abs2(x float64) float64 {
	if x < 0 {
		return -1 * x
	}
	return x
}
