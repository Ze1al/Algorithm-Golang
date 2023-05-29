package offer

// 三角形中最小路径之和

/*
[2]
[3,4]
[6,5,7]
[4,1,8,3]
*/

// 思路一：通过dfs，直接计算

// 思路二：dp
// dp[i][j] 代表从节点走到i j位置，最小路劲和

// dp[i][j] 肯定是从上面走下来的
// dp[i][j] = min(dp[i-1][j-1] + nums[i-1][j-1], dp[i-1][j] + nums[i-1][j])

// dp[0][0] = nums[0][0]

// i正序，j倒序

func minimumTotal(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		dp[i] = make([]int, len(triangle[i]))
	}

	dp[0][0] = triangle[0][0]
	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == len(dp[i-1]) {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j]+triangle[i][j], dp[i-1][j-1]+triangle[i][j])
			}
		}
	}

	val := dp[len(dp)-1][0]
	for i := 1; i < len(dp[len(dp)-1]); i++ {
		if val > dp[len(dp)-1][i] {
			val = dp[len(dp)-1][i]
		}
	}
	return val
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
