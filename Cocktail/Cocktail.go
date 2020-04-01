package main

import "fmt"

func CockTail(arr []int) []int  {
	for i:= 0;i<len(arr)/2;i++{
		left := 0
		right := len(arr)-1
		for left <= right{
			if arr[left]> arr[left+1]{
				arr[left],arr[left+1]=arr[left+1],arr[left]
			}
			left++
			if arr[right-1] > arr[right]{
				arr[right-1],arr[right]= arr[right],arr[right-1]
			}
			right--
		}
		fmt.Println(arr)
	}
	return arr
}


func main()  {
	arr := []int{1,7,2,3,5,8,2,4}
	fmt.Println(CockTail(arr))
}
