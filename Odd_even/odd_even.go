package main

import "fmt"

func OddEven(arr []int) []int  {
	isSorted := false  //奇数位，偶数位 都不需要排序的有序
	for ;isSorted==false;{
		isSorted = true
		for i := 1;i<len(arr)-1;i=i+2{
			if arr[i]>arr[i+1]{
				arr[i],arr[i+1]=arr[i+1],arr[i]
				isSorted=false
			}
		}
		fmt.Println("1",arr)
		for i:= 0;i<len(arr)-1;i=i+2{
			if arr[i]>arr[i+1]{
				arr[i],arr[i+1]=arr[i+1],arr[i]
				isSorted=false
			}

		}
		fmt.Println("0",arr)
	}
	return arr
}

func main()  {
	arr := []int{1,7,2,3,5,8,2,4}
	fmt.Println(OddEven(arr))
}