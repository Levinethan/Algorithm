package main

import "fmt"

func main()  {
	arr := []int{1,2,3,4,5,6,7,8,2,33,4}
	for i:=0;i<len(arr);i++{
		fmt.Println("index",i,arr[i])
	}
	fmt.Println(Binary_search(arr,3))
}

//A找到一个等于3的
//B找到最后一共等于3的
//C找到第一个大于等于2的
//D再找到最后一个小于7的数据
func Binary_search(arr []int,data int) int  {
	low := 0
	high := len(arr)-1
	index :=-1
	for low<high{
		mid:=(low+high)/2
		if arr[mid]>data{
			high=mid-1
		}else if arr[mid]<data{
			low=mid+1
		}else {
			if mid ==0 || arr[mid-1]!=data {  //mid ==0 找第一个 。mid==len(arr)-1 ||arr[mid+1]!=data 最后一个
				index=mid
				fmt.Println("find it",mid)
				break
			}else {
				high=mid-1  //递归继续查找
			}
			//find it
		}

	}
	return index
}
//先处理小数据 然后再处理大数据
