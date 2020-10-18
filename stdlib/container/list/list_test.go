package list_tset

import (
	"container/list"
	"fmt"
)

func Example() {
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value.(int))
	}
	// Output:
	// 1 2 3 4
}
