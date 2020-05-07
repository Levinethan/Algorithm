package Single_Link

import (
	"fmt"
	"strings"
)

// 单链表的接口
type SingleLink interface {
	//需求：增删查改
	GetFirstNode ()*SingleLinkNode //抓取头部节点
	InsertNodeFront(node *SingleLinkNode)   //头部插入
	InserNodeBack(node *SingleLinkNode)    //尾部插入

	//在一个节点前插入或者是节点后插入
	InsertNodeValueFront(dest interface{},node *SingleLinkNode)
	InserNodeValueBack(dest interface{},node *SingleLinkNode)

	FindString (data string)
	GetNodeByIndex (index int) *SingleLinkNode //根据索引抓取指定位置的节点
	DeleteNode (dest *SingleLinkNode) bool   //删除一个节点
	DeleteIndex (index int)   //删除指定位置的节点
	String() string   //返回链表字符串
}

type SingleLinkList struct {
	head *SingleLinkNode  //链表的头部指针
	length int    //链表的长度

}



//*SingleLinkNode 这一步 返回一个指针类型的数据
func NewSingleLinkList() *SingleLinkList  {
	head := NewSingleLinkNode(nil)
	return &SingleLinkList{
		head:   head,
		length: 0,
	}
}


func (list *SingleLinkList)InsertNodeValueFront(dest interface{},node *SingleLinkNode)bool{
	phead := list.head
	isfind := false  //是否找到了这个插入
	for  phead.pNext!=nil {
		if phead.pNext.value==dest{   //找到
			isfind=true
			break
		}
		phead=phead.pNext
	}
	if isfind{
		node.pNext = phead.pNext
		phead.pNext=node
		list.length++    //尾部插入
		return isfind
	}else {
		return isfind
	}


}
func (list *SingleLinkList)InserNodeValueBack(dest interface{},node *SingleLinkNode) bool{
	phead := list.head
	isfind := false  //是否找到了这个插入
	for  phead.pNext!=nil {
		if phead.value==dest{   //找到
			isfind=true
			break
		}
		phead=phead.pNext
	}

	if isfind{
		node.pNext = phead.pNext
		phead.pNext=node
		list.length++    //尾部插入
		return true
	}else {

		return false
	}
}

func (list *SingleLinkList)GetFirstNode () *SingleLinkNode  {
	return list.head.pNext
}

func (list *SingleLinkList)InsertNodeFront (node *SingleLinkNode)  {
	if list.head==nil{
		list.head.pNext = node
		node.pNext =nil
		list.length++  //插入节点 数据追加
	}else {
		bak := list.head   //备份头部指针
		node.pNext= bak.pNext
		bak.pNext=node
		list.length++  //插入节点 数据追加
	}
}

func (list *SingleLinkList)InserNodeBack (node *SingleLinkNode)  {
	if list.head==nil{
		list.head.pNext = node
		node.pNext =nil
		list.length++  //插入节点 数据追加
	}else {
		bak := list.head
		for bak.pNext!=nil{
			bak=bak.pNext //循环到最后
		}
		bak.pNext=node
		list.length++
	}
}

func (list *SingleLinkList) String() string {
	var listString string
	p := list.head
	for p.pNext!=nil{
		listString += fmt.Sprintf("%v-->",p.pNext.value)
		p = p.pNext //循环下去
	}
	return listString
}

func (list *SingleLinkList)GetNodeByIndex (index int) *SingleLinkNode  {
	if index > list.length-1 || index < 0{
		return nil
	}else {
		phead := list.head
		for index>-1{
			phead = phead.pNext //一直向后循环
			index -- //向后循环
		}
		return phead
	}
}

func (list *SingleLinkList)DeleteNode (dest *SingleLinkNode) bool  {
	if dest== nil{
		return false
	}
	phead := list.head
	for phead.pNext!= nil && phead.pNext!=dest{
		phead=phead.pNext
	}
	if phead.pNext==dest{


		phead.pNext=phead.pNext.pNext
		list.length--
		return true
	}else {
		return false
	}
}

func (list *SingleLinkList)DeleteIndex (index int)  {
	if index > list.length-1 || index < 0{
		return
	}else {
		phead := list.head
		for index>0{  //要找到前一个的地址才可以删除
			phead = phead.pNext //一直向后循环
			index -- //向后循环
		}

		phead.pNext=phead.pNext.pNext
		return
	}
}

func (list *SingleLinkList)FindString (data string)  {
	phead := list.head.pNext   //指定头部
	for phead.pNext!=nil{  //循环所有数据
		if strings.Contains(phead.value.(string),data){
			fmt.Println(phead.value)
		}
		phead=phead.pNext
	}
}

func (list *SingleLinkList)GetMid () *SingleLinkNode  {
	if list.head.pNext==nil{

		return nil
	}else {
		phead1 := list.head
		phead2 := list.head
		for phead2!=nil && phead2.pNext!=nil{
			phead1=phead1.pNext
			phead2=phead2.pNext.pNext
		}
		return phead1  //中间节点  1/2节点
		//求1/3  1/4节点的话要怎么求
	}
}


//面试题   反转链表
func (list *SingleLinkList)ReverseList()  {
	if list.head==nil || list.head.pNext==nil{
		//要不是链表为空 要么是只有一个节点
		return
	}else {
		var pre *SingleLinkNode   //前面节点
		var cur *SingleLinkNode = list.head.pNext  //当前节点

		for cur != nil{
			curNext := cur.pNext   //后续节点
			cur.pNext=pre  //链表反转的第一步
			pre = cur   //持续推进
			cur = curNext
		}
		fmt.Println(pre)
		list.head.pNext.pNext=nil
		list.head.pNext=pre
	}
}


//双链表