package dp

// https://leetcode.cn/problems/maximum-length-of-repeated-subarray/
// 两个数组 最长重复子数组
func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1))
	for i, _ := range dp {
		dp[i] = make([]int, len(nums2))
	}

	for i := 0; i < len(nums1); i++ {
		if nums1[i] == nums2[0] {
			dp[i][0] = 1
		}
	}

	var res int

	for j := 0; j < len(nums2); j++ {
		if nums2[j] == nums1[0] {
			dp[0][j] = 1
			res = 1
		}
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}

	return res
}
