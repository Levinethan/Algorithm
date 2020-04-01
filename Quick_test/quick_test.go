package Quick_test

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

const N  =  84331445
type QQ struct {
	QQuser int
	pd string
}
func sortForMerge(arr []QQ,left int,right int)  {
	for i:=left;i<=right;i++{
		temp:=arr[i]//备份数据
		var j int
		for j=i;j>left && arr[j-1].QQuser>temp.QQuser;j--{ //定位
			arr[j]=arr[j-1] //数据往后移动
		}
		arr[j]=temp //插入
	}
}
func swap(arr []QQ,i int,j int)  {
	arr [i],arr[j]=arr[j],arr[i]
}

//3 1 8 2 3 9 7
//


//递归快速排序
func QucikSortX(arr []QQ,left int,right int)  {
	if right-left<15{  //数组剩下三个元素，直接插入排序
		sortForMerge(arr,left,right)
	}else {
		//随机找一共数字, 放在第一个位置
		swap(arr,left,rand.Int()%(right-left+1)+left)
		vdata := arr[left] //坐标数组 比我小去左边  比我大 去右边
		lt := left  //arr[left+1,lt] < vdata
		gt := right+1 //arr[gt.. right]>vdata
		i := left+1 //arr[lt+1,..i]==vdata
		for i<gt{
			if arr[i].QQuser < vdata.QQuser{
				swap(arr,i,lt+1)
				lt++
				i++
			}else if arr[i].QQuser > vdata.QQuser{
				swap(arr,i,gt-1) //移动到大于的地方
				gt--

			}else {
				i++
			}
		}
		swap(arr,left,lt)   //交换头部位置
		QucikSortX(arr,left,lt-1)
		QucikSortX(arr,gt,right)


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
//快速排序核心程序
func QuickSortPlus(arr []QQ)  {
	QucikSortX(arr,0,len(arr)-1)
}

func Test_Quick(t *testing.T)  {
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
			allstrs[i].pd=lines[1] //下面这里  字符串转换为整数
			allstrs[i].QQuser,_=strconv.Atoi(lines[0])

		}

		i++

	}
	fmt.Println("数据载入内存",i)
	//排序
	fmt.Println("开始排序")
	startTime:= time.Now()
	QuickSortPlus(allstrs)
	fmt.Println("排序结束，一共用了",time.Since(startTime))
	for ; ;  {
		fmt.Println("请输入要查询的用户名")
		var inputstr int
		fmt.Scanf("%d",&inputstr) //用户输入

		starttime:=time.Now()
		index:= Binary_search(allstrs,inputstr)
		fmt.Println("index",index)
		if index == -1{
			fmt.Println("找不到")
		}else{
			fmt.Println("找到",allstrs[index])
		}


		fmt.Println("本次查询用了",time.Since(starttime))

	}
}