package backTrace

var phoneNumber map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var res3 []string

func letterCombinations(digits string) []string {
	res3 = []string{}
	if len(digits) == 0 {
		return res3
	}

	backtracking3(digits, "", 0)
	return res3
}

func backtracking3(digits string, path string, startIndex int) {
	if len(digits) == len(path) {
		res3 = append(res3, path)
		return
	}

	for i := startIndex; i < len(digits); i++ {
		letter, ok := phoneNumber[string(digits[i])]
		if !ok {
			continue
		}

		for _, v := range letter {
			path += string(v)
			backtracking3(digits, path, i+1)
			path = path[:len(path)-1]
		}
	}
}
