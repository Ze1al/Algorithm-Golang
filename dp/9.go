package dp

// 0-1 背包问题
/*
问题描述：重量为4的背包，如何装下下面的物品使其价值最高，每件物品只有一件

物品     重量      价值
1        1        15
2        3        20
3        4        30

接上文的背包问题，使用一维数组
*/

func BagProblem2(weight []int, values []int, bagSize int) {
	length := len(weight)
	// 1. 容量为J的背包最大价值
	dp := make([]int, bagSize+1)

	// 2. 初始化 值均为0

	// 3. 遍历，注意遍历顺序，一定是先遍历物品，再遍历背包大小，而且背包大小倒序
	for i := 0; i < length; i++ {
		for j := bagSize; j > 0; j-- {
			dp[j] = max(dp[j], dp[j-weight[i]]+values[i])
		}
	}
	return
}
