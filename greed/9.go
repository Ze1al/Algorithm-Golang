package greed

// 柠檬水找零
// https://leetcode.cn/problems/lemonade-change/
func lemonadeChange(bills []int) bool {
	five, ten, twenty := 0, 0, 0
	for _, v := range bills {
		if v == 5 {
			five += 1
		}
		// 如果是10块的
		if v == 10 {
			if five == 0 {
				return false
			}
			ten += 1
			five -= 1
		}
		// 如果是20块的
		if v == 20 {
			if ten > 0 && five > 0 {
				ten -= 1
				five -= 1
				twenty += 1
				continue
			}

			if ten == 0 && five >= 3 {
				five -= 3
				twenty += 1
				continue
			}
			return false
		}
	}
	return true
}
