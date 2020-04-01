package main

import "fmt"

func main()  {
	var array []int = []int{16,2,5,11,4}


	fmt.Println(CombSort(array))
}
//希尔排序终极版
func CombSort(arr []int) []int  {
	length := len(arr)
	gap := length
	for gap>1{
		gap = gap*10/13
		for i:=0;i+gap<length;i++{
			if arr[i]>arr[i+gap]{
				arr[i],arr[i+gap]=arr[i+gap],arr[i]
			}
		}
	}
	return arr
}