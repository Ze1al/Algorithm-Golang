package offer

import "container/heap"

// 出现频率最高的 k 个数字
// 典型的堆解法

// 二维数组，[1,2] 第一个元素代表nums，第二个元素代表出现的次数

// 大于堆顶元素的，则加入
// 小于堆顶元素的，则代表还有K个元素比他大，则抛弃

func topKFrequent(nums []int, k int) []int {
	freMap := make(map[int]int)
	for _, v := range nums {
		freMap[v] += 1
	}

	iHeap := &IHeap{}
	heap.Init(iHeap)

	for nums, fre := range freMap {
		ele := [2]int{nums, fre}
		heap.Push(iHeap, ele)
		if iHeap.Len() > k {
			heap.Pop(iHeap)
		}
	}

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(iHeap).([2]int)[0]
	}
	return res
}

type IHeap [][2]int

func (iheap IHeap) Len() int {
	return len(iheap)
}

func (iheap IHeap) Less(i, j int) bool {
	return iheap[i][1] < iheap[j][1]
}

func (iheap IHeap) Swap(i, j int) {
	iheap[i], iheap[j] = iheap[j], iheap[i]
}

func (iheap *IHeap) Push(x interface{}) {
	*iheap = append(*iheap, x.([2]int))
}

func (iheap *IHeap) Pop() interface{} {
	old := *iheap
	last := old[len(*iheap)-1]
	*iheap = old[:len(*iheap)-1]
	return last
}
