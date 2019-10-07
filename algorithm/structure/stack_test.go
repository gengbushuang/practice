package structure

import (
	"fmt"
	"practice/algorithm/structure"

	"testing"
)

func TestPush(t *testing.T) {
	link := structure.NewLink()
	var s structure.Stack = link

	s.Push("1")
	s.Push("2")
	s.Push("3")

	d := s.Pop()
	for d != nil {
		fmt.Println(d)
		d = s.Pop()
	}

}

func TestBrackets(t *testing.T) {
	//括号是否是正常的键值对
	var brackets string = "[{}{(())}[]{{[]}}]"

	bracketsMap := map[int32]int32{')': '(', '}': '{', ']': '['}
	normal := true

	var stack structure.Stack = structure.NewLink()

	for _, bracket := range brackets {
		element, ok := bracketsMap[bracket]
		if !ok {
			stack.Push(bracket)
			continue
		}
		if stack.Empty() || element != stack.Peek() {
			normal = false
			break
		}

		stack.Pop()

	}

	if normal && stack.Empty() {
		fmt.Println(brackets, "合格")
	} else {
		fmt.Println(brackets, "不合格")
	}

}
