package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type QQX struct {
	QQuser int
	QQpass string
}
//const N  = 84331445
func main()  {
	allstrs := make([]QQ,N,N)

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
		linstr := string(line)
		lines := strings.Split(linstr,"----")
		if len(lines)==2{  //筛选垃圾数据  标准数据只有2个元素
			allstrs[i].QQpass=lines[1]
			allstrs[i].QQuser,_=strconv.Atoi(lines[0])

		}

		i++

	}
	fmt.Println("数据载入内存")
	time.Sleep(time.Second *10)


	startTime:= time.Now()
	for ; ;  {
		fmt.Println("请输入要查询的数据")
		var QQ int
		fmt.Scanf("%d",&QQ) //查询QQ
		for j := 0 ;j <N;j++{
			if allstrs[i].QQuser==QQ{
				fmt.Println(j,allstrs[j].QQuser,allstrs[j].QQpass) //根据数据查询QQ
			}
		}
		fmt.Println("一共用了",time.Since(startTime))
		break
	}
}


//2G的数据 用的内存非常大  需要优化   我就不进行测试了 怕
//bitcoin 的索引检查 更大 200GB