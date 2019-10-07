package structure

type arrayed struct {
	ds   []interface{}
	head uint
	tail uint
}

func NewArray() {
	return &arrayed{ds: make([]interface{}, 10), head: 0, tail: 0}
}

func (a *arrayed) Add(data interface{}) {
	a.ds = append(a.ds, data)
	a.head = a.head + 1
}

func (a *arrayed) Get(index uint) interface{} {
	if index >= 0 && a.head > index {
		return a.ds[index]
	}
	return nil
}

func (a *arrayed) Set(index uint, data interface{}) bool {
	if index >= 0 && a.head > index {
		a.ds[index] = data
		return true
	}
	return false
}

func (a *arrayed) Remove(index uint) bool {
	if index >= 0 && a.head > index {
		a.ds[index] = nil
		a.size = a.size - 1
		return true
	}
	return false
}

func (a *arrayed) Empty() {
	return a.head == 0
}
