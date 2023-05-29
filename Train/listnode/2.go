package listnode

type SingleNode struct {
	Val  int
	Next *SingleNode
}

type MyLinkedList struct {
	Head *SingleNode
	Size int
}

//
//func Constructor() MyLinkedList {
//	return MyLinkedList{
//		Head: &SingleNode{}, // 亚节点
//		Size: 0,
//	}
//}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Size {
		return -1
	}

	p := this.Head
	for i := 0; i <= index; i++ {
		p = p.Next
	}
	return p.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.Size-1, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	// 如果index大于数组长度，则不插入
	if index > this.Size {
		return
	}
	index = max(index, 0)
	this.Size++

	p := this.Head
	for i := 0; i < index; i++ {
		p = p.Next
	}

	node := &SingleNode{Val: val, Next: p.Next}
	p.Next = node
	return
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}

	p := this.Head
	for i := 0; i <= index-1; i++ {
		p = p.Next
	}
	p.Next = p.Next.Next
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
