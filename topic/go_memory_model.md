# The go Memory Model

Ref: https://golang.org/ref/mem

- [Original Article Summary or Guidance](#original-article-summary-or-guidance)

## Original Article Summary or Guidance

### Introduction

The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to
observe values produced by writes to the same variable in a different goroutine.

see also, [Memory Model Explain](https://github.com/kakukosaku/YNM/blob/master/language_topics/memory_model.md)

### Advice

Programs that modify data being simultaneously accessed by multiple goroutines must serialize such access.

To serialize access, protect the data with channel operations or other synchronization primitives such as those iin the
[sync](https://golang.org/pkg/sync/) and [sync/atomic](https://golang.org/pkg/sync/atomic/) packages.

### Happens Before

Within a single goroutine, reads and writes must behave as if the executed in the order specified by the program. That is,
compilers and processors may reorder the reads and writes executed within a single goroutine only when the reordering does
not change the behavior within that goroutine as defined by the language specification. 

Because of this reordering, the execution order observed by one goroutine may differ from the order perceived by another.
For example, if one goroutine executes a a = 1; b = 2;, another might observe the updated value of b before the updated 
value of a.

To specify the requirements of the reads and writes, we define happens before, a partial order on the execution of memory
operation in a Go program. If event `e1` happens before event `e2`, then we say that `e2` happens after `e1`. **Also, if 
`e1` does not happen before e2 and does not happen after `e2`, then we say that `e1` and `e2` happen concurrently.**

Within a single goroutine, the happens-before order is the order expressed by the program.

A read `r` of a variable `v` is allowed to observe a write `w` to `v` if both of the following hold:

1. `r` does not happen before `w`.
2. There is no other write `w` to `v` that happens after `w` but before `r`.

To guarantee that a read `r` of a variable `v` observes a particular write `w` to `v`, ensure that `w` is the only write
`r` is allowed to observe. That is, r is guaranteed to `w` if both of the following hold:

1. `w` happens before `r`.
2. Any other write to the shared variable `v` either happens before `w` or after `r`.

**翻译一下**

这一段, 严谨的论述了读操作(r)在什么情况下, 可以看到写操作(w)对变量(v)的修改; 以及什么情况下, 保证可以看到写操作(w)对变量(v)的修改.

对于什么情况下可以看到, 有如下充要条件:

1. 读(r) 不早于 写(w)...(happen-before, 是一个重要的概念, 自行理解, 不赘述)
2. 没有其它写(w2)l晚于之前的写(w), 却早于读(r)

这样的条件, 可以看到(但不一定看到!)

对于什么情况下保证可以看到, 有如下充要条件:

1. 写(w) happen-before 读(r)
2. 任何其它对于变量(v)的写操作, 要么发生在写(w)之前, 要么发生在读(r)之后

Withing a single goroutine, there is no concurrency, so the two definitions are equivalent: a read r `observes` the value
written by the most recent write `w` to `v`. When multiple goroutines access a shared variable `v`, they must use synchronization
events to establish happens-before conditions that ensure reads observe the desired writes.

**再翻译一下...**

在同一个goroutine中, 这2个条件其实是"相当"的, 读一定可以看到最近写的结果. 但在多个goroutine并发访问共享变量时, 它们必须通过**同步机制, 建立
happen-before关系**, 才能确保, 某次读能读到期望的写入值

The initialization of variable `v` with the zero value for v's type behaves as a write in the memory model.

**Reads and writes of values larger than a single machine word behave as multiple machine-word-sized operations in an unspecified
order.**

对超过一个字长度的读写, 被视为未指定顺序的对**多个单字长度**的读写!

### Synchronization

- Initialization

Program initialization runs in a single goroutine, but that goroutine may create other goroutines, which run concurrently.

If a package `p` import package `q`, the completion of `q's init functions` happens before the start of any p's.

The start of the function `main.main` happens after all `init` functions have finished. 

- Goroutine creation

The `go` statement that starts a new goroutine happens before the goroutine's execution begins.

For example, in this program:

```go
package main

var a string

func f() {
	print(a)
}

func hello() {
	a = "hello, world"
	go f()
}
```

calling `hello` will print "hello, world" at some point in the future(perhaps after "hello" has returned).

- Goroutine destruction

The exit of goroutine is not guaranteed to happen before any event in the program. For example, in this program:

```go
package main

var a string

func hello() {
	go func() { a = "hello" }()
	print(a)
}
```

the assignment to a is not followed by any synchronization event, so it is not guaranteed to be observed by any other 
goroutine. In fact, an aggressive compiler might delete the entire go statement.

If the effects of a goroutine must be observed by another goroutine, use a synchronization mechanism such as a lock or
channel communication to establish a relative ordering.

- Channel communication

Channel communication is the main method of synchronization between goroutines. Each send on a particular channel is 
matched to a corresponding receive from that channel, usually in a different goroutine.

A send on a channel happens before the corresponding receive from that channel completes.

This program:

```go
package main

var c = make(chan int, 10)
var a string

func f() {
	a = "hello, world"
	c <- 0
}

func main() {
	go f()
	<-c
	print(a)
}
```

is guaranteed to print "hello, world". The write to a happens before the send on c, which happens before the corresponding
 receive on c completes, which happens before the print.

more buffered channel discuss omit...

- Locks

The `sync` packages implements two lock data types, sync.Mutex and sync.RWMutex.

For any sync.Mutex or sync.RWMutex variable l and n < m, call n of l.Unlock() happens before call m of l.Lock() returns.

This program:

```go
package main

import "sync"

var l sync.Mutex
var a string

func f() {
	a = "hello, world"
	l.Unlock()
}

func main() {
	l.Lock()
	go f()
	l.Lock()
	print(a)
}
```

is guaranteed to print "hello, world". The first call to l.Unlock() (in f) happens before the second call to l.Lock() 
(in main) returns, which happens before the print.

- Once

The sync package provides a safe mechanism for initialization in the presence of multiple goroutines through the use of 
the Once type. Multiple threads can execute once.Do(f) for a particular f, but only one will run f(), and the other 
calls block until f() has returned.

A single call of f() from once.Do(f) happens (returns) before any call of once.Do(f) returns. 细品~

In this program:

```go
package main

import "sync"

var a string
var once sync.Once

func setup() {
	a = "hello, world"
}

func doprint() {
	once.Do(setup)
	print(a)
}

func twoprint() {
	go doprint()
	go doprint()
}
```

- Incorrect synchronization

Note that a read r may observe the value written by a write w that happens concurrently with r. Even if this occurs, it
does not imply that reads happening after r will observe writes that happened before w.

In this program:

```go
package main

var a, b int

func f() {
	a = 1
	b = 2
}

func g() {
	print(b)
	print(a)
}

func main() {
	go f()
	g()
}
```

**it can happen that g prints 2 and then 0.**

Double-checked locking is an attempt to avoid the overhead of synchronization. For example, the twoprint program might
be **incorrectly** written as:

```go
package main

import "sync"

var a string
var done bool
var once sync.Once

func setup() {
	a = "hello, world"
	done = true
}

func doprint() {
	if !done {
		once.Do(setup)
	}
	print(a)
}

func twoprint() {
	go doprint()
	go doprint()
}
```

Another incorrect idiom is busy waiting for a value, as in:

```go
package main

var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

func main() {
	go setup()
	for !done {
	}
	print(a)
}
```

As before, there is no guarantee that, in main, observing the write to done implies observing the write to a, so this 
program could print an empty string too. Worse, there is no guarantee that the write to done will ever be observed by 
main, since there are no synchronization events between the two threads.

The loop in main is not guaranteed to finish.

There are subtler variants on this theme, such as this program.

```go
package main

type T struct {
	msg string
}

var g *T

func setup() {
	t := new(T)
	t.msg = "hello, world"
	g = t
}

func main() {
	go setup()
	for g == nil {
	}
	print(g.msg)
}
```

Even if main observes g != nil and exits its loop, there is no guarantee that it will observe the initialized value for g.msg.

In all these examples, the solution is the same: use explicit synchronization.
