package greed

import "fmt"

// 分发糖果
// https://leetcode.cn/problems/candy/
func candy(ratings []int) int {
	candyVee := make([]int, len(ratings))

	// 从左往右
	candyVee[0] = 1
	for i := 1; i <= len(ratings)-1; i++ {
		if ratings[i] > ratings[i-1] {
			candyVee[i] = candyVee[i-1] + 1
		} else {
			candyVee[i] = 1
		}
	}
	fmt.Printf("arr1=%v\n", candyVee)

	// 从右到左
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candyVee[i] = max(candyVee[i], candyVee[i+1]+1)
		}
	}

	fmt.Printf("arr2=%v\n", candyVee)

	return sum(candyVee)
}

func sum(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
}
