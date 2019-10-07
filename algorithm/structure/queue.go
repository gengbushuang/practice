package structure

type Queue interface {
	Offer(data interface{})

	Poll() interface{}

	Peek() interface{}

	//是否已经没有元素
	Empty() bool
}
