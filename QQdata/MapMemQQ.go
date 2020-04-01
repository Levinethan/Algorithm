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


type QQs struct {
	QQuser int
	QQpass string
}
//const N  = 84331445
func main()  {
	allstrs := make(map[int]string,N)  //初始化map映射

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
			QQpass:=lines[1]
			QQuser,_:=strconv.Atoi(lines[0])
			allstrs[QQuser]=QQpass  //映射到Map中

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


		QQpass,err :=allstrs[QQ]
		if err{
			fmt.Println(QQ,QQpass,"存在")
		}else {
			fmt.Println("找不到")
		}

		fmt.Println("一共用了",time.Since(startTime))
		break
	}
}


//map 比数组 内存更快
//但是  内存消耗也更恐怖。。我的电脑带不动