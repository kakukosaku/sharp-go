// Author: kaku
// Date: 2020/10/18
//
// GitHub:
//	https://github.com/kakukosaku
//
package grammar_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	// consumer & producer example code

	// Job chan with 10 buffer:
	ch := make(chan int, 10)

	// producer
	// add int:i to ch every second and then accumulate i, in 10 seconds
	go func() {
		i := 1
		endTime := time.After(time.Duration(10) * time.Second)
	End:
		for {
			select {
			case <-endTime:
				t.Log("[P] end!")
				close(ch)
				break End
			case <-time.Tick(time.Duration(1) * time.Second):
				t.Log("[P] add:", i)
				ch <- i
			}
			i++
		}
	}()

	// result
	rest := make(chan int, 2)

	// consumer
	workerN := 3
	worker := func(workerId int) {
		sum := 0
		for i := range ch {
			t.Logf("[C-%d]got i:%d\n", workerId, i)
			sum += i
		}
		t.Logf()
		rest <- sum
	}
	for i := 0; i < workerN; i++ {
		go worker(i)
	}

	total := 0
	for i := 0; i < workerN; i++ {
		total += <-rest
		if i == workerN-1 {
			close(rest)
		}
	}

	// actual 1+2+3+...+9
	assert.Equal(t, 45, total)
}
