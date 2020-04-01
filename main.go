package main

import (
	"./ArrayList"
	"fmt"
)
func main2()  {
	list := ArrayList.NewArrayList()
	list.Append("a")
	list.Append("b")
	list.Append("c")     //Append 可以添加任意类型的数据  范型的作用在此
	fmt.Println(list.TheSize)

}
//如何定义一个接口 然后接口把这个类挂起来
func main3()  {
	//定义接口对象，赋值的对象必须实现接口的所有方法
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append("a1")
	fmt.Println(list)
	//需要实现接口里面的所有方法 否则报错
	//does not implement (missing Clear method)
}

func main4()  {
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append("a1")
	list.Append("b2")
	list.Append("c3")
	for i := 0;i<4 ;i ++{
		list.Insert(1,"x5")
		fmt.Println(list)
	}//这一步理论上可以使内存溢出
	fmt.Println("delete")
	list.Delete(2)  //删除没问题
	fmt.Println(list)
}

func main()  {
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append("a1")
	list.Append("b2")
	list.Append("c3")
	for i := 0;i<55 ;i ++{
		list.Insert(1,"x5")
		fmt.Println(list)
	}//这一步理论上可以使内存溢出

	fmt.Println(list)
}