package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)
//内存搜索 memory search
//const N  = 84331445
func main()  {
	allstrs := make([]string,N,N)

	path := "D:\\goCN\\QQ.txt"


	qqfile , _ := os.Open(path)  //打开文件
	defer qqfile.Close()		//最后关闭文件
	i := 0
	br := bufio.NewReader(qqfile)
	for {
		line , _ ,end := br.ReadLine()  //读取一行数据
		if end==io.EOF{
			break    //文件关闭
		}
		allstrs[i]=string(line)
		i++

	}
	fmt.Println("数据载入内存")
	time.Sleep(time.Second *20)
	startTime:= time.Now()
	for ; ;  {
		fmt.Println("请输入要查询的数据")
		var inputstr string ="liang"
		fmt.Scanln(&inputstr)
		for j := 0 ;j <N;j++{
			if strings.Contains(allstrs[j],inputstr){
				fmt.Println(allstrs[j]) //字符串搜索
			}
		}
		fmt.Println("一共用了",time.Since(startTime))
		//break
	}

	//fmt.Println("一共用了",time.Since(startTime))


}
