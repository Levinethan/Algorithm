package main

import (

	"fmt"
)

func main()  {
	arr := make([]int,1000,1000)
	for i:=0;i<1000;i++{
		arr[i]=i
	}
	fmt.Println(arr)
	//
	//index := bin_search(arr,124)
	//fmt.Println("index",index)
	//if index==-1{
	//	fmt.Println("找不到")
	//}else {
	//	fmt.Println("找到")
	//}
	//fmt.Println(makeFabArray([]int{1,2,3,4,5,6,7,8,9}))
	index := fab_search(arr,1112)
	fmt.Println("index",index)
	//比二分查找快一点
	//消耗的内存大一点
	if index ==-1{
		fmt.Println("err")
	}else {
		fmt.Println("find it")
	}
}

func bin_search(arr[]int,data int)int  {
	low := 0
	high := len(arr)-1

	i:=0
	for low <= high{
		i++
		fmt.Println("第N次",i)


		leftv := float64(data-arr[low])
		allv:= float64(arr[high]-arr[low])

		diff:= float64(high-low)

		mid := int(float64(low)+leftv/allv*diff) //计算出中间值
		//mid := (low+high)/2
		if mid<0 || mid>len(arr){
			return -1
		}


		if arr[mid]>data{
			high=mid-1
		}else {
			return mid
		}
	}
	return -1
}

func makeFabArray(arr []int) []int  {
	length := len(arr)
	flblen := 2
	first,secord,third:=1,2,3
	for third<length{
		third,first,secord=first+secord,secord,third
		flblen++  //叠加计算斐波拉契数列
	}
	fb := make([]int,flblen) //开辟数组
	fb[0]=1
	fb[1]=1
	for i:=2;i<flblen;i++{
		fb[i]= fb[i-1]+fb[i-2]
	}
	return fb
}

func fab_search(arr []int,val int) int  {
	length := len(arr)
	fabArr := makeFabArray(arr)
	fmt.Println(fabArr)
	filllength := fabArr[len(fabArr)-1]

	fillArr := make([]int,filllength)
	for i,v := range arr{
		fillArr[i]=v
	}


	lastdata := arr[length-1] //填充最后一共大树
	for i:=length;i<filllength;i++{
		fillArr[i]=lastdata
	}
	fmt.Println(fillArr,length)

	left,mid,right := 0,0,length
	kindex := len(fabArr)-1 //游标
	for left<right{
		mid = left+fabArr[kindex-1]-1 //斐波拉契切割
		if val < fillArr[mid]{
			right=mid-1
			kindex--
		}else if val>fillArr[mid]{
			left=mid+1
			kindex-=2
		}else {
			if mid>right{
				return right
			}else {
				return mid
			}
		}
	}
	return -1
}