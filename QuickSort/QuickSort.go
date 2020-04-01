package main

import "fmt"

func main()  {
	arr := []int{11,1119,873,44515,6624,5551}
	fmt.Println("未排序",arr)
	arr = QuickSort(arr)
	fmt.Println(arr)
	index := Binary_search(arr,11)
	if index==-1{
		fmt.Println("未找到")
	}else {
		fmt.Println(arr[index],index,"找到")
	}
}

func QuickSort(arr []int) []int  {
	length := len(arr)
	if length<=1{
		return arr
	}else {
		splitdata := arr[0] //第一个数字
		low := make([]int,0,0)
		mid := make([]int,0,0)
		high := make([]int,0,0)
		mid=append(mid,splitdata) //保存分离的数据
		for i:=1;i<length;i++{
			if arr[i]<splitdata{
				low=append(low,arr[i])
			}else if arr[i]>splitdata{
				high=append(high,arr[i])
			}else {
				mid = append(mid ,arr[i])
			}
		}
		low,high= QuickSort(low),QuickSort(high) //递归循环
		myarr :=append(append(low,mid...),high...)
		return myarr
	}

}

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




