package Queue


type MyQueue interface {
	Size() int //大小
	Front()interface{}   //第一个
	End()  interface{}//最后一个
	IsEmpyt() bool
	EnQueue (data interface{})  //入队
	DeQueue ()interface{}
	Clear()  //清空
}

type Queue struct {
	dataStore []interface{}
	theSize int
}

func NewQueue() *Queue  {   //构造初始化的过程
	myqueue := new(Queue) //开辟结构体
	myqueue.dataStore = make([]interface{},0)
	myqueue.theSize=0
	return myqueue
}

func (myq *Queue)Size() int  {
	return myq.theSize  //返回大小
}

func (myq *Queue)Front()interface{}  {
	//求第一个
	if myq.Size()==0{
		return nil
	}
	return myq.dataStore[0]
}

func (myq *Queue)End()  interface{}  {
	if myq.Size()==0{
		return nil
	}
	return myq.dataStore[myq.Size()-1]  //取出最后一个
}

func (myq *Queue)IsEmpyt() bool  {
	return myq.theSize==0
}

func (myq *Queue)EnQueue (data interface{})  {
	myq.dataStore=append(myq.dataStore,data)  //入队操作
	myq.theSize++
}

func (myq *Queue)DeQueue ()interface{}  {
	if myq.Size()==0{
		return nil
	}
	data := myq.dataStore[0]
	if myq.Size()>1{
		myq.dataStore=myq.dataStore[1:myq.Size()]  //扔掉dataSize【0】
	}
	myq.theSize--
	return data
}

func (myq *Queue)Clear()  {
	myq.dataStore = make([]interface{},0)
	myq.theSize=0
}
