package main

import "fmt"

func main()  {
	arr := []int{1,7,2,3,5,8,2,4}
	fmt.Println(MergeSort(arr))
}

func merge(leftarr []int,rightarr []int) []int  {
	leftindex := 0
	rightindex := 0
	lastarr := []int{}
	//左右索引
	for leftindex <len(leftarr) && rightindex<len(rightarr){
		if leftarr[leftindex] < rightarr[rightindex]{
			lastarr = append(lastarr,leftarr[leftindex])
			leftindex++
		}else if leftarr[leftindex] > rightarr[rightindex]{
			lastarr=append(lastarr,rightarr[rightindex])
			rightindex++
		}else  {
			lastarr = append(lastarr,rightarr[rightindex])
			lastarr = append(lastarr,leftarr[leftindex])
			leftindex++
			rightindex++
		}
	}

	for leftindex < len(leftarr){
		lastarr = append(lastarr,leftarr[leftindex])
		leftindex++
	}
	for rightindex < len(rightarr){
		lastarr = append(lastarr,rightarr[rightindex])
		rightindex++
	}
	return lastarr
}

func MergeSort(arr[]int) []int  {
	length := len(arr)
	if length<=1{
		return arr
	}else if length>1 && length<5 {
		return InserSortX(arr)   //算法改良 增加效率
	} else {
		fmt.Println(arr)
		mid := length/2
		leftarr := MergeSort(arr[:mid])
		rightarr := MergeSort(arr[mid:])
		return merge(leftarr,rightarr)
	}

}

func InserSortX(arr []int) []int  {
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
