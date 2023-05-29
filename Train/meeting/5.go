package meeting

import (
	"strconv"
)

// 猿辅导
// 求表达式的值
func evalRPN(tokens []string) int {
	stack := []int{}

	for _, v := range tokens {
		num, err := strconv.Atoi(v)
		if err == nil {
			stack = append(stack, num)
		} else {
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			if v == "+" {
				stack = append(stack, num1+num2)
			}
			if v == "-" {
				stack = append(stack, num1-num2)
			}
			if v == "/" {
				stack = append(stack, num1/num2)
			}
			if v == "*" {
				stack = append(stack, num1*num2)
			}
		}
	}
	return stack[len(stack)-1]
}
