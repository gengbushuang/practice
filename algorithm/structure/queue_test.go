package structure

import (
	"fmt"
	"practice/algorithm/structure"
	"testing"
)

func TestAdd(t *testing.T) {
	link := structure.NewLink()

	var q structure.Queue = link

	q.Offer(1)
	q.Offer(2)
	q.Offer(3)

	d := q.Poll()
	for d != nil {
		fmt.Println(d)
		d = q.Poll()
	}
}
