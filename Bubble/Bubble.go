package main

import "fmt"

func BubbleSort(arr []int) []int  {
	length := len(arr)
	if length<=1{
		return arr
	}else {
		for i:=0;i<length-1;i++{
			isneed := false
			for j:=0;j<length-1-i;j++{
				if arr[j]>arr[j+1]{
					arr[j],arr[j+1]=arr[j+1],arr[j]
					isneed = true
				}
			}
			if !isneed{
				break
			}
			fmt.Println(arr)
		}
		return arr
	}
}

func main()  {
	arr:= []int{11,6,2,5,1,6,2,7,9}
	fmt.Println(BubbleSort(arr))
}
