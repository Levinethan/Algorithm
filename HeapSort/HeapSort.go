package main

import "fmt"

func HeapSortMax (arr []int,length int) []int  {
	//length := len(arr)
	if length<=1{
		return arr
	}else {
		depth := length/2 -1
		for i := depth;i>=0 ;i--{
			topmax := i
			leftchild:=2*i+1
			rightchild := 2*i+2
			if leftchild<=length-1 &&arr[leftchild]>arr[topmax]{
				topmax=leftchild
			}
			if rightchild<=length-1 && arr[rightchild]>arr[topmax]{
				topmax=rightchild
			}
			if topmax!=i{
				arr[i],arr[topmax]=arr[topmax],arr[i]
			}

		}
		return arr
	}
}

func HeapSort(arr []int) []int  {
	length:= len(arr)
	for i := 0;i<length;i++{
		lastmesslen := length-i
		HeapSortMax(arr,lastmesslen)
		if i < length{
			arr[0],arr[lastmesslen-1]=arr[lastmesslen-1],arr[0]
		}
		fmt.Println("ex",arr)
	}
	return arr
}

func main()  {
	arr := []int{1,0,9,11,7,5,6}
	fmt.Println(HeapSort(arr))
}
