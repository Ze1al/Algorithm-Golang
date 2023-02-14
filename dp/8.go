package dp

// 0-1 背包问题
/*
问题描述：重量为4的背包，如何装下下面的物品使其价值最高，每件物品只有一件

物品     重量      价值
1        1        15
2        3        20
3        4        30

*/

func BagProblem(weight []int, values []int, bagSize int) {
	length := len(weight)

	// 1. 二维数组，dp[i][j] 代表各个[0,i]的物品任取，且容量为j的最大价值
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, bagSize+1)
	}

	// 2. 初始化
	// 容量为0的时候，价值均为0
	for i := 0; i < len(dp); i++ {
		dp[i][0] = 0
	}
	// 只选物品0时候，各个容量对应的最大价值
	for j := 0; j <= bagSize; j++ {
		if j < weight[0] {
			dp[0][j] = 0
		} else {
			dp[0][j] = values[0]
		}
	}

	// 3. 递推公式
	for i := 1; i < len(dp); i++ {
		for j := 1; j < bagSize; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+values[i])
		}
	}

	return
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
