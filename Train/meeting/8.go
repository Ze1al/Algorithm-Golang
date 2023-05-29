package meeting

// LRU
// Get O(1)
// Put
// Init
//

type LRU struct {
	Size   int              `json:"size"`   // 当前大小
	Cap    int              `json:"cap"`    // 容量
	Cache  map[int]*LRUNode `json:"cache"`  // map
	Header *LRUNode         `json:"header"` // 头节点
	Tail   *LRUNode         `json:"tail"`   // 尾节点
}

type LRUNode struct {
	Key  int      `json:"key"`
	Val  int      `json:"val"`
	Pre  *LRUNode `json:"pre"`
	Next *LRUNode `json:"next"`
}

// InitLRU 初始化，cap为输入的容量
func InitLRU(cap int) *LRU {
	dummyHeader := &LRUNode{}
	dummyTail := &LRUNode{}
	cache := make(map[int]*LRUNode)

	dummyHeader.Next = dummyTail
	dummyTail.Pre = dummyHeader

	return &LRU{
		Size:   0,
		Cap:    cap,
		Cache:  cache,
		Header: dummyHeader,
		Tail:   dummyTail,
	}
}

// 查询值，如果查到，则返回value
// 如果没有查到，则返回-1
func (this *LRU) Get(key int) int {
	node, ok := this.Cache[key]
	if !ok {
		return -1 // 没有查询到
	}

	this.removeNode(node)
	this.add2Header(node)
	return node.Val
}

// 存入，如果已经有存在的Key，则更新Value
// 如果不存在该Key，则新增value
func (this *LRU) Put(key, value int) {
	node, ok := this.Cache[key]
	if ok {
		newNode := &LRUNode{
			Key: key,
			Val: value,
		}

		// 将原来的节点移除
		this.removeNode(node)

		// 新加节点移动到头结点
		this.Cache[key] = newNode
		this.add2Header(newNode)
		return
	}

	// 如果加进去长度大于容量，则移除最后一个元素
	if this.Size+1 > this.Cap {
		this.removeTail()
		this.Size -= 1
	}

	// 新增Node
	this.Size += 1
	newNode := &LRUNode{
		Key: key,
		Val: value,
	}

	this.Cache[key] = newNode
	this.add2Header(newNode)
	return
}

func (this *LRU) removeNode(node *LRUNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (this *LRU) add2Header(node *LRUNode) {
	temp := this.Header.Next // 原来的header节点

	this.Header.Next = node
	node.Pre = this.Header

	node.Next = temp
	temp.Pre = node
}

func (this *LRU) removeTail() {
	tail := this.Tail.Pre
	delete(this.Cache, tail.Key)
	this.removeNode(tail)
}
