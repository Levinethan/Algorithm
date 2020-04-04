package ArrayStack


type StackArray interface {
	Clear()  //清空
	Size() int  //大小
	Pull() interface{}   //
	Push(data interface{})   //push in
	IsFull() bool   //Is there full or not
	IsEmpt() bool   //Is there empt or not
}

type Stack struct {
	dataSource []interface{}
	capsize int	//最大范围
	currentsize int //当前实际使用大小
}
//首先需要一个NewStack 返回一个栈   *Stack
func NewStack() *Stack {
	mystack := new(Stack)    //开辟一个结构体
	mystack.dataSource=make([]interface{},0,10)  //对dataSource作为数据库 用make  interface来接收数据
	mystack.capsize=10  // 栈最大空间
	mystack.currentsize=0  //栈起始大小
	return mystack   //返回个人定制的栈
}

func (mystack *Stack)Clear()  {   //清空栈  自动回收内存  就是重新分配内存给0就行了
	mystack.dataSource=make([]interface{},0,10)
	mystack.currentsize=0   //把数据重新设置为0
	mystack.capsize=10
}

func (mystack *Stack)Size() int  {
	return mystack.currentsize
}

func (mystack *Stack)Pull() interface{} {
	if !mystack.IsEmpt(){   //if stack is empty it doesn`t pop anything

		last := mystack.dataSource[mystack.currentsize-1] // pull the last elem
		mystack.dataSource=mystack.dataSource[:mystack.currentsize-1] //delte the last elem
		mystack.currentsize--//  cap -1
		return last
	}
	return nil
}
func (mystack *Stack)Push(data interface{})  {
	if !mystack.IsFull(){	//also if stack is full so you couldn`t pull something in stack
		mystack.dataSource=append(mystack.dataSource,data)
		mystack.currentsize++
		//push in stack
	}
}
func (mystack *Stack)IsFull() bool  {
	if mystack.currentsize==mystack.capsize{
		return true
	}else {
		return false
	}
}
func (mystack *Stack)IsEmpt() bool  {
	if mystack.currentsize==0{
		return true
	}else {
		return false
	}
}
