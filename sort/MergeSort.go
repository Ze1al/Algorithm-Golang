/*
*  @author: didichuxing.com
 */

package sort

// 归并排序，归并排序的时间复杂度O(nlogn) 空间复杂度是O(n)
// 稳定

func Merge(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	mid := n / 2
	leftArr := Merge(nums[:mid])
	rightArr := Merge(nums[mid:])
	res := MergeSort(leftArr, rightArr)
	return res
}

func MergeSort(leftArr []int, rightArr []int) []int {
	arrNew := []int{}
	var i, j int
	for i < len(leftArr) && j < len(rightArr) {
		if leftArr[i] < rightArr[j] {
			arrNew = append(arrNew, leftArr[i])
		} else {
			arrNew = append(arrNew, rightArr[j])
		}
	}
	arrNew = append(arrNew, leftArr[i:]...)
	arrNew = append(arrNew, rightArr[j:]...)
	return arrNew
}
