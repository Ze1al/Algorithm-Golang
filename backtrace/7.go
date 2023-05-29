package backTrace

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/restore-ip-addresses/

var res7 []string

func restoreIpAddresses(s string) []string {
	res7 = []string{}

	backTracking7(s, 0, []string{})
	return res7
}

func backTracking7(s string, startIndex int, path []string) {
	if len(path) == 4 && startIndex == len(s) {
		s2 := strings.Join(path, ".")
		res7 = append(res7, s2)
		return
	}

	for i := startIndex; i < len(s); i++ {
		sub := s[startIndex : i+1]
		if ipIsValid(sub) {
			path = append(path, sub)
			backTracking7(s, i+1, path)
			path = path[:len(path)-1]
		} else {
			break
		}
	}
}

// IP是否合法
func ipIsValid(s string) bool {
	if len(s) == 0 {
		return false
	}

	if s[0] == '0' {
		if len(s) == 1 {
			return true
		} else {
			return false
		}
	}

	num, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	if num < 0 || num > 255 {
		return false
	}
	return true
}
