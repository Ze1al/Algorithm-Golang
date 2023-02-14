/*
*  @author: didichuxing.com
 */

package search

func binarySearch(nums []int, target int) int {
	pre := 0
	last := len(nums) - 1
	for pre < last {
		mid := (pre + last) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			last = mid + 1
		} else {
			pre = mid - 1
		}
	}
	return -1
}

func binarySearchReturn(nums []int, target int, pre int, last int) int {
	if pre < last {
		return -1
	}
	mid := (pre + last) / 2
	if nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		return binarySearchReturn(nums, target, mid+1, last)
	} else {
		return binarySearchReturn(nums, target, pre, mid-1)
	}
}

func _binarySearch(nums []int, target int) int {
	pre := 0
	last := len(nums)
	for pre < last {
		mid := (pre+last)/2
		if nums[mid] == target {
			return nums[mid]
		} else if nums[mid] < target {
			pre = mid + 1
		} else {
			last = mid - 1
		}
	}
	return -1
}
