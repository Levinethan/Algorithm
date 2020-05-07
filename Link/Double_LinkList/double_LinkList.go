package Double_LinkList

import (
	"fmt"
	"strings"
)

//双链表基本结构
//notice： 插入数据 获取长度 对一系列函数进行封装
type DoubleLinkList struct {
	head  *DoubleLinkNode
	length int
}
//返回一个Lins对象
func NewDoubleLinkList() *DoubleLinkList  {
	head := NewDoubleLinkNode(nil)
	return &DoubleLinkList{
		head:   head,
		length: 0,
	}
}
//返回链表长度
func (dlist *DoubleLinkList)GetLength () int  {
	return dlist.length
}

func (dlist *DoubleLinkList)GetFistNode() *DoubleLinkNode  {
	return dlist.head.next  //返回第一个节点
}

func (dlist *DoubleLinkList)InsertHead (node *DoubleLinkNode) {
	phead := dlist.head  //头节点
	if phead.next==nil{
		node.next=nil
		//都等于nil
		phead.next=node
		//只有一个节点直接连接上去
		node.prev=phead //上一个节点
		dlist.length++
	}else {
		phead.next.prev=node //标记上一个节点
		node.next=phead.next  //下一个节点
		phead.next=node   //标记头部节点
		node.prev=phead
		dlist.length++
	}
}

func (dlist *DoubleLinkList)InsertTail(node *DoubleLinkNode)  {
	phead := dlist.head  //头节点
	if phead.next==nil{
		node.next=nil
		//都等于nil
		phead.next=node
		//只有一个节点直接连接上去
		node.prev=phead //上一个节点
		dlist.length++
	}else {
		for phead.next!=nil{
			phead=phead.next
		}
		phead.next=node
		node.prev=phead
		dlist.length++
	}
}

func (dlist *DoubleLinkList)String() string  {
	var listString1 string
	var listString2 string

	phead := dlist.head
	for phead.next!=nil{
		//正向循环
		listString1 += fmt.Sprintf("%v -->",phead.next.value)
		phead=phead.next  //循环
	}
	listString1 += fmt.Sprintf("nil")
	listString1 += "\n"
	for phead!=dlist.head{
		//反向循环
		listString2 += fmt.Sprintf("%v <--",phead.value)
		phead=phead.prev  //循环
	}
	listString2 += fmt.Sprintf("nil")


	return listString1+listString2+"\n"


}

//指定一个节点从 前面/后面插入
func (dlist *DoubleLinkList)InsertValueTail (dest *DoubleLinkNode,node *DoubleLinkNode) bool{
	phead := dlist.head
	for phead.next!=nil  && phead.next!=dest{ //循环查找
		phead=phead.next //找到了
	}
	if phead.next==dest{
		//这里注意几个规范
		if phead.next.next!=nil{
			//对数据做单独处理 因为要在节点前面插入
			phead.next.next.prev=node
		}
		node.next=phead.next.next
		phead.next.next=node
		node.prev=phead.next

		dlist.length++
		return true
	}else {


		return false
	}
}

func (dlist *DoubleLinkList)InsertValueHead (dest *DoubleLinkNode,node *DoubleLinkNode) bool{
	phead := dlist.head  //备份的数据
	for phead.next!=nil  && phead.next!=dest{ //循环查找
		phead=phead.next //找到了
	}
	if phead.next==dest{
		//这里注意几个规范
		if phead.next.next!=nil{
			phead.next.prev=node
		}
		node.next=phead.next
		node.prev=phead
		phead.next=node

		dlist.length++
		return true
	}else {


		return false
	}
}

func (dlist *DoubleLinkList)InsertValueTailByindex (dest interface{},node *DoubleLinkNode) bool{
	phead := dlist.head
	for phead.next!=nil  && phead.next.value!=dest{ //循环查找
		phead=phead.next //找到了
	}
	if phead.next.value==dest{
		//这里注意几个规范
		if phead.next.next!=nil{
			//对数据做单独处理 因为要在节点前面插入
			phead.next.next.prev=node
		}
		node.next=phead.next.next
		phead.next.next=node
		node.prev=phead.next

		dlist.length++
		return true
	}else {


		return false
	}
}

func (dlist *DoubleLinkList)InsertValueHeadByindex (dest interface{},node *DoubleLinkNode) bool{
	phead := dlist.head  //备份的数据
	for phead.next!=nil  && phead.next.value!=dest{ //循环查找
		phead=phead.next //找到了
	}
	if phead.next.value==dest{
		//这里注意几个规范
		if phead.next.next!=nil{
			phead.next.prev=node
		}
		node.next=phead.next
		node.prev=phead
		phead.next=node
		dlist.length++
		return true
	}else {
		return false
	}
}

func (dlist *DoubleLinkList)GetNodeAtIndex(index int)  *DoubleLinkNode {
	if index > dlist.length-1 || index<0{
		return nil
	}

	phead := dlist.head
	for index > -1 {
		phead=phead.next
		index--
		//计算位置 游标

	}
	return phead
}

func (dlist *DoubleLinkList)DeleteNode(node *DoubleLinkNode) bool {
	if node == nil{
		return false
	}else {
		phead := dlist.head
		for phead.next!=nil && phead.next!=node{
			phead=phead.next //循环查找
		}
		if phead.next==node{
			if phead.next.next!=nil{
				phead.next.next.prev=phead
			}
			phead.next=phead.next.next
			dlist.length--
			return true
		}else {
			return false
		}

		return true
	}
}

func (dlist *DoubleLinkList)DeleteNodeByindex(index int) bool {
	if dlist.head == nil{
		return false
	}
	phead := dlist.head
	for index>0{
		phead=phead.next
		index--
	}
	if phead.next.next!=nil{
		phead.next.next.prev=phead
	}
	phead.next=phead.next.next
	dlist.length--
	return true
}

func (dlist *DoubleLinkList)FindString(inputstr string)  {
	phead := dlist.head.next
	for phead.next!=nil{
		if strings.Contains(phead.value.(string),inputstr){
			fmt.Println(phead.value.(string))
		}
		phead=phead.next  //循环
	}
}