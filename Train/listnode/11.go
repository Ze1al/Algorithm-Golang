package listnode

// LRU缓存

type LRUCache struct {
	Capacity int           // 容量
	Size     int           // 当前容量
	Cache    map[int]*Node // map
	Header   *Node         // 头结点，最新使用
	Tail     *Node         // 尾节点，最久未使用
}

type Node struct {
	Key  int
	Val  int
	Pre  *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {
	dummyHeader := &Node{}
	dummyTail := &Node{}

	dummyHeader.Next = dummyTail
	dummyTail.Pre = dummyHeader

	return LRUCache{
		Capacity: capacity,
		Size:     0,
		Cache:    make(map[int]*Node),
		Header:   dummyHeader,
		Tail:     dummyTail,
	}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.Cache[key]
	if !ok {
		return -1
	}
	this.removeNode(v)
	this.add2Header(v)
	return v.Val
}

func (this *LRUCache) Put(key int, value int) {
	v, ok := this.Cache[key]
	if ok {
		node := &Node{
			Key: key,
			Val: value,
		}
		this.Cache[key] = node
		this.removeNode(v)
		this.add2Header(node)
		return
	}

	if this.Size+1 > this.Capacity {
		this.moveTailNode()
	}

	this.Size++
	node := &Node{
		Key: key,
		Val: value,
	}
	this.Cache[key] = node
	this.add2Header(node)
	return
}

func (this *LRUCache) removeNode(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (this *LRUCache) add2Header(node *Node) {
	temp := this.Header.Next

	this.Header.Next = node
	node.Pre = this.Header
	node.Next = temp
	temp.Pre = node
}

func (this *LRUCache) moveTailNode() {
	this.Size--
	node := this.Tail.Pre
	delete(this.Cache, node.Key)
	this.removeNode(this.Tail.Pre)
}
