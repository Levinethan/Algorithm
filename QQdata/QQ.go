package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)
//硬盘搜索

func main()  {
	path := "D:\\goCN\\QQ.txt"
	startTime := time.Now()

	qqfile , _ := os.Open(path)  //打开文件
	defer qqfile.Close()		//最后关闭文件
	i := 0 //统计一共多少行
	br := bufio.NewReader(qqfile)
	for {
		line , _ ,end := br.ReadLine()  //读取一行数据
		if end==io.EOF{
			break    //文件关闭
		}
		linestr := string(line)
		if strings.Contains(linestr,"YUN"){
			fmt.Println(linestr) //字符串搜索
		}
		i++
	}
	fmt.Println("一共用了",time.Since(startTime))
	fmt.Println("数据一共这么多行",i)
}
