package meeting

// 怎么计算十进制中1的个数
func hammingWeight(num uint32) int {
	res := 0

	for i := 0; i < 32; i++ {
		if 1<<i&num != 0 {
			res += 1
		}
	}
	return res
}
