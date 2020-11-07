// Package context_test Example copy/summary from go std lib `context`
package context_test

import (
	"context"
	"fmt"
	"time"
)

func ExampleWithCancel() {
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

	// Note:
	// gen(ctx) 以传入ctx(用context.WithCancel...的方式传递ctx)的方法, 用cancel function控制由 gen(ctx) 拉起的goroutine-A!
	// cancel() 控制 gen(ctx) 拉起的goroutine-A的终止! 避免goroutine leak!
	for n := range gen(ctx) {
		if n == 5 {
			fmt.Println(n)
			break
		}
	}
	// Output:
	// 5
}

func ExampleWithDeadline() {
	shortDuration := 1 * time.Millisecond
	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
	// Output:
	// context deadline exceeded
}

func ExampleWithTimeout() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	ctxKaku := context.WithValue(ctx, "kaku", 1)
	fmt.Println(ctxKaku.Value("kaku").(int))
	ctx.Deadline()
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
	// Output:
	// 1
	// context deadline exceeded
}

func ExampleWithValue() {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	// ctx WithValue key should not be of type string or any other built-in type
	// to avoid collisions between packages using context.
	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))
	// Output:
	// found value: Go
	// key not found: color
}
