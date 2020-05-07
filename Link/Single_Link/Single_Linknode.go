package Single_Link


//链表的节点
type SingleLinkNode struct {
	value interface{}
	pNext  *SingleLinkNode
}

//构造一个节点
func NewSingleLinkNode(data interface{}) *SingleLinkNode  {
	return &SingleLinkNode{
		value: data,
		pNext: nil,
	}
}

//围绕单节点 处理相关欢节

//返回  值
func (node *SingleLinkNode)Value ()  interface{} {
	return node.value
}


//返回节点
func (node *SingleLinkNode)PNext () *SingleLinkNode  {
	return node.pNext
}

//项目有一个原则
//  一个文件只放一个类