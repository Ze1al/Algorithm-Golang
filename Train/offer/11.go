package offer

import (
	"container/heap"
	"sort"
)

// 数据流中第K大的数字
// 解法，堆排序
type KthLargest struct {
	sort.IntSlice
	k int
}

func Constructor11(k int, nums []int) KthLargest {
	k1 := KthLargest{k: k}
	for _, v := range nums {
		k1.Add(v)
	}
	return k1
}

func (this *KthLargest) Push(x interface{}) {
	this.IntSlice = append(this.IntSlice, x.(int))
}

func (this *KthLargest) Pop() interface{} {
	a := this.IntSlice
	v := a[len(a)-1]
	this.IntSlice = a[:len(a)-1]
	return v
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this, val)
	if this.Len() > this.k {
		heap.Pop(this)
	}

	return this.IntSlice[0]
}
