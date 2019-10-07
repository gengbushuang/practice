package structure

import (
	"practice/algorithm/structure"
	"testing"
)

func TestAdd(t *testing.T) {

	link := structure.NewLink()
	link.AddFirst("a")
	link.AddFirst("b")
	link.AddFirst("c")

	link.Show()
}
