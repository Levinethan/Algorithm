package main

import "fmt"

func Select(arr []int) []int  {
	length:= len(arr)
	if length <= 1{
		return arr
	}else {
		for i := 0;i<length-1;i++{
			min := i
			for j:= i+1;j<length;j++{
				if arr[min]< arr[j]{
					min=j
				}
			}
			if i!=min{
				arr[i],arr[min]=arr[min],arr[i]
			}
		}
		return  arr

	}

}

func main()  {
	arr := []int{1,9,3,4,7,2,8,4,7,2,4,7,0}
	fmt.Println(Select(arr))
}

func SelectSortMaxx(arr []int) int  {
	length := len(arr)
	if length<=1{
		return arr[0]
	}else {
		max := arr[0]
		for i:=1;i<length;i++{
			if arr[i]>max{
				max=arr[i]
			}
		}
		return max
	}
}
