package main

import (
	"fmt"
	"time"
)

var container chan bool //通信管道
var flag bool
var count int
func main()  {
	var array []int = []int{16,2,5,11,4}
	flag = true
	container = make(chan bool,5)
	for i :=0;i<len(array);i++{
		go tosleep(array[i])
	}
	//监听chan
	go listen(len(array))
	for flag{
		time.Sleep(1*time.Second)
	}
}

func listen(size int)  {
	for flag{
		select {
		case <-container:
			count ++ //计数器
			if count>=size{ //等待5个数字采集完成就退出
				flag=false
				break
			}

		}
	}
}

func tosleep(data int)  {
	time.Sleep(time.Duration(data)*time.Microsecond*1000)
	fmt.Println("sleep",data)
	container<- true //管道输入
}
