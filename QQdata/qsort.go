package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type QQ struct {
	QQuser int
	QQpass string
}

func main()  {

	const N  = 84331445
	path := "D:\\goCN\\QQ.txt"
	allstrs := make([]QQ,N,N)

	qqfile , _ := os.Open(path)  //打开文件
	defer qqfile.Close()		//最后关闭文件
	i := 0
	br := bufio.NewReader(qqfile)
	for {
		line , _ ,end := br.ReadLine()  //读取一行数据
		if end==io.EOF{
			break    //文件关闭
		}
		linstr := string(line)
		lines := strings.Split(linstr,"----")
		if len(lines)==2{  //筛选垃圾数据  标准数据只有2个元素
			allstrs[i].QQpass=lines[1] //下面这里  字符串转换为整数
			allstrs[i].QQuser,_=strconv.Atoi(lines[0])

		}
		if i==N{
			break
		}

		i++

	}
	fmt.Println("数据载入内存")
	time.Sleep(time.Second *10)
	//排序
	fmt.Println("开始排序")
	startTime:= time.Now()
	allstrs=QuickSortStruct(allstrs)
	fmt.Println("排序结束，一共用了",time.Since(startTime))


	startTimes:= time.Now()
	for ; ;  {
		fmt.Println("请输入要查询的数据")
		var QQ int
		fmt.Scanf("%d",&QQ) //查询QQ
		//顺序查找  修改为二分查找
		index:=Binary_search(allstrs,QQ)
		if index==-1{
			fmt.Println("数据找不到")
		}else {
			fmt.Println("数据找到",index,allstrs[i].QQuser,allstrs[i].QQpass)
		}
		fmt.Println("一共用了",time.Since(startTimes))
		break
	}
}


func QuickSortStruct(arr []QQ) []QQ  {
	length := len(arr)
	if length<=1{
		return arr
	}else {
		splitdata := arr[0].QQuser //第一个数字
		low := make([]QQ,0,0)
		mid := make([]QQ,0,0)
		high := make([]QQ,0,0)
		mid=append(mid,arr[0]) //保存分离的数据
		for i:=1;i<length;i++{
			if arr[i].QQuser<splitdata{
				low=append(low,arr[i])
			}else if arr[i].QQuser>splitdata{
				high=append(high,arr[i])
			}else {
				mid = append(mid ,arr[i])
			}
		}
		low,high= QuickSortStruct(low),QuickSortStruct(high) //递归循环
		myarr :=append(append(low,mid...),high...)
		return myarr
	}

}

func Binary_search(arr []QQ,data int) int  {
	low := 0
	high := len(arr)-1
	for low<high{
		mid:=(low+high)/2
		if arr[mid].QQuser>data{
			high=mid-1
		}else if arr[mid].QQuser<data{
			low=mid+1
		}else {
			return mid
		}
	}
	return -1
}

//8g内存爆了。

//136414786
//花了2秒   如果用插入排序,选择排序需要3天

