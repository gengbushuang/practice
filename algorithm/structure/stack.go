package structure

type Stack interface {

	//元素入栈
	Push(data interface{})

	//返回栈顶元素，不出栈
	Peek() interface{}

	//返回栈顶元素，出栈，移除元素
	Pop() interface{}
	//是否已经没有元素
	Empty() bool
}
