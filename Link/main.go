package main

import (
	"./Double_LinkList"
	"./Single_Link"
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main1()  {
	list := Single_Link.NewSingleLinkList()
	node1 := Single_Link.NewSingleLinkNode(1)
	node12 := Single_Link.NewSingleLinkNode(2)
	node13 := Single_Link.NewSingleLinkNode(3)
	list.InserNodeBack(node1)
	fmt.Println(list)
	list.InserNodeBack(node12)
	fmt.Println(list)
	list.InserNodeBack(node13)
	fmt.Println(list)
	node4 := Single_Link.NewSingleLinkNode(4)
	//list.InserNodeValueBack(2,node4)
	list.InsertNodeValueFront(2,node4)
	fmt.Println(list)
	fmt.Println(list.GetNodeByIndex(2))

	//list.DeleteNode(node4)
	list.DeleteIndex(1)
	fmt.Println(list)
}
func main2()  {
	list := Single_Link.NewSingleLinkList()
	path := "D:\\goCN\\140W某信用卡购物网数据.csv"
	file ,_ := os.Open(path)
	i:=0
	br := bufio.NewReader(file)
	for {
		line ,_,end := br.ReadLine()
		if end==io.EOF{
			break
		}
		linestr := string(line)
		nodestr := Single_Link.NewSingleLinkNode(linestr)
		list.InsertNodeFront(nodestr) //数据插入链表
		i++
	}
	fmt.Println("内存载入完成")

	for ; ;  {
		fmt.Println("请输入要查找的用户名")
		var inputstr string
		fmt.Scanln(&inputstr)

		start := time.Now()
		list.FindString(inputstr)



		fmt.Println("本次查询用了",time.Since(start))
	}

}

//链表的中间节点
func main3()  {
	list := Single_Link.NewSingleLinkList()
	node1 := Single_Link.NewSingleLinkNode(1)
	node2 := Single_Link.NewSingleLinkNode(2)
	node3 := Single_Link.NewSingleLinkNode(3)
	node4 := Single_Link.NewSingleLinkNode(4)
	list.InserNodeBack(node1)
	fmt.Println(list)
	list.InserNodeBack(node2)
	fmt.Println(list)
	list.InserNodeBack(node3)
	fmt.Println(list)
	list.InserNodeBack(node4)
	fmt.Println(list)
	list.ReverseList()
	fmt.Println(list)
}

func main()  {
	dlist :=Double_LinkList.NewDoubleLinkList()
	a := Double_LinkList.NewDoubleLinkNode(1)
	b := Double_LinkList.NewDoubleLinkNode(2)
	c := Double_LinkList.NewDoubleLinkNode(3)
	d := Double_LinkList.NewDoubleLinkNode(4)
	e := Double_LinkList.NewDoubleLinkNode(5)
	dlist.InsertHead(a)
	dlist.InsertHead(b)
	dlist.InsertHead(c)
	dlist.InsertHead(d)
	dlist.InsertHead(e)
	//g := Double_LinkList.NewDoubleLinkNode(6)
	//dlist.InsertValueTail(a,g)  //在a的后面插入6
	//fmt.Println(dlist.String())
	//h := Double_LinkList.NewDoubleLinkNode(6)
	//dlist.InsertValueHead(a,h)  //在a的前面插入6
	//fmt.Println(dlist.String())
	//j := Double_LinkList.NewDoubleLinkNode(6)
	//dlist.InsertValueTailByindex(3,j)
	//fmt.Println(dlist.String())
	//dlist.DeleteNode(b)
	//fmt.Println(dlist.String())
	fmt.Println(dlist.String())
	dlist.DeleteNodeByindex(3)
	fmt.Println(dlist.String())

}