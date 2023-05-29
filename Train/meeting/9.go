package meeting

func convColumnNum(num int) string {

	var res string
	for num > 0 {
		// 如果小于26，则代表除到了最小一位，则直接添加
		if num < 26 {
			letter := num + 'A'
			res += string(rune(letter))
			break
		} else {
			first := num / 26
			num = num % 26

			letter := first + 'A'
			res = res + string(rune(letter))
		}
	}
	return res
}
