package dp

// 完全背包
/*
相较于0-1背包，其实只有遍历顺序不同
for(int i = 0; i < weight.size(); i++) { // 遍历物品
	for(int j = bagWeight; j >= weight[i]; j--) { // 遍历背包容量
		dp[j] = max(dp[j], dp[j - weight[i]] + value[i]);
    }
}

完全背包，遍历顺序发生变化
for(int i = 0; i < weight.size(); i++) { // 遍历物品
	for(int j = weight[i]; j <= bagWeight; j--) { // 遍历背包容量
		dp[j] = max(dp[j], dp[j - weight[i]] + value[i]);
    }
}
*/

func completeBag(weight, value []int, bagWeight int) int {
	dp := make([]int, bagWeight+1)

	for i := 0; i < len(weight); i++ {
		for j := weight[i]; j <= bagWeight; j++ {
			dp[i] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}

	return dp[bagWeight]
}
