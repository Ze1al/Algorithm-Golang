package offer

import "strconv"

// 二进制求和
func addBinary(a string, b string) string {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}

	var res string
	p1, p2 := len(a)-1, len(b)-1
	c := 0

	for p1 >= 0 || p2 >= 0 {
		var val1, val2 int
		if p1 >= 0 {
			val1, _ = strconv.Atoi(string(a[p1]))
			p1 -= 1
		}
		if p2 >= 0 {
			val2, _ = strconv.Atoi(string(b[p2]))
			p2 -= 1
		}
		val := (val1 + val2 + c) % 2
		c = (val1 + val2 + c) / 2

		res = strconv.Itoa(val) + res
	}

	if c != 0 {
		res = strconv.Itoa(c) + res
	}

	return res
}
