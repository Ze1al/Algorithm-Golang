package list

// 水果成篮
// https://leetcode.cn/problems/fruit-into-baskets/
func totalFruit(fruits []int) int {
	res := 0

	left := 0
	fruitMap := make(map[int]int)
	for right := 0; right < len(fruits); right++ {
		rightNode := fruits[right]
		fruitMap[rightNode]++

		for len(fruitMap) > 2 {
			leftNode := fruits[left]
			fruitMap[leftNode]--
			left++
			if fruitMap[leftNode] == 0 {
				delete(fruitMap, leftNode)
			}
		}

		if right-left+1 > res {
			res = right - left + 1
		}
	}
	return res
}
