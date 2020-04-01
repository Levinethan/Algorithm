package main

import "fmt"

func InserSort(arr []int) []int  {
	length := len(arr)
	if length <=1{
		return arr
	}else {
		for i:=1;i<length;i++{
			backup:= arr[i]
			j:= i-1
			for j>=0 &&backup<arr[j]{
				arr[j+1]=arr[j]
				j--
			}
			arr[j+1]=backup
			fmt.Println(arr)
		}
		return arr

	}
}

func main()  {
	arr := []int{1,19,29,8,3,7,4,6,5,10}
	fmt.Println(InserSort(arr))
}
