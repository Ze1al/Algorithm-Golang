/*
*  @author: didichuxing.com
 */

package sort

import (
	"math/rand"
	"time"
)

// 单路快排，时间复杂度O(nlogn), 最差时间复杂度是O(n2)
// 快排是不稳定的，受到原来数组的顺序影响，若原来的数组就是就有序，则会将数据全分到一遍，最后导致变成冒泡排序
// 单路快排没有考虑数据与基准值相等的情况，容易一边大一边小
// 双路快排对于相等的值就是不交换，多路快排则针对相等的值单独提取

func QSortRecursion(arr []int) {
	if len(arr) < 2 {
		return
	}

	head, tail := 0, len(arr)-1
	randomNum := getRandNum(len(arr) - 1)
	arr[0], arr[randomNum] = arr[randomNum], arr[0]
	poivt := arr[0]

	for i := 1; i <= tail; {
		if arr[i] <= poivt {
			arr[i], arr[head] = arr[head], arr[i]
			head++
			i++
		} else {
			arr[i], arr[tail] = arr[tail], arr[i]
			tail--
		}
	}
	QSortRecursion(arr[:head])
	QSortRecursion(arr[head+1:])
}

func getRandNum(num int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(num)
}

func quickSort(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	head, tail := 0, len(nums)
	randomNum := getRandNum(n)
	nums[0], nums[randomNum] = nums[randomNum], nums[0]
	pviot := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] <= pviot {
			nums[i], nums[head] = nums[head], nums[i]
			i++
			head++
		} else {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		}
	}
	quickSort(nums[:head])
	quickSort(nums[head+1:])
}
