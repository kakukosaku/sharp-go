// Author: kaku
// Date: 2020/10/18
//
// GitHub:
//	https://github.com/kakukosaku
//
package grammar_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"sync/atomic"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	// consumer & producer example code

	// Job chan with 5 buffer:
	jobs := make(chan int, 5)
	var jobsCounter int32 = 0

	// use ctx timeout to finish this *Producer goroutine*
	ctx, cancel := context.WithCancel(context.Background())

	// *Producer goroutine*
	// add int:i to ch every second and then accumulate i
	go func(ctx context.Context, jobs chan<- int, jobsCounter *int32) {
		i := 0
		for {
			i++
			select {
			case <-ctx.Done():
				t.Log("[P] end cause by cancel called!")
				close(jobs)
				return
			default:
				jobs <- -i
				atomic.AddInt32(jobsCounter, 1)
				t.Log("[P] add:", -i)
			}
		}
	}(ctx, jobs, &jobsCounter)

	// result channel with zero-buf
	rest := make(chan int)

	// *Consumer goroutine*
	worker := func(workerId int, jobs <-chan int, rest chan<- int) {
		for i := range jobs {
			t.Logf("[C-%d] got i:%d\n", workerId, i)
			rest <- -i
		}
		t.Logf("[C-%d] end!", workerId)
	}

	// start workerN consumer goroutine
	workerN := 3
	for i := 0; i < workerN; i++ {
		go worker(i, jobs, rest)
	}

	// in another goroutine finish producer
	go func() {
		select {
		case <-time.After(10 * time.Second):
			cancel()
		}
	}()

	// calculate result!
	total := func(ctx context.Context, jobs <-chan int, rest <-chan int, jobsCounter *int32) int {
		var total int
		var c int32 = 0
		for {
			select {
			// get result every 1s, this rate limit the P/C rate actually!
			case <-time.After(1 * time.Second):
				total += <-rest
				c++
				if c == atomic.LoadInt32(jobsCounter) {
					return total
				}
			}
		}
	}(ctx, jobs, rest, &jobsCounter)

	// actual 1+2+3+...+18 why is 18?
	// 5 buf + 10(1 every 1s)
	assert.Equal(t, 171, total)
}
