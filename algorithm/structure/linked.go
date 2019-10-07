package structure

import (
	"fmt"
)

//链表节点结构
type node struct {
	d interface{}
	//后驱节点
	next *node
	//前驱节点
	prev *node
}

type linked struct {
	first *node
	last  *node
	size  int
}

func NewLink() *linked {
	return &linked{size: 0}
}

func (link *linked) Add(data interface{}) {
	link.AddLast(data)
}

func (link *linked) Get(index int) (interface{}, bool) {
	if !link.isElementIndex(index) {
		return nil, false
	}
	n := link.getNode(index)
	return n.d, true
}

/*
头节点添加数据
*/
func (link *linked) AddFirst(data interface{}) {
	f := link.first

	newNode := &node{
		d: data,
		//新加入的节点后驱节点指向上个头节点
		next: f,
		prev: nil,
	}

	link.first = newNode
	if f == nil {
		//如果上次头节点是空，就跟尾节点指向同一个节点
		link.last = newNode
	} else {
		//如果上次头节点不为空，把上个头节点前驱节点指向新加入节点
		f.prev = newNode
	}

	link.size = link.size + 1
}

/*
尾节点添加数据
*/
func (link *linked) AddLast(data interface{}) {
	l := link.last

	newNode := &node{
		d:    data,
		next: nil,
		//新加入的节点前驱节点指向上个尾节点
		prev: l,
	}

	link.last = newNode

	if l == nil {
		//如果上次尾节点是空，就跟头节点指向同一个节点
		link.first = newNode
	} else {
		//如果上次尾节点不为空，把上个尾节点后驱节点指向新加入节点
		l.next = newNode
	}

	link.size = link.size + 1
}

func (link *linked) getNode(index int) *node {
	n := link.size

	if index < (n >> 1) {
		nNext := link.first
		for i := 0; i < index; i++ {
			nNext = nNext.next
		}
		return nNext
	} else {
		nPrev := link.last
		for i := n - 1; i < index; i-- {
			nPrev = nPrev.prev
		}
		return nPrev
	}
}

/*
删除指定的节点
*/
func (link *linked) Remove(index int) (interface{}, bool) {

	if !link.isElementIndex(index) {
		return nil, false
	}

	n := link.getNode(index)
	link.removeNode(n)
	return n.d, true
}

func (link *linked) removeFirst(n *node) {
	//当前要删除的节点的后驱节点
	nNode := n.next

	link.first = nNode
	if nNode == nil {
		link.last = nil
	} else {
		nNode.prev = nil
	}
	link.size = link.size - 1
}

func (link *linked) removeLast(n *node) {
	//当前要删除的节点的前驱节点
	pNode := n.prev

	link.last = pNode

	if pNode == nil {
		link.first = nil
	} else {
		pNode.next = nil
	}
	link.size = link.size - 1
}

/*
删除node节点
*/
func (link *linked) removeNode(n *node) {
	//当前要删除的节点的后驱节点
	nNode := n.next
	//当前要删除的节点的前驱节点
	pNode := n.prev

	if pNode == nil {
		//如果前驱节点为空，就证明要删除的是头节点，重新设置头节点
		link.first = nNode
	} else {
		//前驱节点不为空，就重新把前驱的后驱节点指向要删除的后驱节点
		pNode.next = nNode
	}

	if nNode == nil {
		//如果后驱节点为空，就证明要删除的是尾节点，重新在设置尾节点
		link.last = pNode
	} else {
		//如果后驱节点不为空，就把前驱节点指向要删除的前驱节点
		nNode.prev = pNode
	}

	link.size = link.size - 1
}

func (link *linked) isElementIndex(index int) bool {
	return link.size >= 0 && link.size > index
}

//反转一个链表
func (link *linked) Reverse() {
	nNode := link.first

	link.first, link.last = link.last, link.first

	for nNode != nil {
		tNode := nNode.next
		nNode.next, nNode.prev = nNode.prev, nNode.next
		nNode = tNode
	}

}

func (link *linked) Show() {
	fmt.Println("\n-----------------next\n")
	for nNode := link.first; nNode != nil; nNode = nNode.next {
		fmt.Print(nNode, " ")
	}

	fmt.Println("\n-----------------prev\n")
	for pNode := link.last; pNode != nil; pNode = pNode.prev {
		fmt.Print(pNode, " ")
	}
}

func (link *linked) Empty() bool {
	return link.size == 0
}

////////////////////////////栈的实现
func (link *linked) Push(data interface{}) {
	link.AddFirst(data)
}

func (link *linked) Pop() interface{} {
	f := link.first
	if f == nil {
		return nil
	}
	link.removeFirst(f)
	return f.d
}

////////////////////////////队列的实现

func (link *linked) Offer(data interface{}) {
	link.AddLast(data)
}

func (link *linked) Poll() interface{} {
	f := link.first
	if f == nil {
		return nil
	}
	link.removeFirst(f)
	return f.d
}

///////////////////////////队列先入元素先出 栈后入元素先出
func (link *linked) Peek() interface{} {
	if link.first != nil {
		return link.first.d
	}
	return nil
}
