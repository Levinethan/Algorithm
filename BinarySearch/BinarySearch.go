package main

import "fmt"

func Binary_search(arr []int,data int) int  {
	low := 0
	high := len(arr)-1
	for low<high{
		mid:=(low+high)/2
		if arr[mid]>data{
			high=mid-1
		}else if arr[mid]<data{
			low=mid+1
		}else {
			return mid
		}

	}
	return -1
}

func main()  {
	arr:=[]int{1,19,4,8,3,5,6,0}
	fmt.Println("未排序",arr)
	index := Binary_search(arr,4)
	if index==-1{
		fmt.Println("未找到")
	}else {
		fmt.Println(arr[index],index,"找到")
	}
}