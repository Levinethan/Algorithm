package Link_stack_queue



type QueueLink struct {
	rear *Node   //指向头节点
	front *Node  //指向尾部节点
}

type LinkQueue interface {
	Length() int
	Enqueue (value interface{})
	Dequeue ()interface{}

}

func NewLinkQueue() *QueueLink  {
	return & QueueLink{}
}

func (qlk *QueueLink)Length() int  {
	pnext := qlk.front

	length :=0
	for pnext .pNext!=nil{
		pnext=pnext.pNext   //节点的循环跳跃
		length++
	}
	return length //返回长度
}
func (qlk *QueueLink)Enqueue (value interface{})  {
	newNode := &Node{
		data:  value,
		pNext: nil,
	} //新节点
	if qlk.front==nil{
		//如果头节点为空 直接插入数据
		qlk.front=newNode
		qlk.rear = newNode
	}else {
		qlk.rear.pNext = newNode  //地址往前挪动
		qlk.rear=qlk.rear.pNext   //尾部追加

	}
}
func (qlk *QueueLink)Dequeue ()interface{}  {
	if qlk.front==nil{
		return nil
	}
	newNode := qlk.front //记录头部位置
	if qlk.front==qlk.rear{  //只剩下一个元素时
		qlk.front=nil
		qlk.rear=nil
	}else {
		qlk.front=qlk.front.pNext  //删除一个
	}
	return newNode.data
}
