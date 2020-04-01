package ArrayList

import "errors"

//设计迭代器
//主要解决访问数组方便的问题

type Iterator interface {
	//接口
	//实现接口
	HasNext()  bool //是否有下一个
	Next() (interface{},error) //下一个
	Remove() //删除
	GetIndex() int //得到索引
}

type Iterable interface {
	Iterator() Iterator  //这就是迭代器设计模式  构造初始化接口

}
//构造指针访问这个数组
type ArraylistIterator struct {
	//类    一个类里面可以包含其他类
	list *ArrayList   //数组指针  这一步 关联了类与类 意义：包含一个ArrayList 对象
	currentindex int //当前索引
}

func (list *ArrayList)Iterator() Iterator  {  //Iterator 返回的是结构体  但是这个结构体需要把函数方法全部实现才可以使用
											//Iterator -> interface -> func
	//接口实现
	it := new(ArraylistIterator)   //构造迭代器
	
	it.currentindex = 0  //代表第一个
	
	it.list=list  //等于当前值
	return  it    //返回it arraylistiterator 接口
}

func (it *ArraylistIterator) HasNext() bool {
	return it.currentindex < it.list.TheSize //是否有下一个
}

func(it *ArraylistIterator) Next() (interface{},error)  {
	if !it.HasNext(){
		return nil,errors.New("with out next node")
	}
	value , err := it.list.Get(it.currentindex)  //抓取当前数据
	it.currentindex++
	return value,err
}

func (it *ArraylistIterator)Remove()  {
	it.currentindex--
	it.list.Delete(it.currentindex)
}

func (it *ArraylistIterator)GetIndex() int   {
	return it.currentindex
}