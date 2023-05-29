package double_point

// 赎金信
func canConstruct(ransomNote string, magazine string) bool {
	letterMap := make(map[byte]int)

	for i := 0; i < len(magazine); i++ {
		letterMap[magazine[i]] += 1
	}
	for i := 0; i < len(ransomNote); i++ {
		num, ok := letterMap[ransomNote[i]]
		if ok && num != 0 {
			letterMap[ransomNote[i]] -= 1
		} else {
			return false
		}
	}
	return true
}
