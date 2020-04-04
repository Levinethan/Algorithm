package Link_stack_queue


type Node struct {
	data  interface{}
	pNext *Node  //指向下一个节点
}

type LinkStack interface {
	IsEmpty ()bool
	Push(data interface{}) //压入
	Pop() interface{}
	Length() int //链式栈 长度
}

func NewStack() *Node  {
	return &Node{}  //返回一个节点指针
}

func (n *Node)IsEmpty ()bool  {
	if n.pNext==nil{
		return true
	}else {
		return  false
	}
}
func (n *Node) Push(data interface{}) {
	newnode :=&Node{
		data:  data,
		pNext:nil,
	}
	newnode.pNext=n.pNext
	n.pNext=newnode  //头部插入
}
func (n *Node) Pop() interface{} {
	if n.IsEmpty() == true{
		return nil
	}
	value := n.pNext.data //要弹出的数据
	n.pNext=n.pNext.pNext //删除
	return value
}
func (n *Node) Length() int {
	pnext := n
	length :=0
	for pnext .pNext!=nil{
		pnext=pnext.pNext   //节点的循环跳跃
		length++
	}
	return length //返回长度
}
