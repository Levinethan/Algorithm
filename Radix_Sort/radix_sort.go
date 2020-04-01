package main

import (
	"fmt"
)

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

func RadixSort(arr[]int) [] int {
	max := SelectSortMaxx(arr)  //寻找数组的极大值
	for bit:=1 ; max/bit>0 ; bit*=10{
		arr=BitSort(arr,bit) //每次处理一个级别的排序
		fmt.Println(arr)
		//按照数量级分段
	}
	return arr
}

func BitSort(arr []int,bit int)[]int  {
	length := len(arr)
	bitcounts := make([]int,10) //统计长度
	for i:=0;i<length;i++{
		num := (arr[i]/bit)%10  //分层处理  bit=1000 三位数不参与排序
		bitcounts[num]++ //统计余数相等的个数
	}
	fmt.Println(bitcounts)
	// 0 1  2  3 4 5
	for i:=1;i<10;i++{
		bitcounts[i]+=bitcounts[i-1]  //叠加计算位置
	}
	fmt.Println(bitcounts)


	tmp := make([]int,10)
	for i:=length-1;i>=0;i--{
		num:= (arr[i]/bit)%10
		tmp[bitcounts[num]-1]=arr[i]  //计算排序的位置
		bitcounts[num]--
	}
	for i:= 0;i<length;i++{
		arr[i]=tmp[i]
	}
	return arr

}

func main()  {
	arr := []int{11,1119,873,44515,6624,5551}
	fmt.Println(RadixSort(arr))
}