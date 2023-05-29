/*
*  @author: didichuxing.com
 */

package sort

import (
	"testing"
)

func Test_qsort(t *testing.T) {
	type args struct {
		array []int
		low   int
		high  int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"Test", args{array: []int{3, 41, 24, 76, 11, 45, 3, 3, 64, 21, 69, 19, 36}, low: 0, high: 13}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_quickSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"Test", args{[]int{3, 41, 24, 76, 11, 45, 3, 3, 64, 21, 69, 19, 36}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}

}
