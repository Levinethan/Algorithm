package main

import "fmt"

func main()  {
	arr := []int{1,2,3,4,4,3,2,2,3,1,1,2,3,4}
	fmt.Println(arr)
	arr = BuckerSort(arr)
	fmt.Println(arr)
}

func BuckerSort(arr []int) []int  {
	length:= len(arr)
	if length<1{
		return arr
	}else {
		num := length
		max := SelectSortMax(arr)
		buckets := make([][]int,num)
		index := 0
		for i:=0;i<length;i++{
			index = arr[i]*(num-1)/max   //木桶的自动分配算法
			buckets[index]=append(buckets[index],arr[i])
		}
		fmt.Println(buckets)
		tmppose := 0 //木桶排序
		for i :=0;i<num;i++{
			bucketslen := len(buckets[i])
			if bucketslen>0{
				buckets[i]=Select(buckets[i])   //木桶内部排序
				copy(arr[tmppose:],buckets[i])   //拷贝数据
				tmppose+=bucketslen
			}
		}
		return arr
	}
}

func SelectSortMax(arr []int) int  {
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
