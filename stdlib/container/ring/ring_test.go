package ring_test

import (
	"container/ring"
	"fmt"
)

func Example() {
	r := ring.New(5)
	n := r.Len()
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}
	r.Do(func(p interface{}) {
		fmt.Printf("%d ", p.(int))
	})
	// Output:
	// 0 1 2 3 4
}
