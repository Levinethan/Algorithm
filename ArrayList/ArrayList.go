package ArrayList

import (
	"errors"
	"fmt"
)

//范型type
// 接口
type List interface {
	Size() int   //数组大小
	Get(index int ) (interface{},error)//抓取第几个元素  interface 代表返回的类型 然后是error
	Set(index int,newval interface{}) error   //修改元素  new新数据 标记一下新类型  尾部错误处理error
	Insert (index int,newval interface{}) error //插入元素与修改类似
	Append (newval interface{})  //追加
	Clear()  //清空数组
	Delete(index int) error //删除 第几个index
	String() string //返回字符串
}
//数据结构  字符 整数
type ArrayList struct {
	dataStore [] interface{}   //datastore 数组存储  范型
	TheSize int 				//数组大小
}

//结论  第一个是接口 第二个是范型  interface{}这个大括号的意思 是 我需要它为int的时候为int 需要string的时候为string
//注意  interface 与 interface{}不是一个东西 1.接口 2.范用类型

func (list *ArrayList)checkisFull()  {  //因为NewArray 方法设定的默认数组大小为10 所以需要判断
//如果大小>=10 则开辟双倍内存空间    开辟  拷贝  赋值
	if list.TheSize == cap(list.dataStore){
		fmt.Println("A",list.dataStore)
		//判断内存使用               make 第二个参数如果是0的话 只是允许有这么多空间 没有给你开辟空间
		newdataStore := make([]interface{},2*list.TheSize,2*list.TheSize)  //开辟双倍内存
		//copy(newdataStore,list.dataStore)   //拷贝
		for i:=0;i<len(list.dataStore);i++{
			newdataStore[i]=list.dataStore[i]
		}
		list.dataStore=newdataStore  //赋值
		fmt.Println("B",list.dataStore)
		fmt.Println("C",newdataStore)
	}   //cap 求数据空间 实际使用有多少

}

func NewArrayList() *ArrayList  {  //这一步是为NewArray 开辟内存
	list := new(ArrayList)   //首先新建一个list  初始化数据结构

	list.dataStore = make([]interface{},0,10)  //然后对数组list 开辟内存

	list.TheSize = 0

	return list
}


//之前写过一个类  arraylist
//
func (list *ArrayList)Size() int  {
	return list.TheSize   //返回数组大小

}

func(list *ArrayList) Append(newval interface{})  {
	list.dataStore = append(list.dataStore,newval)
	list.TheSize ++ //索引移动  游标移动
}

//抓取数据
func (list *ArrayList)Get(index int) (interface{},error)  {
	if index < 0 || index >=list.TheSize{
		return nil,errors.New("越界")
	}
	return list.dataStore[index],nil
}

// func Get(index int)  我们写一个get函数是可以的 但是光这么写不够 需要包装进类里 struct

func (list *ArrayList)String()string  {   //返回数组字符串
	return fmt.Sprint(list.dataStore)
}

func (list *ArrayList)Set (index int, newval interface{}) error  {
	if index < 0 || index >=list.TheSize{
		return errors.New("越界")
	}
	list.dataStore[index] = newval //设置
	return nil
}
//too many arguement to return 仔细看方法返回的类型
func (list *ArrayList)Insert(index int, newval interface{}) error  {
	if index < 0 || index >=list.TheSize{
		return errors.New("越界")
	}
	//插入往往需要开辟内存  删除的时候把后面的往前移动
	list.checkisFull()  //检测内存，如果满了 自动追加
	list.dataStore = list.dataStore[:list.TheSize+1]  //插入的时候 使用的范围是要开辟一块 移动一位
	for  i:= list.TheSize;i>index;i-- {
		list.dataStore[i]=list.dataStore[i-1]
	}//从后往前移动

	//实现插入
	list.dataStore[index] = newval
	list.TheSize++ //索引追加

	return nil
}

func (list *ArrayList)Clear () {
	list.dataStore = make([]interface{},0,10)  //然后对数组list 开辟内存

	list.TheSize = 0
	//因为go语言自动回收内存
	//那么clear 这一步就可以把原来的内存废弃 然后再重新开辟一个
}

func (list *ArrayList)Delete (index int)error {
	//删除index 简洁性写法 调用append 把0~index 叠加过来 然后与index+1 叠加到最后
	list.dataStore = append(list.dataStore[:index],list.dataStore[index+1:]...)//重新叠加  跳过index
								//[0,index) index  [index+1 , n)  叠加/合并
	list.TheSize--  //往前移动

	return nil
}