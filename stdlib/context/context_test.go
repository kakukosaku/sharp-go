package context_test

import (
	"context"
	"fmt"
	"time"
)

func Example()  {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		// 重点在于, gen 发起的goroutine-A无需要手动close一个single channel 在 select中
		// 主函数退出时 defer cancel 使用 ctx wait 在 ctx.Done() goroutine-A 也退出, 避免goroutine leak
		if n == 5 {
			fmt.Println(n)
			break
		}
	}
	// Output:
	// 5
	time.After()
}